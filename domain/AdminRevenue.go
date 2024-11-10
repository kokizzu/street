package domain

import (
	"street/model/mBusiness"
	"street/model/mBusiness/rqBusiness"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminRevenue.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminRevenue.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminRevenue.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminRevenue.go
//go:generate farify doublequote --file AdminRevenue.go

type (
	AdminRevenueIn struct {
		RequestCommon
		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Sales rqBusiness.Sales `json:"sales" form:"sales" query:"sales" long:"sales" msg:"sales"`
		PropKey string `json:"propKey" form:"propKey" query:"propKey" long:"propKey" msg:"propKey"`
	}

	AdminRevenueOut struct {
		ResponseCommon
		Revenues []*mBusiness.AdminRevenue `json:"revenues" form:"revenues" query:"revenues" long:"revenues" msg:"revenues"`
	}
)

const (
	AdminRevenueAction = `admin/revenue`
)

func (d *Domain) AdminRevenue(in *AdminRevenueIn) (out AdminRevenueOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Cmd {
	case zCrud.CmdUpsert:
	case zCrud.CmdList:
		r := rqBusiness.NewSales(d.BusinessOltp)
		out.Revenues = r.FindRevenues()
	}

	return
}
