package domain

import (
	"sort"
	"sync"

	"golang.org/x/sync/errgroup"

	"street/model/xGmap"
)

// https://developers.google.com/maps/documentation/places/web-service/search-nearby
//

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserNearbyFacilities.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserNearbyFacilities.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserNearbyFacilities.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserNearbyFacilities.go
//go:generate farify doublequote --file UserNearbyFacilities.go

type (
	UserNearbyFacilitiesIn struct {
		RequestCommon `json:"request_common"`

		CenterLat  float64 `json:"centerLat" form:"centerLat" query:"centerLat" long:"centerLat" msg:"centerLat"`
		CenterLong float64 `json:"centerLong" form:"centerLong" query:"centerLong" long:"centerLong" msg:"centerLong"`
	}

	UserNearbyFacilitiesOut struct {
		ResponseCommon

		Facilities []xGmap.Place `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	}
)

const (
	UserNearbyFacilitiesAction = `user/nearbyFacilitites`

	ErrFailedRetrieveFacilities = `failed to get nearby facilities`
)

func (d *Domain) UserNearbyFacilities(in *UserNearbyFacilitiesIn) (out UserNearbyFacilitiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	eg := errgroup.Group{}

	m := sync.Mutex{}
	var errCount int

	add := func(rows []xGmap.Place, err error) error {
		m.Lock()
		defer m.Unlock()
		out.Facilities = append(out.Facilities, rows...)
		errCount++
		return err
	}

	const totalCall = 5

	eg.Go(func() error {
		res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeHospital)
		return add(res, err)
	})

	eg.Go(func() error {
		res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeSchool)
		return add(res, err)
	})

	eg.Go(func() error {
		res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeRestaurant)
		return add(res, err)
	})

	eg.Go(func() error {
		res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeSubwayStation)
		return add(res, err)
	})

	eg.Go(func() error {
		res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeConvenienceStore)
		return add(res, err)
	})

	_ = eg.Wait()

	if errCount == totalCall {
		out.SetError(500, ErrFailedRetrieveFacilities)
		return
	}

	// sort by distance
	sort.Slice(out.Facilities, func(i, j int) bool {
		return out.Facilities[i].DistanceKm < out.Facilities[j].DistanceKm
	})

	return

}
