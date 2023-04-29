package domain

import (
	"github.com/kokizzu/gotro/S"

	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file user.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type user.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type user.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type user.go
//go:generate farify doublequote --file user.go

type (
	GuestRegisterIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestRegisterOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
)

const (
	GuestRegisterAction = `guest/register`

	ErrGuestRegisterEmailInvalid       = `email must be valid`
	ErrGuestRegisterPasswordTooShort   = `password must be at least 12 characters`
	ErrGuestRegisterEmailUsed          = `email already used`
	ErrGuestRegisterUserCreationFailed = `user creation failed`

	minPassLength = 12
)

func (d *Domain) GuestRegister(in *GuestRegisterIn) (out GuestRegisterOut) {
	in.Email = S.Trim(S.ValidateEmail(in.Email))
	if in.Email == `` {
		out.SetError(400, ErrGuestRegisterEmailInvalid)
		return
	}
	if len(in.Password) < minPassLength {
		out.SetErrorf(400, ErrGuestRegisterPasswordTooShort)
		return
	}
	exists := rqAuth.NewUsers(d.AuthOltp)
	exists.Email = in.Email
	if exists.FindByEmail() {
		out.SetError(400, ErrGuestRegisterEmailUsed)
		return
	}
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email
	user.Password = S.EncryptPassword(in.Password)
	user.CreatedAt = in.Now
	if !user.DoInsert() {
		out.SetError(500, ErrGuestRegisterUserCreationFailed)
		return
	}
	user.Password = ``
	out.User = user.Users
	return
}

type (
	GuestLoginIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestLoginOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
)

const (
	GuestLoginAction = `guest/login`

	ErrGuestLoginEmailInvalid             = `email must be valid`
	ErrGuestLoginEmailOrPasswordIncorrect = `incorrect email or password`
	ErrGuestLoginPasswordOrEmailIncorrect = `incorrect password or email`
)

func (d *Domain) GuestLogin(in *GuestLoginIn) (out GuestLoginOut) {
	in.Email = S.Trim(S.ValidateEmail(in.Email))
	if in.Email == `` {
		out.SetError(400, ErrGuestLoginEmailInvalid)
		return
	}
	user := rqAuth.NewUsers(d.AuthOltp)
	user.Email = in.Email
	if !user.FindByEmail() {
		out.SetError(400, ErrGuestLoginEmailOrPasswordIncorrect)
		return
	}
	if err := S.CheckPassword(user.Password, in.Password); err != nil {
		out.SetError(400, ErrGuestLoginPasswordOrEmailIncorrect)
		return
	}
	user.Password = ``
	out.User = *user
	return
}
