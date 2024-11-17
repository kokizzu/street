package domain

import (
	"street/model/mBusiness"
	"street/model/mBusiness/rqBusiness"
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
		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}

	UserBuyersOut struct {
		ResponseCommon
		Buyers []mBusiness.Buyer `json:"buyers" form:"buyers" query:"buyers" long:"buyers" msg:"buyers"`
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

	switch in.Cmd {
	case zCrud.CmdList:
		r := rqBusiness.NewSales(d.BusinessOltp)
		buyers := r.FindBuyerMonthly(in.YearMonth)

		out.Buyers = buyers
	}


	return
}
