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
	UserRegisterIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	UserRegisterOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
)

const UserRegisterAction = `UserRegister`
const minPassLength = 12

const (
	ErrUserRegisterEmailInvalid       = `email must be valid`
	ErrUserRegisterPasswordTooShort   = `password must be at least 12 characters`
	ErrUserRegisterEmailUsed          = `email already used`
	ErrUserRegisterUserCreationFailed = `user creation failed`
)

func (d *Domain) UserRegister(in *UserRegisterIn) (out UserRegisterOut) {
	in.Email = S.Trim(S.ValidateEmail(in.Email))
	if in.Email == `` {
		out.SetError(400, ErrUserRegisterEmailInvalid)
		return
	}
	if len(in.Password) < minPassLength {
		out.SetErrorf(400, ErrUserRegisterPasswordTooShort)
		return
	}
	exists := rqAuth.NewUsers(d.UserOltp)
	exists.Email = in.Email
	if exists.FindByEmail() {
		out.SetError(451, ErrUserRegisterEmailUsed)
		return
	}
	user := wcAuth.NewUsersMutator(d.UserOltp)
	user.Email = in.Email
	user.Password = S.EncryptPassword(in.Password)
	user.CreatedAt = in.Now
	if !user.DoInsert() {
		out.SetError(500, ErrUserRegisterUserCreationFailed)
		return
	}
	user.Password = ``
	out.User = user.Users
	return
}
