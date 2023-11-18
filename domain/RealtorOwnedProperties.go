package domain

import (
	"street/model/mProperty"
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

		ShowMeta bool `json:"showMeta" form:"showMeta" query:"showMeta" long:"showMeta" msg:"showMeta"`
	}
	RealtorOwnedPropertiesOut struct {
		ResponseCommon

		Properties []rqProperty.PropertyWithNote `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`

		Meta []zCrud.Field `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
)

const (
	RealtorOwnedPropertiesAction = `realtor/ownedProperties`
)

var (
	RealtorOwnedPropertiesMeta = []zCrud.Field{
		{
			Name:      mProperty.SizeM2,
			Label:     `Size (m2)`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.MainUse,
			Label:     `Main Use / Facility`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.MainBuildingMaterial,
			Label:     `Main Building Material`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.ConstructCompletedDate,
			Label:     `Construct Completed Date`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.BuildingLamination,
			Label:     `Building Lamination`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.Address,
			Label:     `Address`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.District,
			Label:     `District`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.About,
			Label:     `About`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.ContactEmail,
			Label:     `ContactEmail`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeEmail,
		},
		{
			Name:      mProperty.ContactPhone,
			Label:     `ContactPhone`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mProperty.CountryCode,
			Label:     `Country`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.CreatedAt,
			Label:     `Created At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mProperty.UpdatedAt,
			Label:     `Last Update`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
	}
)

func (d *Domain) RealtorOwnedProperties(in *RealtorOwnedPropertiesIn) (out RealtorOwnedPropertiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {

		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	out.Properties = r.FindOwnedByPagination(sess.UserId, &in.Pager, &out.Pager)
	if in.ShowMeta {
		out.Meta = RealtorOwnedPropertiesMeta
	}
	return
}
