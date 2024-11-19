package domain

import (
	"strconv"

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
		*rqProperty.PropertyWithNote

		Lat float64 `json:"lat" form:"lat" query:"lat" long:"lat" msg:"lat"`
		Lng float64 `json:"lng" form:"lng" query:"lng" long:"lng" msg:"lng"`

		DistanceKM float64 `json:"distanceKM" form:"distanceKM" query:"distanceKM" long:"distanceKM" msg:"distanceKM"`

		Liked     bool  `json:"liked" form:"liked" query:"liked" long:"liked" msg:"liked"`
		LikeCount int64 `json:"likeCount" form:"likeCount" query:"likeCount" long:"likeCount" msg:"likeCount"`

		id uint64
	}
)

const (
	UserSearchPropAction = `user/searchProp`

	ErrSearchPropFailed = `search prop failed`

	DefaultLimit = 10
)

func mergePropertyWithSerialNumber(inputProperties []Property) ([]Property, []uint64) {
	filterProperties := []Property{}
	filterPropIds := []uint64{}

	// Merge property history with same serial number
	totalSizePropBySerialNumber := make(map[string]*Property)

	for _, prop := range inputProperties {
		prop := prop
		propSize, err := strconv.ParseFloat(prop.SizeM2, 64)

		// Check for errors
		if err != nil {
			continue
		}

		if totalSizePropBySerialNumber[prop.SerialNumber] == nil {
			totalSizePropBySerialNumber[prop.SerialNumber] = &prop
		} else {
			// Existing size
			existingPropSize, err := strconv.ParseFloat(totalSizePropBySerialNumber[prop.SerialNumber].SizeM2, 64)

			if err != nil {
				continue
			}

			// Accumulate size of property
			appendSize := existingPropSize + propSize
			totalSizePropBySerialNumber[prop.SerialNumber].SizeM2 = strconv.FormatFloat(appendSize, 'f', 2, 64)
		}
	}

	for _, prop := range totalSizePropBySerialNumber {
		prop := prop
		filterProperties = append(filterProperties, *prop)
		filterPropIds = append(filterPropIds, prop.Id)
	}

	return filterProperties, filterPropIds
}

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
	var propIds = make([]uint64, 0, in.Limit)

	// Get satisfied property with expected condition
	satisfiedProperties := make([]Property, 0, in.Limit)

	ok := prop.FindByLatLongAndCountry(d.PropOltp, sess.Country, in.CenterLat, in.CenterLong, in.Limit, in.Offset, func(row []any) bool {
		item := Property{PropertyWithNote: &rqProperty.PropertyWithNote{Property: &rqProperty.Property{}}}
		item.FromArray(row)
		if item.DeletedAt > 0 { // skip deleted property
			return true
		}
		if len(item.Coord) >= 2 {
			item.Lat = X.ToF(item.Coord[0])
			item.Lng = X.ToF(item.Coord[1])
		}
		item.id = item.Id
		item.DistanceKM = conf.DistanceKm(item.Lat, item.Lng, in.CenterLat, in.CenterLong)
		if item.DistanceKM > in.MaxDistanceKM {
			return false
		}

		satisfiedProperties = append(satisfiedProperties, item)
		return true
	})

	// Merge property history with same serial number
	filterProperties, filterPropIds := mergePropertyWithSerialNumber(satisfiedProperties)

	out.Properties = filterProperties
	propIds = filterPropIds

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
			CountryCode: v.CountryCode,
			PropertyId: v.Id,
		})
	}
	
	return
}