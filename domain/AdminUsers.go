package domain

import (
	"street/model/mAuth"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminUsers.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminUsers.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminUsers.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminUsers.go
//go:generate farify doublequote --file AdminUsers.go

type (
	AdminUsersIn struct {
		RequestCommon

		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`

		// for modifying user
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
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
			Name:      mAuth.Id,
			Label:     `ID`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
		},
		{
			Name:      mAuth.UserName,
			Label:     `Username`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mAuth.Email,
			Label:     `Email`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeEmail,
		},
		{
			Name:      mAuth.FullName,
			Label:     `Full Name`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mAuth.CreatedAt,
			Label:     `Created At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.UpdatedAt,
			Label:     `Updated At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.DeletedAt,
			Label:     `Deleted At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.VerifiedAt,
			Label:     `Verified At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.Country,
			Label:     `Country`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mAuth.LastLoginAt,
			Label:     `Last Login At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
	},
}

func (d *Domain) AdminUsers(in *AdminUsersIn) (out AdminUsersOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.User.Id

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminUsersMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
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
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:

		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Id = in.User.Id
		if user.Id > 0 {
			if !user.FindById() {
				out.SetError(400, ErrAdminUserIdNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if user.DeletedAt == 0 {
					user.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if user.DeletedAt > 0 {
					user.SetDeletedAt(0)
				}
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
			if user.Id == 0 {
				user.SetCreatedAt(in.UnixNow())
			}
		}
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
	case zCrud.CmdList:
		r := rqAuth.NewUsers(d.AuthOltp)
		out.Users = r.FindByPagination(&AdminUsersMeta, &in.Pager, &out.Pager)
	}

	return
}
