package domain

import (
	"fmt"
	"street/model/mProperty/rqProperty"
	"street/model/mStorage/rqStorage"

	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserListing.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserListing.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserListing.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserListing.go
//go:generate farify doublequote --file UserListing.go

type (
	UserListingIn struct {
		RequestCommon
		Id uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}

	UserListingOut struct {
		ResponseCommon

		Property *rqProperty.PropertyWithNote `json:"property" form:"property" query:"property" long:"property" msg:"property"`
	}
)

const (
	UserListingAction = `user/listing`
)

func (d *Domain) UserListing(in *UserListingIn) (out UserListingOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	r := rqProperty.NewProperty(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrRealtorPropertyNotFound)
		return
	}
	in.RefId = in.Id
	r.NormalizeFloorList()
	r.Adapter = nil

	propertyWithNote := r.ToPropertyWithNote()
	out.Property = &propertyWithNote


	img3d := rqStorage.NewDesignFiles(d.StorOltp)
	img3dCountryPropId := fmt.Sprintf("%s:%d", out.Property.CountryCode, out.Property.Id)
	img3d.CountryPropId = img3dCountryPropId
	if img3d.FindByCountryPropId() {
		out.Property.Image3dUrl = img3d.FilePath	
	}

	L.Print(img3d)

	return
}
