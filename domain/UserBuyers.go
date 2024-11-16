package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserBuyers.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserBuyers.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserBuyers.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserBuyers.go
//go:generate farify doublequote --file UserBuyers.go

type (
	UserBuyersIn struct {
		RequestCommon
	}

	UserBuyersOut struct {
		ResponseCommon
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

	return
}
