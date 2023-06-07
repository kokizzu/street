package domain

import (
	"street/model/mAuth/rqAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file admin.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type admin.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type admin.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type admin.go
//go:generate farify doublequote --file admin.go

type (
	AdminUserCrudIn struct {
		RequestCommon

		Action string `json:"action" form:"action" query:"action" long:"action" msg:"action"`

		zCrud.PaginationIn
	}
	AdminUserCrudOut struct {
		ResponseCommon

		zCrud.PaginationOut

		Users [][]any `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	AdminUserListAction = `admin/userList`
)

func (d *Domain) AdminUserList(in *AdminUserCrudIn) (out AdminUserCrudOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Action {
	case zCrud.ActionForm:

	case zCrud.ActionUpsert, zCrud.ActionDelete, zCrud.ActionRestore:

		fallthrough
	case zCrud.ActionList:
		r := rqAuth.NewUsers(d.AuthOltp)

		// TODO: return all columns meta/schema
		out.Users = r.FindByPagination(&in.PaginationIn, &out.PaginationOut)

		// TODO: return [][]any based on order

	}

	return
}
