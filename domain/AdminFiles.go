package domain

import (
	"street/model/mStorage"
	"street/model/mStorage/rqStorage"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminFiles.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminFiles.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminFiles.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminFiles.go
//go:generate farify doublequote --file AdminFiles.go

type (
	AdminFilesIn struct {
		RequestCommon

		Action string `json:"action" form:"action" query:"action" long:"action" msg:"action"`

		// for modifying files
		File rqStorage.Files `json:"file" form:"file" query:"file" long:"file" msg:"file"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	AdminFilesOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		File *rqStorage.Files `json:"file" form:"file" query:"file" long:"file" msg:"file"`

		Files [][]any `json:"files" form:"files" query:"files" long:"files" msg:"files"`
	}
)

const (
	AdminFilesAction = `admin/files`
)

var (
	AdminFilesMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mStorage.Id,
				Label:     `ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.OriginalPath,
				Label:     `Original Path`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.ResizedPath,
				Label:     `Resized Path`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.OriginalSize,
				Label:     `Size`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.Mime,
				Label:     `MIME`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.Purpose,
				Label:     `Purpose`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.RefId,
				Label:     `Ref ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.AccessCount,
				Label:     `Access Count`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mStorage.LastAccessAt,
				Label:     `Last Access`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
		},
	}
)

func (d *Domain) AdminFiles(in *AdminFilesIn) (out AdminFilesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminPropertiesMeta
	}

	switch in.Action {
	case zCrud.ActionList:
		r := rqStorage.NewFiles(d.StorOltp)
		out.Files = r.FindByPagination(&AdminFilesMeta, &in.Pager, &out.Pager)

	}

	return
}
