package domain

import (
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

		Country     string `json:"country" form:"country" query:"country" long:"country" msg:"country"`
		CountryCode string `json:"country_code" form:"country_code" query:"country_code" long:"country_code" msg:"country_code"`
	}
)

const (
	UserGpsCountryAction = `user/GpsCountry`

	ErrUserGpsCountryUserNotFound       = `gps user not found`
	ErrUserGpsCountryFailedGetCountry   = `failed to get gps country`
	ErrUserGpsCountryFailedStoreCountry = `failed to store gps country`
)

func (d *Domain) UserGpsCountry(in *UserGpsCountryIn) (out UserGpsCountryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	country, iso2, err := d.Gmap.GetCountryByLatLng(in.CenterLat, in.CenterLong)
	if err != nil {
		out.SetError(500, ErrUserGpsCountryFailedGetCountry)
		return
	}
	out.Country = country
	out.CountryCode = iso2

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

	user.SetCountry(iso2)

	if !user.DoUpsert() {
		out.SetError(500, ErrUserGpsCountryFailedStoreCountry)
		return
	}

	return
}
