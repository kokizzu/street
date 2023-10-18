package domain

import (
	"strconv"
	"strings"

	"github.com/kokizzu/gotro/L"

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
		Id string `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}
	UserPropertyOut struct {
		ResponseCommon
		Property    *rqProperty.Property        `json:"property" form:"property" query:"property" long:"property" msg:"property"`
		PropertyUS  *rqProperty.PropertyUS      `json:"property" form:"property" query:"property" long:"property" msg:"property"`
		PropHistory *rqProperty.PropertyHistory `json:"propHistory" form:"propHistory" query:"propHistory" long:"propHistory" msg:"propHistory"`
		Meta        []zCrud.Field               `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	UserPropertyAction           = `user/property`
	ErrUserPropertyNotFound      = `Property not found`
	ErrUserPropHistoryIdNotFound = `user prop history id not found`
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

func (d *Domain) UserProperty(in *GuestPropertyIn) (out UserPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}
	idSplit := strings.Split(in.Id, "US")
	if len(idSplit) == 2 {
		idUint, _ := strconv.ParseUint(idSplit[1], 10, 64)
		r := rqProperty.NewPropertyUS(d.PropOltp)
		r.Id = idUint
		if !r.FindById() {
			out.SetError(400, ErrUserPropertyNotFound)
			return
		}
		out.PropertyUS = r
		out.Meta = UserPropertiesMeta
		return
	}
	idUint, err := strconv.ParseUint(in.Id, 10, 64)
	if err != nil {
		out.SetError(400, ErrUserPropertyNotFound)
		return
	}

	ph := rqProperty.NewPropertyHistory(d.PropOltp)
	ph.Id = idUint
	out.PropHistory = ph
	L.Print(ph)
	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = idUint
	if !r.FindById() {
		out.SetError(400, ErrUserPropertyNotFound)
		return
	}
	r.NormalizeFloorList()
	out.Property = r
	out.Meta = UserPropertiesMeta
	return
}
