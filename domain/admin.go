package domain

import (
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file admin.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type admin.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type admin.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type admin.go
//go:generate farify doublequote --file admin.go

type (
	AdminUserListIn struct {
		RequestCommon

		zCrud.PaginationIn
	}
	AdminUserListOut struct {
		ResponseCommon

		zCrud.PaginationOut
	}
)

const (
	AdminUserListAction = `admin/userList`
)

func (d *Domain) AdminUserList(in *AdminUserListIn) (out AdminUserListOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	// TODO: query user list based on specific order

	return
}
