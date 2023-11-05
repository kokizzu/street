package domain

import (
	"fmt"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/X"
	"github.com/segmentio/fasthash/fnv1a"

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

		AskReview bool `json:"askReview" form:"askReview" query:"askReview" long:"askReview" msg:"askReview"`

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

	prop.SetUpdatedAt(in.UnixNow())
	prop.SetUpdatedBy(sess.UserId)
	if prop.Id == 0 {
		prop.SetCreatedAt(in.UnixNow())
		prop.SetCreatedBy(sess.UserId)
	}
	prop.SetUniqPropKey(fmt.Sprintf(`%d_%d`, sess.UserId, fnv1a.HashString64(in.Property.FormattedAddress)))

	dup := rqProperty.NewProperty(d.PropOltp)
	dup.UniqPropKey = prop.UniqPropKey
	if dup.FindByUniqPropKey() {
		if dup.Id != prop.Id {
			out.SetError(400, ErrRealtorUpsertPropertyAddressAlreadyAdded)
			return
		}
	}

	prop.NormalizeFloorList()

	var prevReview string
	if in.AskReview || in.Property.Id == 0 {
		if prop.ApprovalState == `` {
			prevReview = `previously already approved just updated by realtor`
		} else {
			prevReview = `previous review: ` + prop.ApprovalState
		}
		in.AskReview = prop.SetApprovalState(mProperty.ApprovalStatePending)
	} else if prop.ApprovalState == `` {
		prop.SetApprovalState(`pending, property updated by realtor, need to ask review again`)
	}

	if prop.DoUpsert() {
		if in.Property.Id == 0 {
			// Get user email, send message to their email
			d.runSubtask(func() {
				err := d.Mailer.SendNotifCreatePropertyEmail(sess.Email,
					fmt.Sprintf("%s/realtor/property/%v", d.WebCfg.WebProtoDomain, in.Property.Id),
				)
				L.IsError(err, `SendNotifAddPropertyEmail`)
			})
		}
	} else {
		out.SetError(500, ErrRealtorUpsertPropertySaveFailed)
		return
	}

	if in.AskReview {
		d.runSubtask(func() {
			err := d.Mailer.SendPendingPropertyReviewToAdmin(
				fmt.Sprintf("%s/realtor/property/%v", d.WebCfg.WebProtoDomain, prop.Id),
				X.ToS(prop.Id), prevReview,
			)
			L.IsError(err, `SendPendingPropertyReviewToAdmin`)
		})
	}

	prop.Adapter = nil
	out.Property = &prop.Property
	return
}
