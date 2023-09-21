package domain

import (
	"fmt"

	"street/model/mProperty/rqProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestProperty.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestProperty.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestProperty.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestProperty.go
//go:generate farify doublequote --file GuestProperty.go

type (
	GuestPropertyIn struct {
		RequestCommon
		Id uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}
	GuestPropertyOut struct {
		ResponseCommon
		Property *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
	}
)

const (
	GuestPropertyAction      = `guestProperty`
	ErrGuestPropertyNotFound = `Property not found`
)

func (d *Domain) GuestProperty(in *GuestPropertyIn) (out GuestPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrGuestPropertyNotFound)
		return
	}
	r.NormalizeFloorList()
	out.Property = r
	fmt.Println(r)
	return
}
