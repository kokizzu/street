package domain

import (
	"github.com/kokizzu/gotro/M"

	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminProperties.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminProperties.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminProperties.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminProperties.go
//go:generate farify doublequote --file AdminProperties.go

type (
	AdminPropertiesIn struct {
		RequestCommon

		Cmd string

		// for modifying property
		Property rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	AdminPropertiesOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		Property *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`

		// listing
		Properties [][]any `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`
	}
)

const (
	AdminPropertiesAction = `admin/properties`

	ErrAdminPropertyIdNotFound = `property id not found`
	ErrAdminPropertySaveFailed = `property save failed`
)

var (
	AdminPropertiesMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mProperty.Id,
				Label:     `ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mProperty.UniqPropKey,
				Label:     `Prop Key`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mProperty.SerialNumber,
				Label:     `Serial Number`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
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
				Name:      mProperty.NumberOfFloors,
				Label:     `Number Of Floors`,
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
				Name:      mProperty.Coord,
				Label:     `Coord`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mProperty.LastPrice,
				Label:     `Last Price`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText, // TODO: need to fix schema
			},
			{
				Name:      mProperty.Purpose, // rent/sell
				Label:     `Purpose`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText, // TODO: change to select
			},
			{
				Name:      mProperty.HouseType, // house/apartment
				Label:     `House Type`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText, // TODO: change to select
			},
			//{
			//	Name: 	mProperty.Images,
			//	Label: 	`Images`,
			//	DataType: zCrud.DataTypeString,
			//	InputType: zCrud.InputTypeText, // TODO: change to image upload
			//},
			{
				Name:      mProperty.Bedroom,
				Label:     `Bedroom`,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mProperty.Bathroom,
				Label:     `Bathroom`,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mProperty.AgencyFeePercent,
				Label:     `Agency Fee Percent`,
				DataType:  zCrud.DataTypeFloat,
				InputType: zCrud.InputTypeNumber,
			},
			//{
			//	Name:      mProperty.FloorList,
			//	Label:     `Floor List`,
			//	DataType:  zCrud.DataTypeString,
			//	InputType: zCrud.InputTypeText, // TODO: change to json editor
			//},
			{
				Name:      mProperty.Country,
				Label:     `Country`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText, // TODO: change to select
			},
			{
				Name:      mProperty.DeletedAt,
				Label:     `Deleted At`,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
		},
	}
)

func (d *Domain) AdminProperties(in *AdminPropertiesIn) (out AdminPropertiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.Property.Id

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminPropertiesMeta
	}

	switch in.Action {
	case zCrud.ActionForm:
		if in.Property.Id <= 0 {
			out.Meta = &AdminPropertiesMeta
			return
		}

		prop := rqProperty.NewProperty(d.PropOltp)
		prop.Id = in.Property.Id
		if !prop.FindById() {
			out.SetError(400, ErrAdminPropertyIdNotFound)
			return
		}
		prop.NormalizeFloorList()
		out.Property = prop
	case zCrud.ActionUpsert, zCrud.ActionDelete, zCrud.ActionRestore:

		prop := wcProperty.NewPropertyMutator(d.PropOltp)
		prop.Id = in.Property.Id
		if prop.Id > 0 {
			if !prop.FindById() {
				out.SetError(400, ErrAdminPropertyIdNotFound)
				return
			}

			if in.Action == zCrud.ActionDelete {
				if prop.DeletedAt == 0 {
					prop.SetDeletedAt(in.UnixNow())
				}
			} else if in.Action == zCrud.ActionRestore {
				if prop.DeletedAt > 0 {
					prop.SetDeletedAt(0)
				}
			}
		} else {
			prop.SetCreatedAt(in.UnixNow())
		}

		haveMutation := prop.SetAll(in.Property, M.SB{
			`priceHistoriesSell`: true,
			`priceHistoriesRent`: true,
		}, M.SB{})

		if haveMutation {
			prop.SetUpdatedAt(in.UnixNow())
			prop.SetUpdatedBy(sess.UserId)
			if prop.Id == 0 {
				prop.SetCreatedAt(in.UnixNow())
				prop.SetCreatedBy(sess.UserId)
			}
		}
		if !prop.DoUpsert() {
			out.SetError(500, ErrAdminPropertySaveFailed)
			break
		}

		out.Property = &prop.Property

		if in.Pager.Page == 0 {
			break
		}
		fallthrough
	case zCrud.ActionList:
		r := rqProperty.NewProperty(d.PropOltp)
		out.Properties = r.FindByPagination(&AdminPropertiesMeta, &in.Pager, &out.Pager)
	}

	return
}
