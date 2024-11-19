package domain

import (
	"github.com/kokizzu/gotro/X"

	"street/conf"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/saProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserSearchProp.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserSearchProp.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserSearchProp.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserSearchProp.go
//go:generate farify doublequote --file UserSearchProp.go

type (
	UserSearchPropUSIn struct {
		RequestCommon `json:"request_common"`

		CenterLat  float64 `json:"centerLat" form:"centerLat" query:"centerLat" long:"centerLat" msg:"centerLat"`
		CenterLong float64 `json:"centerLong" form:"centerLong" query:"centerLong" long:"centerLong" msg:"centerLong"`
		Offset     int     `json:"offset" form:"offset" query:"offset" long:"offset" msg:"offset"`
		Limit      int     `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`

		MaxDistanceKM float64 `json:"maxDistanceKM" form:"maxDistanceKM" query:"maxDistanceKM" long:"maxDistanceKM" msg:"maxDistanceKM"`
	}

	UserSearchPropUSOut struct {
		ResponseCommon

		Properties []PropertyUS `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`
	}

	PropertyUS struct {
		*rqProperty.PropertyUS
		Lat float64 `json:"lat" form:"lat" query:"lat" long:"lat" msg:"lat"`
		Lng float64 `json:"lng" form:"lng" query:"lng" long:"lng" msg:"lng"`

		DistanceKM float64 `json:"distanceKM" form:"distanceKM" query:"distanceKM" long:"distanceKM" msg:"distanceKM"`

		Liked     bool  `json:"liked" form:"liked" query:"liked" long:"liked" msg:"liked"`
		LikeCount int64 `json:"likeCount" form:"likeCount" query:"likeCount" long:"likeCount" msg:"likeCount"`

		id uint64
	}
)

const (
	UserSearchPropUSAction = `user/searchProp/US`

	ErrSearchPropUSFailed = `search prop failed`

	DefaultLimitPropUS = 10
)

func (d *Domain) UserSearchPropUS(in *UserSearchPropUSIn) (out UserSearchPropUSOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	prop := rqProperty.NewPropertyUS(d.PropOltp)

	if in.Limit == 0 {
		in.Limit = DefaultLimit
	}

	if in.MaxDistanceKM <= 0 {
		in.MaxDistanceKM = 2
	}

	out.Properties = make([]PropertyUS, 0, in.Limit)
	propIds := make([]uint64, 0, in.Limit)
	ok := prop.FindPropUSByLatLong(in.CenterLat, in.CenterLong, in.Limit, in.Offset, func(row []any) bool {
		item := PropertyUS{PropertyUS: &rqProperty.PropertyUS{}}
		item.FromArray(row)
		if item.DeletedAt > 0 { // skip deleted property
			return true
		}
		if len(item.Coord) >= 2 {
			item.Lat = X.ToF(item.Coord[0])
			item.Lng = X.ToF(item.Coord[1])
		}
		item.id = item.Id
		item.NormalizeFloorList()
		item.DistanceKM = conf.DistanceKm(item.Lat, item.Lng, in.CenterLat, in.CenterLong)
		if item.DistanceKM > in.MaxDistanceKM {
			return false
		}
		out.Properties = append(out.Properties, item)
		propIds = append(propIds, item.Id)
		return true
	})
	if !ok {
		out.SetError(500, ErrSearchPropFailed)
		return
	}

	// get liked by this user and like count
	if len(propIds) > 0 {
		upl := rqProperty.NewUserPropLikes(d.PropOltp)
		upl.UserId = sess.UserId
		likedMap := upl.LikedMap(propIds)

		plc := rqProperty.NewPropLikeCount(d.PropOltp)
		countMap := plc.CountMap(propIds)

		for i := range out.Properties {
			out.Properties[i].Liked = likedMap[out.Properties[i].id]
			out.Properties[i].LikeCount = countMap[out.Properties[i].id]
		}
	}

	var (
		city string = ``
		state string = ``
	)

	if len(out.Properties) > 0 {
		city = out.Properties[0].City
		state = out.Properties[0].State
	}
	
	d.insertScannedAreas(saProperty.ScannedAreas{
		ActorId: sess.UserId,
		CreatedAt: in.TimeNow(),
		Latitude: in.Lat,
		Longitude: in.Long,
		City: city,
		State: state,
	})

	for _, v := range out.Properties {
		d.insertScannedProps(saProperty.ScannedProperties{
			ActorId: sess.UserId,
			CreatedAt: in.TimeNow(),
			CountryCode: `US`,
			PropertyId: v.Id,
		})
	}

	return
}
