package domain

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
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

		LimitEach int `json:"limitEach" form:"limitEach" query:"limitEach" long:"limitEach" msg:"limitEach"`
	}

	UserNearbyFacilitiesOut struct {
		ResponseCommon

		Facilities []xGmap.Place `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	}
)

const (
	UserNearbyFacilitiesAction = `user/nearbyFacilities`

	ErrFailedRetrieveFacilities     = `failed to get nearby facilities`
	ErrFailedReadingFacilitiesCache = `failed to read nearby facilities from cache`

	WarnFailedWritingFacilitiesCache = `failed to write nearby facilities cache`
)

func (d *Domain) UserNearbyFacilities(in *UserNearbyFacilitiesIn) (out UserNearbyFacilitiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	var facilities []xGmap.Place

	const cacheDistance = 10 // 1 = 111.11 km | 10 = 11.11 km | 100 = 1.11 km
	cacheFilename := fmt.Sprintf(`%s/nearbyFacilities_%.0f_%.0f.json`, d.CacheDir, in.CenterLat*cacheDistance, in.CenterLong*cacheDistance)
	if L.FileExists(cacheFilename) {
		bytes := []byte(L.ReadFile(cacheFilename))
		err := json.Unmarshal(bytes, &facilities)
		if !L.IsError(err, ErrFailedReadingFacilitiesCache) {
			d.Gmap.RecalculateFacilitiesDistance(in.CenterLat, in.CenterLong, facilities)
		}

		sort.Slice(facilities, func(i, j int) bool {
			return facilities[i].DistanceKm < facilities[j].DistanceKm
		})

		out.AddTrace(`cache hit`)
	}

	if len(facilities) == 0 {
		eg := errgroup.Group{}

		m := sync.Mutex{}
		var errCount int

		add := func(rows []xGmap.Place, err error) error {
			m.Lock()
			defer m.Unlock()

			facilities = append(facilities, rows...)
			if err != nil {
				L.Print(err)
				errCount++
			}
			return nil
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
			res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeBusStation)
			return add(res, err)
		})

		eg.Go(func() error {
			res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeSubwayStation)
			return add(res, err)
		})

		eg.Go(func() error {
			res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeAirport)
			return add(res, err)
		})

		eg.Go(func() error {
			res, err := d.Gmap.NearbyFacilities(in.CenterLat, in.CenterLong, xGmap.TypeTrainStation)
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
		sort.Slice(facilities, func(i, j int) bool {
			return facilities[i].DistanceKm < facilities[j].DistanceKm
		})

		// write to cache
		jsonStr := X.ToJsonPretty(facilities)
		if !L.CreateFile(cacheFilename, jsonStr) {
			out.AddTrace(WarnFailedWritingFacilitiesCache)
		}
	}

	if in.LimitEach <= 0 {
		in.LimitEach = 5
	}

	facilitiesCountByType := make(map[string]int)
	for _, facility := range facilities {
		if facilitiesCountByType[facility.Type] < in.LimitEach {
			out.Facilities = append(out.Facilities, facility)
			facilitiesCountByType[facility.Type] += 1
		}
	}

	return

}
