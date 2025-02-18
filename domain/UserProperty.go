package domain

import (
	"github.com/kokizzu/gotro/S"

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
	UserPropertyIn struct {
		RequestCommon
		Id          uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
		CountryCode string
	}
	UserPropertyOut struct {
		ResponseCommon
		Property      *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
		PropHistories []*rqProperty.PropertyHistory
		Meta          []zCrud.Field `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	UserPropertyAction      = `user/property`
	ErrUserPropertyNotFound = `user property not found`
)

var (
	UserPropertiesMeta = []zCrud.Field{
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
			Name:      mProperty.Attribute,
			Label:     `Attribute`,
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

func (d *Domain) UserProperty(in *UserPropertyIn) (out UserPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.CountryCode == `US` { // for now
		r := rqProperty.NewPropertyUS(d.PropOltp)
		out.Property = r.ToProperty()
		out.Meta = UserPropertiesMeta
		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrUserPropertyNotFound)
		return
	}
	in.RefId = in.Id
	r.NormalizeFloorList()
	out.Property = r
	out.Meta = UserPropertiesMeta

	ph := rqProperty.NewPropertyHistory(d.PropOltp)
	serialNumber := S.LeftOf(r.UniqPropKey, `#`)
	hist := ph.FindBySerialNumber(serialNumber)
	out.PropHistories = hist

	return
}
