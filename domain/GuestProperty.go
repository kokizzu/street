package domain

import (
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/zCrud"

	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestProperty.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestProperty.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestProperty.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestProperty.go
//go:generate farify doublequote --file GuestProperty.go

type (
	GuestPropertyIn struct {
		RequestCommon
		Id          uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
		CountryCode string `json:"countryCode" form:"countryCode" query:"countryCode" long:"countryCode" msg:"countryCode"`
	}
	GuestPropertyOut struct {
		ResponseCommon
		Property *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
		PropertyExtraUS *rqProperty.PropertyExtraUS `json:"propertyExtraUS" form:"propertyExtraUS" query:"propertyExtraUS" long:"propertyExtraUS" msg:"propertyExtraUS"`
		Meta     []zCrud.Field        `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	GuestPropertyAction             = `guest/property`
	ErrGuestPropertyNotFound        = `property not found`
	ErrGuestPropertyCountryNotFound = `property country not found`
	ErrGuestPropertyExtraUSNotFound = `property extra US not found`
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
			Name: mProperty.FacilityInfoJson,
			Label: `Facility Info`,
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name: `parkingFeatures`,
			Label: `Parking Features`,
			DataType: zCrud.DataTypeString,
		},
		{
			Name: `ListingBrokerName`,
			Label: `Broker Name`,
			DataType: zCrud.DataTypeString,
		},
		{
			Name: `ListingBrokerNumber`,
			Label: `Broker Phone`,
			DataType: zCrud.DataTypeString,
		},
		{
			Name: `ListingAgentName`,
			Label: `Agent Name`,
			DataType: zCrud.DataTypeString,
		},
		{
			Name: `ListingAgentNumber`,
			Label: `Agent Phone`,
			DataType: zCrud.DataTypeString,
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

func (d *Domain) GuestProperty(in *GuestPropertyIn) (out GuestPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	if in.CountryCode == `US` {
		r := rqProperty.NewPropertyUS(d.PropOltp)
		r.Id = in.Id
		if !r.FindById() {
			out.SetError(400, ErrGuestPropertyCountryNotFound)
			return
		}
		out.Property = r.ToProperty()

		rx := rqProperty.NewPropertyExtraUS(d.PropOltp)
		rx.PropertyKey = r.UniqPropKey
		if !rx.FindByPropertyKey() {
			L.Print(ErrGuestPropertyExtraUSNotFound)
		}
		out.PropertyExtraUS = rx.ToPropertyExtra()
		out.Meta = GuestPropertiesMeta
		return
	}

	if in.CountryCode == `TW` {
		r := rqProperty.NewPropertyTW(d.PropOltp)
		r.Id = in.Id
		if !r.FindById() {
			out.SetError(400, ErrGuestPropertyCountryNotFound)
			return
		}
		out.Property = r.ToProperty()
		out.Meta = GuestPropertiesMeta
		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrGuestPropertyNotFound)
		return
	}

	if r.ApprovalState != mProperty.ApprovalStatePending && r.ApprovalState != `` {
		out.SetError(400, ErrGuestPropertyNotFound)
		return
	}
	r.NormalizeFloorList()
	out.Property = r
	out.Meta = GuestPropertiesMeta
	return
}
