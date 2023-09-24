package domain

import (
	"fmt"
	"os"

	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/zCrud"
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
		Property      *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
		Meta          []zCrud.Field        `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		OgURL         string               `json:"og_url" form:"og_url" query:"og_url" long:"og_url" msg:"og_url"`
		OgImgURL      string               `json:"og_imgUrl" form:"og_imgUrl" query:"og_imgUrl" long:"og_imgUrl" msg:"og_imgUrl"`
		OgDescription string               `json:"og_description" form:"og_description" query:"og_description" long:"og_description" msg:"og_description"`
		OgCreatedAt   int64                `json:"og_createdAt" form:"og_createdAt" query:"og_createdAt" long:"og_createdAt" msg:"og_createdAt"`
		OgUpdatedAt   int64                `json:"og_updatedAt" form:"og_updatedAt" query:"og_updatedAt" long:"og_updatedAt" msg:"og_updatedAt"`
	}
)

const (
	GuestPropertyAction      = `guest/property`
	ErrGuestPropertyNotFound = `Property not found`
)

var (
	GuestPropertiesMeta = []zCrud.Field{
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
			Name:      mProperty.Note,
			Label:     `Note`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.Country,
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

func (d *Domain) GuestProperty(in *GuestPropertyIn) (out GuestPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	domainName := os.Getenv("WEB_PROTO_DOMAIN")
	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrGuestPropertyNotFound)
		return
	}
	r.NormalizeFloorList()
	out.Property = r
	out.Meta = GuestPropertiesMeta
	out.OgURL = fmt.Sprintf("%s/%s/%d", domainName, GuestPropertyAction, r.Id)
	out.OgImgURL = fmt.Sprintf("%s%s", domainName, r.Images[0])
	out.OgDescription = r.Note
	out.OgCreatedAt = r.CreatedAt
	out.OgUpdatedAt = r.UpdatedAt
	return
}
