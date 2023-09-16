package domain

import (
	"github.com/kokizzu/gotro/M"

	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminPropHistories.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminPropHistories.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminPropHistories.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminPropHistories.go
//go:generate farify doublequote --file AdminPropHistories.go

type (
	AdminPropHistoriesIn struct {
		RequestCommon

		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`

		PropHistory rqProperty.PropertyHistory `json:"propHistory" form:"propHistory" query:"propHistory" long:"propHistory" msg:"propHistory"`

		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	AdminPropHistoriesOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		PropHistory *rqProperty.PropertyHistory `json:"propHistory" form:"propHistory" query:"propHistory" long:"propHistory" msg:"propHistory"`

		PropHistories [][]any `json:"propHistories" form:"propHistories" query:"propHistories" long:"propHistories" msg:"propHistories"`
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
			Name:      mProperty.TransactionType,
			Label:     `Trx Type`,
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
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mProperty.PricePerUnit,
			Label:     `Price Per Unit`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mProperty.Price,
			Label:     `Price`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
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
	out.refId = in.PropHistory.Id

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminPropHistoriesMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
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
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:

		ph := wcProperty.NewPropertyHistoryMutator(d.PropOltp)
		ph.Id = in.PropHistory.Id
		if ph.Id > 0 {
			if !ph.FindById() {
				out.SetError(400, ErrAdminPropHistoryIdNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if ph.DeletedAt == 0 {
					ph.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if ph.DeletedAt > 0 {
					ph.SetDeletedAt(0)
				}
			}
		} else {
			ph.SetCreatedAt(in.UnixNow())
		}

		haveMutation := ph.SetAll(in.PropHistory, M.SB{}, M.SB{})

		if haveMutation {
			ph.SetUpdatedAt(in.UnixNow())
			ph.SetUpdatedBy(sess.UserId)
			if ph.Id == 0 {
				ph.SetCreatedAt(in.UnixNow())
				ph.SetCreatedBy(sess.UserId)
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
	case zCrud.CmdList:
		ph := rqProperty.NewPropertyHistory(d.PropOltp)
		out.PropHistories = ph.FindByPagination(&AdminPropHistoriesMeta, &in.Pager, &out.Pager)
	}

	return
}
