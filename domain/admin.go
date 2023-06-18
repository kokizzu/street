package domain

import (
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file admin.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type admin.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type admin.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type admin.go
//go:generate farify doublequote --file admin.go

type (
	AdminUsersIn struct {
		RequestCommon

		Action string `json:"action" form:"action" query:"action" long:"action" msg:"action"`

		// for modifying user
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		zCrud.PagerIn
	}
	AdminUsersOut struct {
		ResponseCommon

		zCrud.PagerOut

		Meta []zCrud.Field

		// modified user or form user
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		// listing
		Users [][]any `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	AdminUsersAction = `admin/users`

	ErrAdminUserIdNotFound = `user id not found`
	ErrAdminUserSaveFailed = `user save failed`
)

func (d *Domain) AdminUsers(in *AdminUsersIn) (out AdminUsersOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.mustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	fillMeta := false

	switch in.Action {
	case zCrud.ActionForm:
		fillMeta = true

		user := rqAuth.NewUsers(d.AuthOltp)
		user.Id = in.User.Id
		if !user.FindById() {
			out.SetError(400, ErrAdminUserIdNotFound)
		}
	case zCrud.ActionUpsert, zCrud.ActionDelete, zCrud.ActionRestore:

		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Id = in.User.Id
		if user.Id > 0 {
			user.FindById()
		} else {
			user.SetCreatedAt(in.UnixNow())
		}

		// admin can override validation
		user.SetUserName(in.User.UserName)
		user.SetEmail(in.User.Email)
		user.SetFullName(in.User.FullName)

		if !user.DoUpsert() {
			out.SetError(500, ErrAdminUserSaveFailed)
			break
		}

		fallthrough
	case zCrud.ActionList:
		fillMeta = true
		r := rqAuth.NewUsers(d.AuthOltp)
		out.Users = r.FindByPagination(&in.PagerIn, &out.PagerOut)
	}

	if fillMeta {
		out.Meta = []zCrud.Field{
			{
				Name:      `id`,
				Label:     `ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      `userName`,
				Label:     `Username`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      `email`,
				Label:     `Email`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeEmail,
			},
			{
				Name:      `fullName`,
				Label:     `Full Name`,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      `createdAt`,
				Label:     `Created At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      `updatedAt`,
				Label:     `Updated At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      `deletedAt`,
				Label:     `Deleted At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      `verifiedAt`,
				Label:     `Verified At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      `lastLoginAt`,
				Label:     `Last Login At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
		}
	}

	return
}
