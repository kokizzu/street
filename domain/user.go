package domain

import (
	"log"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/M"
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
	UserLogoutIn struct {
		RequestCommon
	}
	UserLogoutOut struct {
		ResponseCommon
		LogoutAt int64 `json:"loggedOut" form:"loggedOut" query:"loggedOut" long:"loggedOut" msg:"loggedOut"`
	}
)

const (
	UserLogoutAction = `user/logout`

	ErrUserSessionRemovalFailed = `user session removal failed`
)

func (d *Domain) UserLogout(in *UserLogoutIn) (out UserLogoutOut) {
	out.LogoutAt = d.expireSession(in.SessionToken, &out.ResponseCommon)
	return
}

type (
	UserProfileIn struct {
		RequestCommon
	}
	UserProfileOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	UserProfileAction = `user/profile`

	ErrUserProfileNotFound = `user not found`
)

func (d *Domain) UserProfile(in *UserProfileIn) (out UserProfileOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(403, ErrUserProfileNotFound)
		return
	}
	out.actor = user.Id

	user.CensorFields()
	out.User = user
	out.Segments = sess.Segments
	return
}

type (
	UserChangePasswordIn struct {
		RequestCommon
		OldPass string `json:"oldPass" form:"oldPass" query:"oldPass" long:"oldPass" msg:"oldPass"`
		NewPass string `json:"newPass" form:"newPass" query:"newPass" long:"newPass" msg:"newPass"`
	}
	UserChangePasswordOut struct {
		ResponseCommon
		ok bool
	}
)

const (
	UserChangePasswordAction = `user/changePassword`

	ErrUserChangePasswordUserNotFound    = `user to change password not found`
	ErrUserChangePasswordWrongOldPass    = `old password does not match`
	ErrUserChangePasswordNewPassTooShort = `new password too short`
	ErrUserChangePasswordSaveUserFailed  = `failed saving user`
)

func (d *Domain) UserChangePassword(in *UserChangePasswordIn) (out UserChangePasswordOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if len(in.NewPass) < minPassLength {
		out.SetError(400, ErrUserChangePasswordNewPassTooShort)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserChangePasswordUserNotFound)
		return
	}

	if err := user.CheckPassword(in.OldPass); err != nil {
		out.SetError(400, ErrUserChangePasswordWrongOldPass)
		return
	}

	user.SetEncryptedPassword(in.NewPass, in.UnixNow())
	if !user.DoUpdateById() {
		out.SetError(500, ErrUserChangePasswordSaveUserFailed)
		return
	}

	out.ok = true
	return
}

type (
	UserDeactivateIn struct {
		RequestCommon
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}

	UserDeactivateOut struct {
		ResponseCommon
	}
)

const (
	UserDeactivateAction = `user/deactivate`

	ErrUserDeactivateNotFound     = `user to deactivate not found`
	ErrUserDeactivateWrongPass    = `wrong password`
	ErrUserDeactivateUpdateFailed = `user deactivation failed`
	ErrUserDeactivateLogoutFailed = `user deactivation logout all session failed`
)

func (d *Domain) UserDeactivate(in *UserDeactivateIn) (out UserDeactivateOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserDeactivateNotFound)
		return
	}

	if err := user.CheckPassword(in.Password); err != nil {
		log.Println(err)
		out.SetError(400, ErrUserDeactivateWrongPass)
		return
	}

	user.SetDeletedAt(in.UnixNow())
	arr := S.Split(user.Email, `@`)
	arr[0] += `DEL` + I.UToS(user.Id)
	newEmail := A.StrJoin(arr, `@`)

	user.SetEmail(newEmail)

	if !user.DoUpdateById() {
		out.SetError(500, ErrUserDeactivateUpdateFailed)
		return
	}

	sm := wcAuth.NewSessionsMutator(d.AuthOltp)
	logins, errStr := sm.ForceLogoutAll(user.Id, in.UnixNow())
	if errStr != `` {
		log.Println(errStr)
		out.SetError(500, ErrUserDeactivateLogoutFailed)
		return
	}

	_ = logins
	return
}

type (
	UserUpdateProfileIn struct {
		RequestCommon
		UserName string `json:"userName" form:"userName" query:"userName" long:"userName" msg:"userName"`
		FullName string `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}

	UserUpdateProfileOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
)

const (
	UserUpdateProfileAction = `user/updateProfile`

	ErrUpdateProfileUsernameAlreadyUsed = `username already used`
	ErrUpdateProfileEmailAlreadyUsed    = `email already used`
	ErrUpdateProfileFailed              = `update profile failed`
)

func (d *Domain) UserUpdateProfile(in *UserUpdateProfileIn) (out UserProfileOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserProfileNotFound)
		return
	}

	if in.UserName != `` && user.UserName != in.UserName {
		dup := rqAuth.NewUsers(d.AuthOltp)
		dup.UserName = S.ValidateIdent(in.UserName)
		if dup.FindByUserName() && dup.Id != user.Id {
			out.SetError(400, ErrUpdateProfileUsernameAlreadyUsed)
			return
		}
		user.SetUserName(dup.UserName)
	}

	if in.Email != `` && user.Email != in.Email {
		dup := rqAuth.NewUsers(d.AuthOltp)
		dup.Email = S.ValidateEmail(in.Email)
		if dup.FindByEmail() && dup.Id != user.Id {
			out.SetError(400, ErrUpdateProfileEmailAlreadyUsed)
			return
		}
		user.SetEmail(dup.Email)
		user.SetVerifiedAt(0) // must also unset verifiedAt
	}

	if in.FullName != `` && user.FullName != in.FullName {
		user.SetFullName(in.FullName)
	}

	if !user.DoUpdateById() {
		user.HaveMutation()
		out.SetError(400, ErrUpdateProfileFailed)
		return
	}

	out.AddDbChangeLogs(user.Logs())

	user.CensorFields()
	out.User = &user.Users
	return
}
