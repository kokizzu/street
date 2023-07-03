package domain

import (
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zCrud"
)

type (
	AdminPropertiesIn struct {
		RequestCommon

		Action string `json:"action" form:"action" query:"action" long:"action" msg:"action"`

		// for modifying user
		Property rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn
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
				Label:     `Main Use`,
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
		},
	}
)

func (d *Domain) AdminProperties(in *AdminPropertiesIn) (out AdminPropertiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
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
		}
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

		prop.SetUniqPropKey(in.Property.UniqPropKey)
		prop.SetSerialNumber(in.Property.SerialNumber)
		prop.SetSizeM2(in.Property.SizeM2)
		prop.SetMainUse(in.Property.MainUse)
		prop.SetMainBuildingMaterial(in.Property.MainBuildingMaterial)
		prop.SetConstructCompletedDate(in.Property.ConstructCompletedDate)
		prop.SetNumberOfFloors(in.Property.NumberOfFloors)
		prop.SetBuildingLamination(in.Property.BuildingLamination)
		prop.SetAddress(in.Property.Address)
		prop.SetDistrict(in.Property.District)
		prop.SetNote(in.Property.Note)
		prop.SetCoord(in.Property.Coord)

		if prop.HaveMutation() {
			prop.SetUpdatedAt(in.UnixNow())
			prop.SetUpdatedBy(sess.UserId)
			if prop.Id == 0 {
				prop.SetCreatedAt(in.UnixNow())
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

type (
	AdminPropHistoriesIn struct {
		RequestCommon

		Action string `json:"action"`

		PropHistory *rqProperty.PropertyHistory `json:"propertyHistory"`

		WithMeta bool `json:"withMeta"`

		Pager zCrud.PagerIn `json:"pager"`
	}

	AdminPropHistoriesOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		PropHistory *rqProperty.PropertyHistory

		PropHistories [][]any
	}
)

const (
	AdminPropHistoriesAction = `admin/propHistories`

	ErrAdminPropHistoryIdNotFound = `admin prop history id not found`
	ErrAdminPropHistorySaveFailed = `admin prop history save failed`
)

var AdminPropHistoriesMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `Id`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
		},
		{
			Name:      mProperty.PropertyKey,
			Label:     `Prop Key`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.TransactionKey,
			Label:     `Trx Key`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.TransactionSign,
			Label:     `Trx Sign`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.TransactionTime,
			Label:     `Trx Time`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.TransactionDateNormal,
			Label:     `Trx Date Normal`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.TransactionNumber,
			Label:     `Trx Number`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.PriceNTD,
			Label:     `Price NTD`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.PricePerUnit,
			Label:     `Price Per Unit`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mProperty.Price,
			Label:     `Price`,
			DataType:  zCrud.DataTypeInt,
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
	},
}

func (d *Domain) AdminPropHistories(in *AdminPropHistoriesIn) (out AdminPropHistoriesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminPropHistoriesMeta
	}

	switch in.Action {
	case zCrud.ActionForm:
		if in.PropHistory.Id <= 0 {
			out.Meta = &AdminPropertiesMeta
			return
		}

		ph := rqProperty.NewPropertyHistory(d.PropOltp)
		ph.Id = in.PropHistory.Id
		if !ph.FindById() {
			out.SetError(400, ErrAdminPropHistoryIdNotFound)
		}
		out.PropHistory = ph
	case zCrud.ActionUpsert, zCrud.ActionDelete, zCrud.ActionRestore:

		ph := wcProperty.NewPropertyHistoryMutator(d.PropOltp)
		ph.Id = in.PropHistory.Id
		if ph.Id > 0 {
			if !ph.FindById() {
				out.SetError(400, ErrAdminPropHistoryIdNotFound)
				return
			}

			if in.Action == zCrud.ActionDelete {
				if ph.DeletedAt == 0 {
					ph.SetDeletedAt(in.UnixNow())
				}
			} else if in.Action == zCrud.ActionRestore {
				if ph.DeletedAt > 0 {
					ph.SetDeletedAt(0)
				}
			}
		} else {
			ph.SetCreatedAt(in.UnixNow())
		}

		ph.SetPropertyKey(in.PropHistory.PropertyKey)
		ph.SetTransactionKey(in.PropHistory.TransactionKey)
		ph.SetTransactionSign(in.PropHistory.TransactionSign)
		ph.SetTransactionTime(in.PropHistory.TransactionTime)
		ph.SetTransactionDateNormal(in.PropHistory.TransactionDateNormal)
		ph.SetTransactionNumber(in.PropHistory.TransactionNumber)
		ph.SetPriceNtd(in.PropHistory.PriceNtd)
		ph.SetPricePerUnit(in.PropHistory.PricePerUnit)
		ph.SetPrice(in.PropHistory.Price)
		ph.SetAddress(in.PropHistory.Address)
		ph.SetDistrict(in.PropHistory.District)
		ph.SetNote(in.PropHistory.Note)

		if ph.HaveMutation() {
			ph.SetUpdatedAt(in.UnixNow())
			ph.SetUpdatedBy(sess.UserId)
			if ph.Id == 0 {
				ph.SetCreatedAt(in.UnixNow())
			}
		}
		if !ph.DoUpsert() {
			out.SetError(500, ErrAdminPropHistorySaveFailed)
			break
		}

		out.PropHistory = &ph.PropertyHistory

		if in.Pager.Page == 0 {
			break
		}
		fallthrough
	case zCrud.ActionList:
		r := rqProperty.NewPropertyHistory(d.PropOltp)
		out.PropHistories = r.FindByPagination(&AdminPropHistoriesMeta, &in.Pager, &out.Pager)
	}

	return
}
