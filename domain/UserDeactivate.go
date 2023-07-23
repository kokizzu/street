package domain

import (
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"

	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserDeactivate.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserDeactivate.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserDeactivate.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserDeactivate.go
//go:generate farify doublequote --file UserDeactivate.go

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
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
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
		d.Log.Err(err).Msg(ErrUserDeactivateWrongPass)
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
		d.Log.Error().Msg(errStr)
		out.SetError(500, ErrUserDeactivateLogoutFailed)
		return
	}

	_ = logins
	return
}
