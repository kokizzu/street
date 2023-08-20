package domain

import (
	"street/model/mProperty/rqProperty"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file RealtorOwnedProperties.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type RealtorOwnedProperties.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type RealtorOwnedProperties.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type RealtorOwnedProperties.go
//go:generate farify doublequote --file RealtorOwnedProperties.go

type (
	RealtorOwnedPropertiesIn struct {
		RequestCommon

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	RealtorOwnedPropertiesOut struct {
		ResponseCommon

		Properties []rqProperty.Property `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
)

const (
	RealtorOwnedPropertiesAction = `realtor/ownedProperties`
)

func (d *Domain) RealtorOwnedProperties(in *RealtorOwnedPropertiesIn) (out RealtorOwnedPropertiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {

		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	out.Properties = r.FindOwnedByPagination(sess.UserId, &in.Pager, &out.Pager)
	return
}
