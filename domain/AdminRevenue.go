package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminRevenue.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminRevenue.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminRevenue.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminRevenue.go
//go:generate farify doublequote --file AdminRevenue.go

type (
	AdminRevenueIn struct {
		RequestCommon
	}

	AdminRevenueOut struct {
		ResponseCommon
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

	return
}
