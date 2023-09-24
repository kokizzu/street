package domain

import (
	"street/model/mAuth"
	"street/model/mAuth/saAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminAccessLogs.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminAccessLogs.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminAccessLogs.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminAccessLogs.go
//go:generate farify doublequote --file AdminAccessLogs.go

type (
	AdminAccessLogsIn struct {
		RequestCommon

		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
	}

	AdminAccessLogsOut struct {
		ResponseCommon

		Pager zCrud.PagerOut      `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Logs  []saAuth.ActionLogs `json:"logs" form:"logs" query:"logs" long:"logs" msg:"logs"`
		Meta  *zCrud.Meta         `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	AdminAccessLogsAction = `admin/accessLogs`
)

var (
	AdminAccessLogsMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mAuth.CreatedAt,
				Label:     `Created At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mAuth.RequestId,
				Label:     `Request ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.ActorId,
				Label:     `Actor ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Action,
				Label:     `Action`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.StatusCode,
				Label:     `Status Code`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Traces,
				Label:     `Traces`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.Error,
				Label:     `Error`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.IpAddr4,
				Label:     `IP Address 4`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.IpAddr6,
				Label:     `IP Address 6`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.UserAgent,
				Label:     `User Agent`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.Lat,
				Label:     `Latitude`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeFloat,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Long,
				Label:     `Longitude`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeFloat,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Latency,
				Label:     `Latency`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeFloat,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.RefId,
				Label:     `Ref ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
		},
	}
)

func (d *Domain) AdminAccessLogs(in *AdminAccessLogsIn) (out AdminAccessLogsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminAccessLogsMeta
	}

	// if not set, always override by createdAt descending
	if len(in.Pager.Order) == 0 {
		in.Pager.Order = []string{`-createdAt`}
	}

	r := saAuth.NewActionLogs(d.AuthOlap)
	out.Logs = r.FindByPagination(&in.Pager, &out.Pager)

	return
}
