package domain

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"

	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserGpsCountry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserGpsCountry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserGpsCountry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserGpsCountry.go
//go:generate farify doublequote --file UserGpsCountry.go

type (
	UserGpsCountryIn struct {
		RequestCommon

		CenterLat  float64 `json:"centerLat" form:"centerLat" query:"centerLat" long:"centerLat" msg:"centerLat"`
		CenterLong float64 `json:"centerLong" form:"centerLong" query:"centerLong" long:"centerLong" msg:"centerLong"`

		CheckOnly bool `json:"checkOnly" form:"checkOnly" query:"checkOnly" long:"checkOnly" msg:"checkOnly"`
	}

	UserGpsCountryOut struct {
		ResponseCommon

		CountryName string `json:"countryName" form:"countryName" query:"countryName" long:"countryName" msg:"countryName"`
		CountryIso2 string `json:"countryIso2" form:"countryIso2" query:"countryIso2" long:"countryIso2" msg:"countryIso2"`
	}
)

const (
	UserGpsCountryAction = `user/GpsCountry`

	ErrUserGpsCountryUserNotFound       = `gps user not found`
	ErrUserGpsCountryFailedGetCountry   = `failed to get gps country`
	ErrUserGpsCountryFailedStoreCountry = `failed to store gps country`
)

var userGpsCountryCache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(1*time.Hour))

func (d *Domain) UserGpsCountry(in *UserGpsCountryIn) (out UserGpsCountryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	cacheKey := fmt.Sprintf(`%.3f %.3f`, in.CenterLat, in.CenterLong) // 111 meter difference
	if latLong, err := userGpsCountryCache.Get(cacheKey); err == nil {
		str := string(latLong)
		arr := strings.Split(str, `|`)
		if len(arr) == 2 {
			out.CountryName = arr[0]
			out.CountryIso2 = arr[1]
		}
	}
	if out.CountryName == `` || out.CountryIso2 == `` {
		country, iso2, err := d.Gmap.GetCountryByLatLng(in.CenterLat, in.CenterLong)
		if err != nil {
			out.SetError(500, ErrUserGpsCountryFailedGetCountry)
			return
		}
		out.CountryName = country
		out.CountryIso2 = iso2
		_ = userGpsCountryCache.Set(cacheKey, []byte(fmt.Sprintf(`%s|%s`, country, iso2)))
	}

	if in.CheckOnly {
		return
	}

	// set to user profile
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(403, ErrUserGpsCountryUserNotFound)
		return
	}

	user.SetCountry(out.CountryIso2)

	if !user.DoUpsert() {
		out.SetError(500, ErrUserGpsCountryFailedStoreCountry)
		return
	}

	return
}
