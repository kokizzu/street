package domain

import (
	"fmt"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"

	"street/model/mAuth/rqAuth"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminPropertiesTW.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminPropertiesTW.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminPropertiesTW.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminPropertiesTW.go
//go:generate farify doublequote --file AdminPropertiesTW.go

type (
	AdminPropertiesTWIn struct {
		RequestCommon

		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`

		// for modifying property
		Property rqProperty.PropertyTW `json:"property" form:"property" query:"property" long:"property" msg:"property"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	AdminPropertiesTWOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		Property *rqProperty.PropertyTW `json:"property" form:"property" query:"property" long:"property" msg:"property"`

		// listing
		Properties [][]any `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`
	}
)

const (
	AdminPropertiesTWAction = `admin/properties/TW`

	ErrAdminPropertyTWIdNotFound = `propertyTW id not found`
	ErrAdminPropertyTWSaveFailed = `propertyTW save failed`
)

var (
	AdminPropertiesTWMeta = zCrud.Meta{
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
				Name:      mProperty.SizeM2,
				Label:     `Size (m2)`,
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
				Name:      mProperty.Address,
				Label:     `Address`,
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
				InputType: zCrud.InputTypeText,
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
			//	Name:      mProperty.Images,
			//	Label:     `Images`,
			//	DataType:  zCrud.DataTypeString,
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
				Name:      mProperty.DeletedAt,
				Label:     `Deleted At`,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mProperty.ApprovalState,
				Label:     `Approval State`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
		},
	}
)

func (d *Domain) AdminPropertiesTW(in *AdminPropertiesTWIn) (out AdminPropertiesTWOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.Property.Id

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminPropertiesTWMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.Property.Id <= 0 {
			out.Meta = &AdminPropertiesTWMeta
			return
		}

		prop := rqProperty.NewPropertyTW(d.PropOltp)
		prop.Id = in.Property.Id
		if !prop.FindById() {
			out.SetError(400, ErrAdminPropertyTWIdNotFound)
			return
		}
		// prop.NormalizeFloorList()
		out.Property = prop
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		prop := wcProperty.NewPropertyTWMutator(d.PropOltp)
		prop.Id = in.Property.Id
		if prop.Id > 0 {
			if !prop.FindById() {
				out.SetError(400, ErrAdminPropertyTWIdNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if prop.DeletedAt == 0 {
					prop.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if prop.DeletedAt > 0 {
					prop.SetDeletedAt(0)
				}
			}
		} else {
			prop.SetCreatedAt(in.UnixNow())
		}

		oldState := prop.ApprovalState
		newState := S.Trim(in.Property.ApprovalState)
		haveMutation := prop.SetAll(in.Property, M.SB{
			mProperty.PriceHistoriesSell: true,
			mProperty.PriceHistoriesRent: true,
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
			out.SetError(500, ErrAdminPropertyTWSaveFailed)
			break
		}

		prop.Adapter = nil
		out.Property = &prop.PropertyTW

		var sendMailFunc func(email, number, lang, link string) error
		var sendMailName string

		if newState == `` && oldState != `` {
			out.AddTrace(`state:accepted`)
			sendMailFunc = d.Mailer.SendNotifPropertyAcceptedEmail
			sendMailName = `SendNotifPropertyAcceptedEmail`
		} else if newState != mProperty.ApprovalStatePending {
			out.AddTrace(`state:rejected`)
			sendMailFunc = d.Mailer.SendNotifPropertyRejectedEmail
			sendMailName = `SendNotifPropertyRejectedEmail`
		}

		if sendMailFunc != nil {
			user := rqAuth.NewUsers(d.AuthOltp)
			user.Id = prop.CreatedBy
			if user.FindById() {
				d.runSubtask(func() {
					err := sendMailFunc(user.Email,
						fmt.Sprintf(`#%d`, prop.Id),
						user.Language,
						fmt.Sprintf("%s/realtor/property/%v", d.WebCfg.WebProtoDomain, in.Property.Id),
					)
					L.IsError(err, sendMailName)
				})
			} else {
				out.AddTrace(`failFindRealtor:` + I.UToS(user.Id))
				// continue anyway if failed to send email
			}
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqProperty.NewPropertyTW(d.PropOltp)
		out.Properties = r.FindByPagination(&AdminPropertiesTWMeta, &in.Pager, &out.Pager)
	}

	return
}
