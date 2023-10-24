package domain

import (
	"fmt"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/segmentio/fasthash/fnv1a"

	"street/conf"
	"street/model/mAuth/rqAuth"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file RealtorUpsertProperty.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type RealtorUpsertProperty.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type RealtorUpsertProperty.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type RealtorUpsertProperty.go
//go:generate farify doublequote --file RealtorUpsertProperty.go

type (
	RealtorUpsertPropertyIn struct {
		RequestCommon

		Property rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
	}
	RealtorUpsertPropertyOut struct {
		ResponseCommon

		Property *rqProperty.Property `json:"property" form:"property" query:"property" long:"property" msg:"property"`
	}
)

const (
	RealtorUpsertPropertyAction = `realtor/upsertProperty`

	ErrRealtorUpsertPropertyNotExists           = `property id not exists`
	ErrRealtorUpsertPropertyCoordInvalid        = `property coordinate invalid`
	ErrRealtorUpsertPropertyEmptyImage          = `property image empty`
	ErrRealtorUpsertPropertyAddressAlreadyAdded = `property address already added by you in the past`
	ErrRealtorUpsertPropertyNotOwnedByYou       = `property not owned by you`
	ErrRealtorUpsertPropertySaveFailed          = `property save failed`
)

func (d *Domain) RealtorUpsertProperty(in *RealtorUpsertPropertyIn) (out RealtorUpsertPropertyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	// all logged-in user is potential realtor
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if len(in.Property.Coord) != 2 {
		out.SetError(400, ErrRealtorUpsertPropertyCoordInvalid)
		return
	}

	if len(in.Property.Images) == 0 {
		out.SetError(400, ErrRealtorUpsertPropertyEmptyImage)
	}

	prop := wcProperty.NewPropertyMutator(d.PropOltp)

	if in.Property.Id > 0 {
		prop.Id = in.Property.Id
		if !prop.FindById() {
			out.SetError(400, ErrRealtorUpsertPropertyNotExists)
			return
		}
		if prop.CreatedBy != sess.UserId { // allow owner to edit
			if !sess.IsSuperAdmin { // but allow superadmin to edit
				out.SetError(400, ErrRealtorUpsertPropertyNotOwnedByYou)
				return
			}
		}
	}

	prop.SetAll(in.Property, M.SB{
		mProperty.PriceHistoriesSell: true,
		mProperty.PriceHistoriesRent: true,
		mProperty.ApprovalState:      true,
	}, M.SB{}) // TODO: add list of fields to exclude

	prop.SetApprovalState(mProperty.ApprovalStatePending)
	prop.SetUpdatedAt(in.UnixNow())
	prop.SetUpdatedBy(sess.UserId)
	if prop.Id == 0 {
		prop.SetCreatedAt(in.UnixNow())
		prop.SetCreatedBy(sess.UserId)
	}
	prop.SetUniqPropKey(fmt.Sprintf(`%d_%d`, sess.UserId, fnv1a.HashString64(in.Property.FormattedAddress)))
	fmt.Println("Altitude = ", in.Property.Altitude)

	dup := rqProperty.NewProperty(d.PropOltp)
	dup.UniqPropKey = prop.UniqPropKey
	if dup.FindByUniqPropKey() {
		if dup.Id != prop.Id {
			out.SetError(400, ErrRealtorUpsertPropertyAddressAlreadyAdded)
			return
		}
	}

	if prop.DoUpsert() {
		if in.Property.Id == 0 {
			// Get user email, send message to their email
			user := rqAuth.NewUsers(d.AuthOltp)
			err := d.Mailer.SendNotifCreatePropertyEmail(user.Email,
				fmt.Sprintf("%s/realtor/ownedProperty/%v", conf.EnvWebConf().WebProtoDomain, in.Property.Id),
			)
			L.IsError(err, `SendNotifAddPropertyEmail`)
		}
	}

	if !prop.DoUpsert() {
		out.SetError(500, ErrRealtorUpsertPropertySaveFailed)
		return
	}

	prop.Adapter = nil
	out.Property = &prop.Property
	return
}
