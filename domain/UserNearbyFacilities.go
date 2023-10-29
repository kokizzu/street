package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/kokizzu/gotro/L"
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
	ErrFailedWritingFacilitiesCache = `failed to write nearby facilities into cache`
	ErrFailedReadingFacilitiesCache = `failed to read nearby facilities from cache`
)

func (d *Domain) UserNearbyFacilities(in *UserNearbyFacilitiesIn) (out UserNearbyFacilitiesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	var facilities []xGmap.Place

	cacheFilename := fmt.Sprintf(`%s/nearbyFacilities_%.0f_%.0f.json`, d.CacheDir, in.CenterLat*100, in.CenterLong*100)
	if L.FileExists(cacheFilename) {
		bytes := []byte(L.ReadFile(cacheFilename))
		err := json.Unmarshal(bytes, &facilities)
		if !L.IsError(err, ErrFailedReadingFacilitiesCache) {
			d.Gmap.RecalculateFacilitiesDistance(in.CenterLat, in.CenterLong, facilities)
		}

		sort.Slice(facilities, func(i, j int) bool {
			return facilities[i].DistanceKm < facilities[j].DistanceKm
		})
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
		sort.Slice(facilities, func(i, j int) bool {
			return facilities[i].DistanceKm < facilities[j].DistanceKm
		})

		// write to cache
		bytes, err := json.Marshal(facilities)
		if !L.IsError(err, `json.Marshal: %#v`, facilities) {
			err := os.WriteFile(cacheFilename, bytes, 0644)
			L.IsError(err, ErrFailedWritingFacilitiesCache)
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
