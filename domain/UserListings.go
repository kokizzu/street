package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserListings.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserListings.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserListings.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserListings.go
//go:generate farify doublequote --file UserListings.go

type (
	UserListingsIn struct {
		RequestCommon
	}

	UserListingsOut struct {
		ResponseCommon
	}
)

const (
	UserListingsAction = `user/listings`
)

func (d *Domain) UserListings(in *UserListingsIn) (out UserListingsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	return
}
