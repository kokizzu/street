package domain

import (
	"street/model/mAuth/rqAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestDebug.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestDebug.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestDebug.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestDebug.go
//go:generate farify doublequote --file GuestDebug.go

type (
	GuestDebugIn struct {
		RequestCommon
	}
	GuestDebugOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`
	}
)

const (
	GuestDebugAction = `guest/debug`
)

func (d *Domain) GuestDebug(in *GuestDebugIn) (out GuestDebugOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon
	return
}

type (
	GuestRegisterIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestRegisterOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		verifyEmailUrl string
	}
)
