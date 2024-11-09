package domain

import (
	"street/model/mAuth/rqAuth"
	"street/model/mBusiness"
	"street/model/mBusiness/rqBusiness"
	"street/model/mBusiness/wcBusiness"
	"street/model/mProperty/rqProperty"
	"street/model/zCrud"

	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file RealtorRevenue.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type RealtorRevenue.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type RealtorRevenue.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type RealtorRevenue.go
//go:generate farify doublequote --file RealtorRevenue.go

type (
	RealtorRevenueIn struct {
		RequestCommon
		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Sales rqBusiness.Sales `json:"sales" form:"sales" query:"sales" long:"sales" msg:"sales"`
		PropKey string `json:"propKey" form:"propKey" query:"propKey" long:"propKey" msg:"propKey"`
	}
	RealtorRevenueOut struct {
		ResponseCommon
		Revenues []rqBusiness.Revenue `json:"revenues" form:"revenues" query:"revenues" long:"revenues" msg:"revenues"`
	}
)

const (
	RealtorRevenueAction = `realtor/revenue`

	ErrRealtorRevenueBuyerNotFound = `buyer not found to add a new sales`
	ErrRealtorRevenuePropertyNotFound = `property not found to add a new sales`
	ErrRealtorRevenueInvalidPrice = `invalid price, must be a number`
	ErrRealtorRevenueInvalidSalesDate = `invalid sales date, must be in format YYYY-MM-DD`
	ErrRealtorRevenueSaveFailed = `failed to save a new sales`
)

func (d *Domain) RealtorRevenue(in *RealtorRevenueIn) (out RealtorRevenueOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Cmd {
	case zCrud.CmdUpsert:
		sales := wcBusiness.NewSalesMutator(d.BusinessOltp)
		sales.SetRealtorId(sess.UserId)

		switch in.Sales.PropertyCountry {
		case `US`:
			propUS := rqProperty.NewPropertyUS(d.PropOltp)
			propUS.UniqPropKey = in.PropKey
			if !propUS.FindByUniqPropKey() {
				out.SetError(400, ErrRealtorRevenuePropertyNotFound)
				return
			}
			sales.SetPropertyCountry(`US`)
			sales.SetPropertyId(propUS.Id)
		case `TW`:
			propTW := rqProperty.NewPropertyTW(d.PropOltp)
			propTW.Id = in.Sales.PropertyId
			if !propTW.FindById() {
				out.SetError(400, ErrRealtorRevenuePropertyNotFound)
				return
			}
			sales.SetPropertyCountry(`TW`)
			sales.SetPropertyId(propTW.Id)
		default:
			prop := rqProperty.NewProperty(d.PropOltp)
			prop.Id = in.Sales.PropertyId
			if !prop.FindById() {
				out.SetError(400, ErrRealtorRevenuePropertyNotFound)
				return
			}
			sales.SetPropertyCountry(``)
			sales.SetPropertyId(prop.Id)
		}

		if in.Sales.BuyerEmail != `` {
			buyer := rqAuth.NewUsers(d.AuthOltp)
			buyer.Email = in.Sales.BuyerEmail
			if !buyer.FindByEmail() {
				out.SetError(400, ErrRealtorRevenueBuyerNotFound)
				return
			}
			sales.SetBuyerId(buyer.Id)
			sales.SetBuyerEmail(buyer.Email)
		} else {
			sales.SetBuyerId(0)
			sales.SetBuyerEmail(``)
		}

		price := S.ToInt(in.Sales.Price)
		if price == 0 {
			out.SetError(400, ErrRealtorRevenueInvalidPrice)
			return
		}
		sales.SetPrice(in.Sales.Price)

		if !mBusiness.IsValidDate(in.Sales.SalesDate) {
			out.SetError(400, ErrRealtorRevenueInvalidSalesDate)
			return
		}
		sales.SetSalesDate(in.Sales.SalesDate)
		sales.SetCreatedAt(in.UnixNow())
		sales.SetCreatedBy(sess.UserId)
		sales.SetUpdatedAt(in.UnixNow())
		sales.SetUpdatedBy(sess.UserId)

		if !sales.DoInsert() {
			out.SetError(400, ErrRealtorRevenueSaveFailed)
			return
		}
	case zCrud.CmdList:

	}
	
	return
}
