package domain

import (
	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserGpsLocation.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserGpsLocation.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserGpsLocation.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserGpsLocation.go
//go:generate farify doublequote --file UserGpsLocation.go

type (
	UserGpsLocationIn struct {
		RequestCommon

		CenterLat  float64 `json:"centerLat" form:"centerLat" query:"centerLat" long:"centerLat" msg:"centerLat"`
		CenterLong float64 `json:"centerLong" form:"centerLong" query:"centerLong" long:"centerLong" msg:"centerLong"`
	}

	UserGpsLocationOut struct {
		ResponseCommon

		Country string `json:"country" form:"country" query:"country" long:"country" msg:"country"`
	}
)

const (
	UserGpsLocationAction = `user/gpsLocation`

	ErrUserGpsLocationUserNotFound     = `gps user not found`
	ErrUserGpsLocationFailedSetCountry = `failed to set gps country`
)

func (d *Domain) UserGpsLocation(in *UserGpsLocationIn) (out UserGpsLocationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	var countryCode2digit = `TW`
	// TODO: call to googlemap api

	// set to user profile
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(403, ErrUserGpsLocationUserNotFound)
		return
	}
	user.SetCountry(countryCode2digit)
	if !user.DoUpsert() {
		out.SetError(500, ErrUserGpsLocationFailedSetCountry)
		return
	}

	return
}
