package domain

import (
	"street/model/mAuth/rqAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserBuyers.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserBuyers.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserBuyers.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserBuyers.go
//go:generate farify doublequote --file UserBuyers.go

type (
	UserBuyersIn struct {
		RequestCommon
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	UserBuyersOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Users [][]any        `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	UserBuyersAction = `user/buyers`
)

func (d *Domain) UserBuyers(in *UserBuyersIn) (out UserBuyersOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminUsersMeta
	}

	switch in.Cmd {
	case zCrud.CmdList:
		r := rqAuth.NewUsers(d.AuthOltp)
		out.Users = r.FindByPagination(&AdminUsersMeta, &in.Pager, &out.Pager)
	}

	return
}
