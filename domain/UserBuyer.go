package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserBuyer.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserBuyer.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserBuyer.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserBuyer.go
//go:generate farify doublequote --file UserBuyer.go

type (
	UserBuyerIn struct {
		RequestCommon
	}

	UserBuyerOut struct {
		ResponseCommon
	}
)

const (
	UserBuyerAction = `user/buyer`
)

func (d *Domain) UserBuyer(in *UserBuyerIn) (out UserBuyerOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	return
}
