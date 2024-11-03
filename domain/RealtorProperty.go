package domain

import (
	"street/model/mProperty/rqProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file RealtorProperty.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type RealtorProperty.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type RealtorProperty.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type RealtorProperty.go
//go:generate farify doublequote --file RealtorProperty.go

type (
	RealtorPropertyIn struct {
		RequestCommon

		Id uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}
	RealtorPropertyOut struct {
		ResponseCommon

		Property *rqProperty.PropertyWithNote `json:"property" form:"property" query:"property" long:"property" msg:"property"`
	}
)

const (
	RealtorPropertyAction = `realtor/property`

	ErrPropertyNotFound = `realtor property not found`
)

func (d *Domain) RealtorProperty(in *RealtorPropertyIn) (out RealtorPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrPropertyNotFound)
		return
	}
	in.RefId = in.Id
	r.NormalizeFloorList()
	r.Adapter = nil

	propertyWithNote := r.ToPropertyWithNote()
	out.Property = &propertyWithNote
	return
}
