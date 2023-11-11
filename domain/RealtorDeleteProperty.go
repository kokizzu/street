package domain

import (
	"street/model/mProperty/wcProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file RealtorDeleteProperty.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type RealtorDeleteProperty.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type RealtorDeleteProperty.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type RealtorDeleteProperty.go
//go:generate farify doublequote --file RealtorDeleteProperty.go

type (
	RealtorDeletePropertyIn struct {
		RequestCommon

		Id uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}
	RealtorDeletePropertyOut struct {
		ResponseCommon

		Ok bool `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
	}
)

const (
	RealtorDeletePropertyAction = `realtor/deleteProperty`

	ErrPropertyToDeleteNotFound   = `realtor property to delete not found`
	ErrPropertyyToDeletedNotOwned = `realtor property to delete not owned`
	ErrFailedToDeleteProperty     = `failed to delete realtor property`
)

func (d *Domain) RealtorDeleteProperty(in *RealtorDeletePropertyIn) (out RealtorDeletePropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	r := wcProperty.NewPropertyMutator(d.PropOltp)
	r.Id = in.Id
	if !r.FindById() {
		out.SetError(400, ErrPropertyToDeleteNotFound)
		return
	}
	if r.CreatedBy != sess.UserId {
		out.SetError(400, ErrPropertyyToDeletedNotOwned)
		return
	}

	r.NormalizeFloorList()
	r.SetDeletedAt(in.UnixNow())
	if !r.DoUpsert() {
		out.SetError(500, ErrFailedToDeleteProperty)
		return
	}

	out.Ok = true

	return
}
