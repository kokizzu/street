package domain

import (
	"street/model/mStorage"
	"street/model/mStorage/rqStorage"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file Admin3DFiles.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type Admin3DFiles.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type Admin3DFiles.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type Admin3DFiles.go
//go:generate farify doublequote --file Admin3DFiles.go

type (
	Admin3DFilesIn struct {
		RequestCommon

		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`

		// for modifying files
		File rqStorage.DesignFiles `json:"file" form:"file" query:"file" long:"file" msg:"file"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	Admin3DFilesOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		File *rqStorage.DesignFiles `json:"file" form:"file" query:"file" long:"file" msg:"file"`

		Files [][]any `json:"files" form:"files" query:"files" long:"files" msg:"files"`
	}
)

const (
	Admin3DFilesAction = `admin/3dFiles`
)

var (
	Admin3DFilesMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mStorage.Id,
				Label:     `ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name: mStorage.FilePath,
				Label: `Filepath`,
				ReadOnly: true,
				DataType: zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name: mStorage.CountryPropId,
				Label: `Country / Property ID`,
				ReadOnly: true,
				DataType: zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
		},
	}
)

func (d *Domain) Admin3DFiles(in *Admin3DFilesIn) (out Admin3DFilesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &Admin3DFilesMeta
	}

	switch in.Cmd {
	case zCrud.CmdList:
		r := rqStorage.NewDesignFiles(d.StorOltp)
		out.Files = r.FindByPagination(&Admin3DFilesMeta, &in.Pager, &out.Pager)
	}

	return
}
