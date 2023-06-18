package domain

import (
	"github.com/kokizzu/gotro/L"

	"street/model/mAuth/rqAuth"
	"street/model/mAuth/saAuth"
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

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn
	}
	AdminUsersOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		// modified user or form user
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		// listing
		Users [][]any `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	AdminUsersAction = `admin/users`

	ErrAdminUserIdNotFound      = `user id not found`
	ErrAdminUserSaveFailed      = `user save failed`
	ErrAdminUsersEmailDuplicate = `email already by another user`
	ErrAdminUsernameDuplicate   = `username already user by another user`
)

var AdminUsersMeta = zCrud.Meta{
	Fields: []zCrud.Field{
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
	},
}

func (d *Domain) AdminUsers(in *AdminUsersIn) (out AdminUsersOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminUsersMeta
	}

	switch in.Action {
	case zCrud.ActionForm:
		if in.User.Id <= 0 {
			out.Meta = &AdminUsersMeta
			return
		}

		user := rqAuth.NewUsers(d.AuthOltp)
		user.Id = in.User.Id
		if !user.FindById() {
			out.SetError(400, ErrAdminUserIdNotFound)
		}
		user.CensorFields()
		out.User = user
	case zCrud.ActionUpsert, zCrud.ActionDelete, zCrud.ActionRestore:

		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Id = in.User.Id
		if user.Id > 0 {
			if !user.FindById() {
				out.SetError(400, ErrAdminUserIdNotFound)
				return
			}
		} else {
			user.SetCreatedAt(in.UnixNow())
		}

		// admin can override validation
		if user.SetUserName(in.User.UserName) {
			dup := rqAuth.NewUsers(d.AuthOltp)
			dup.UserName = in.User.UserName
			if dup.FindByUserName() && dup.Id != user.Id {
				out.SetError(400, ErrAdminUsernameDuplicate)
				return
			}
		}
		if user.SetEmail(in.User.Email) {
			user.SetVerifiedAt(0)
			dup := rqAuth.NewUsers(d.AuthOltp)
			dup.Email = in.User.Email
			if dup.FindByEmail() && dup.Id != user.Id {
				out.SetError(400, ErrAdminUsersEmailDuplicate)
				return
			}
		}
		user.SetFullName(in.User.FullName)

		if user.HaveMutation() {
			user.SetUpdatedAt(in.UnixNow())
			user.SetUpdatedBy(sess.UserId)
		}
		L.Describe(user)

		if !user.DoUpsert() {
			out.SetError(500, ErrAdminUserSaveFailed)
			break
		}

		user.CensorFields()
		out.User = &user.Users

		if in.Pager.Page == 0 {
			break
		}
		fallthrough
	case zCrud.ActionList:
		r := rqAuth.NewUsers(d.AuthOltp)
		out.Users = r.FindByPagination(&AdminUsersMeta, &in.Pager, &out.Pager)
	}

	return
}

type (
	AdminDashboardIn struct {
		RequestCommon
	}

	AdminDashboardOut struct {
		ResponseCommon

		RegisteredUserTotal int64

		RequestsPerDate   map[string]int
		UniqueUserPerDate map[string]int
		UniqueIpPerDate   map[string]int

		CountPerActionsPerDate map[string]map[string]int
	}
)

const (
	AdminDashboardAction = `admin/dashboard`
)

func (d *Domain) AdminDashboard(in *AdminDashboardIn) (out AdminDashboardOut) {

	rq := rqAuth.NewUsers(d.AuthOltp)
	out.RegisteredUserTotal = rq.Total()

	sa := saAuth.NewActionLogs(d.AuthOlap)
	out.RequestsPerDate = sa.StatRequestsPerDate()
	out.UniqueUserPerDate = sa.StatUniqueUserPerDate()
	out.UniqueIpPerDate = sa.StatUniqueIpPerDate()

	out.CountPerActionsPerDate = sa.StatPerActionsPerDate()
	return
}
