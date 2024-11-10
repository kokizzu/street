package domain

import (
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserLikeProp.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserLikeProp.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserLikeProp.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserLikeProp.go
//go:generate farify doublequote --file UserLikeProp.go

type (
	UserLikePropIn struct {
		RequestCommon

		PropId uint64 `json:"propId,string" form:"propId" query:"propId" long:"propId" msg:"propId"`
		Like   bool   `json:"like" form:"like" query:"like" long:"like" msg:"like"`
	}

	UserLikePropOut struct {
		ResponseCommon
	}
)

const (
	UserLikePropAction = `user/likeProp`

	ErrUserLikePropNotFound     = `property to like not found`
	ErrUserLikePropFailedUnlike = `failed to unlike property`
	ErrUserLikePropFailedLike   = `failed to like property`
)

func (d *Domain) UserLikeProp(in *UserLikePropIn) (out UserLikePropOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.PropId

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	prop := rqProperty.NewProperty(d.PropOltp)
	prop.Id = in.PropId
	if !prop.FindById() {
		out.SetError(404, ErrUserLikePropNotFound)
		return
	}

	like := wcProperty.NewUserPropLikesMutator(d.PropOltp)
	like.UserId = sess.UserId
	like.PropId = in.PropId
	like.CreatedAt = in.UnixNow()

	delta := 0

	if like.FindByUserIdPropId() { // liked
		if !in.Like { // to be unliked
			if !like.DoDeletePermanentByUserIdPropId() { // delete like
				out.SetError(500, ErrUserLikePropFailedUnlike)
				return
			}
			delta = -1
		}
	} else { // never liked, or like deleted
		if in.Like { // to be liked
			if !like.DoInsert() { // insert like
				out.SetError(500, ErrUserLikePropFailedLike)
				return
			}

			delta = +1
		}
	}

	if delta != 0 {
		// increase/decrease counter
		count := wcProperty.NewPropLikeCountMutator(d.PropOltp)
		count.PropId = in.PropId
		count.DoInc(delta) // ignore error we can synchronize by query preriodically
	}

	return
}
