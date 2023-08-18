package domain

import (
	"github.com/kokizzu/gotro/X"

	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserSearchProp.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserSearchProp.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserSearchProp.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserSearchProp.go
//go:generate farify doublequote --file UserSearchProp.go

type (
	UserSearchPropIn struct {
		RequestCommon `json:"request_common"`

		CenterLat  float64 `json:"centerLat" form:"centerLat" query:"centerLat" long:"centerLat" msg:"centerLat"`
		CenterLong float64 `json:"centerLong" form:"centerLong" query:"centerLong" long:"centerLong" msg:"centerLong"`
		Offset     int     `json:"offset" form:"offset" query:"offset" long:"offset" msg:"offset"`
		Limit      int     `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`

		MaxDistanceKM float64 `json:"maxDistanceKM" form:"maxDistanceKM" query:"maxDistanceKM" long:"maxDistanceKM" msg:"maxDistanceKM"`
	}

	UserSearchPropOut struct {
		ResponseCommon

		Properties []Property `json:"properties" form:"properties" query:"properties" long:"properties" msg:"properties"`
	}

	Property struct {
		*rqProperty.Property
		Lat float64 `json:"lat" form:"lat" query:"lat" long:"lat" msg:"lat"`
		Lng float64 `json:"lng" form:"lng" query:"lng" long:"lng" msg:"lng"`

		DistanceKM float64 `json:"distanceKM" form:"distanceKM" query:"distanceKM" long:"distanceKM" msg:"distanceKM"`

		id uint64
	}
)

const (
	UserSearchPropAction = `user/searchProp`

	ErrSearchPropFailed = `search prop failed`

	DefaultLimit = 10
)

func (d *Domain) UserSearchProp(in *UserSearchPropIn) (out UserSearchPropOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	prop := rqProperty.NewProperty(d.PropOltp)

	if in.Limit == 0 {
		in.Limit = DefaultLimit
	}

	if in.MaxDistanceKM <= 0 {
		in.MaxDistanceKM = 2
	}

	out.Properties = make([]Property, 0, in.Limit)
	ok := prop.FindByLatLong(in.CenterLat, in.CenterLong, in.Limit, in.Offset, func(row []any) bool {
		item := Property{Property: &rqProperty.Property{}}
		item.FromArray(row)
		if len(item.Coord) >= 2 {
			item.Lat = X.ToF(item.Coord[0])
			item.Lng = X.ToF(item.Coord[1])
		}
		item.id = item.Id
		item.DistanceKM = mProperty.DistanceKm(item.Lat, item.Lng, in.CenterLat, in.CenterLong)
		if item.DistanceKM > in.MaxDistanceKM {
			return false
		}
		out.Properties = append(out.Properties, item)
		return true
	})
	if !ok {
		out.SetError(500, ErrSearchPropFailed)
		return
	}
	return
}
