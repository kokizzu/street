package xGmap

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/L"

	"street/conf"
)

type Gmap struct {
	conf.GmapConf
}

func (g Gmap) BuildFullLocationSearchUrl(address string) string {
	const googleApiUrl = `https://maps.googleapis.com/maps/api/place/findplacefromtext/json`

	listFields := "formatted_address,name,rating,opening_hours,geometry"
	inputType := "textquery"
	req, err := http.NewRequest("GET", googleApiUrl, nil)
	L.PanicIf(err, `BuildFullLocationSearchUrl`) // note: do not use this function for API
	q := req.URL.Query()
	q.Add("fields", listFields)
	q.Add("input", address)
	q.Add("inputtype", inputType)
	q.Add("key", g.GmapConf.PrivateApiKey)
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func (g Gmap) StreetViewImageFromLatLong(width, height int, lat, lng float64, fov, heading, pitch int) io.ReadCloser {
	// https://developers.google.com/maps/documentation/streetview/request-streetview
	// heading indicates the compass heading of the camera. Accepted values are from 0 to 360 (both values indicating North, with 90 indicating East, and 180 South). If no heading is specified, a value will be calculated that directs the camera towards the specified location, from the point at which the closest photograph was taken.
	// fov (default is 90) determines the horizontal field of view of the image. The field of view is expressed in degrees, with a maximum allowed value of 120. When dealing with a fixed-size viewport, as with a Street View image of a set size, field of view in essence represents zoom, with smaller numbers indicating a higher level of zoom.
	// pitch (default is 0) specifies the up or down angle of the camera relative to the Street View vehicle. This is often, but not always, flat horizontal. Positive values angle the camera up (with 90 degrees indicating straight up); negative values angle the camera down (with -90 indicating straight down).
	// radius (default is 50) sets a radius, specified in meters, in which to search for a panorama, centered on the given latitude and longitude. Valid values are non-negative integers.
	url := fmt.Sprintf(
		`https://maps.googleapis.com/maps/api/streetview?size=%dx%d&location=%.15f,%.15f&fov=%d&heading=%d&pitch=%d&key=%s&return_error_code=true`,
		width, height, lat, lng, fov, heading, pitch, g.PublicApiKey,
	)
	resp, err := http.Get(url)
	if L.IsError(err, `StreetViewImageFromLatLong`) {
		return nil
	}
	if resp.StatusCode != http.StatusOK || resp == nil || resp.Body == nil {
		return nil
	}
	return resp.Body
}

const (
	// allowed types: https://developers.google.com/maps/documentation/places/web-service/supported_types
	TypeRestaurant       = `restaurant`
	TypeSchool           = `school`
	TypeHospital         = `hospital`
	TypeSubwayStation    = `subway_station`
	TypeConvenienceStore = `convenience_store`
)

type Place struct {
	Name       string  `json:"name,omitempty"`
	Address    string  `json:"address,omitempty"`
	Lat        float64 `json:"lat,omitempty"`
	Lng        float64 `json:"lng,omitempty"`
	Type       string  `json:"type,omitempty"`
	DistanceKm float64 `json:"distanceKm,omitempty"`
}

type GmapSearchNearbyResult struct {
	//HTMLAttributions []any `json:"html_attributions"`
	Results []struct {
		//BusinessStatus string `json:"business_status"` // if keyword= not set, business_status of CLOSED_TEMPORARILY or CLOSED_PERMANENTLY will not be returned.
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			//Viewport struct {
			//	Northeast struct {
			//		Lat float64 `json:"lat"`
			//		Lng float64 `json:"lng"`
			//	} `json:"northeast"`
			//	Southwest struct {
			//		Lat float64 `json:"lat"`
			//		Lng float64 `json:"lng"`
			//	} `json:"southwest"`
			//} `json:"viewport"`
		} `json:"geometry"`
		//Icon                string `json:"icon"`
		//IconBackgroundColor string `json:"icon_background_color"`
		//IconMaskBaseURI     string `json:"icon_mask_base_uri"`
		Name string `json:"name"`
		//OpeningHours        struct {
		//	OpenNow bool `json:"open_now"`
		//} `json:"opening_hours"`
		//Photos []struct {
		//	Height           int64    `json:"height"`
		//	HTMLAttributions []string `json:"html_attributions"`
		//	PhotoReference   string   `json:"photo_reference"`
		//	Width            int64    `json:"width"`
		//} `json:"photos"`
		//PlaceID  string `json:"place_id"`
		//PlusCode struct {
		//	CompoundCode string `json:"compound_code"`
		//	GlobalCode   string `json:"global_code"`
		//} `json:"plus_code"`
		//PriceLevel       int64    `json:"price_level"`
		//Rating           int64    `json:"rating"`
		//Reference        string   `json:"reference"`
		//Scope            string   `json:"scope"`
		//Types            []string `json:"types"`
		//UserRatingsTotal int64    `json:"user_ratings_total"`
		Vicinity string `json:"vicinity"`
	} `json:"results"`
	Status string `json:"status"`
}

func (g Gmap) NearbyFacilities(lat float64, long float64, typ string) (res []Place, err error) {
	// https://developers.google.com/maps/documentation/places/web-service/search

	for z := 0; z < 2; z++ {
		var url string
		if z == 0 {
			url = fmt.Sprintf( // may be failed if google havent set the type: https://stackoverflow.com/questions/16123878/my-places-api-returns-zero-results
				`https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%.15f,%.15f&rankby=distance&type=%s&key=%s`,
				lat, long, typ, g.PrivateApiKey,
			)
		} else {
			url = fmt.Sprintf( // use keyword search instead
				`https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%.15f,%.15f&rankby=distance&keyword=%s&key=%s`,
				lat, long, typ, g.PrivateApiKey,
			)
		}

		resp, err := http.Get(url)
		if L.IsError(err, `Gmap) NearbyFacilities.http.Get`) {
			return nil, err
		}
		// intentionally ignore http status
		if resp == nil || resp.Body == nil {
			return nil, errors.New(`Gmap) NearbyFacilities.http.EmptyBody`)
		}

		// read all body
		body, err := io.ReadAll(resp.Body)
		if L.IsError(err, `Gmap) NearbyFacilities.io.ReadAll`) {
			return nil, err
		}

		var result GmapSearchNearbyResult
		err = json.Unmarshal(body, &result)
		if L.IsError(err, `Gmap) NearbyFacilities.json.Unmarshal`) {
			return nil, err
		}

		for _, row := range result.Results {
			res = append(res, Place{
				Name:       row.Name,
				Address:    row.Vicinity,
				Lat:        row.Geometry.Location.Lat,
				Lng:        row.Geometry.Location.Lng,
				Type:       typ,
				DistanceKm: conf.DistanceKm(lat, long, row.Geometry.Location.Lat, row.Geometry.Location.Lng),
			})
		}

		if result.Status == `ZERO_RESULT` {
			continue
		}
	}

	return res, nil
}

func (g Gmap) GetCountryByLatLng(lat float64, long float64) (country, iso2 string, err error) {
	url := fmt.Sprintf(
		`https://maps.googleapis.com/maps/api/geocode/json?latlng=%.15f,%.15f&sensor=false&key=%s`,
		lat, long, g.PrivateApiKey)

	resp, err := http.Get(url)
	if L.IsError(err, `Gmap) NearbyFacilities.http.Get`) {
		return ``, ``, err
	}
	// intentionally ignore http status
	if resp == nil || resp.Body == nil {
		return ``, ``, errors.New(`Gmap) NearbyFacilities.http.EmptyBody`)
	}

	// read all body
	body, err := io.ReadAll(resp.Body)
	if L.IsError(err, `Gmap) NearbyFacilities.io.ReadAll`) {
		return ``, ``, err
	}

	var data map[string]interface{}

	// Unmarshal the JSON into the map
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return ``, ``, errors.New(`Error unmarshaling JSON`)
	}

	var (
		countryName string
		countryIso2 string
	)
	addressComponents := data["results"].([]interface{})[0].(map[string]interface{})["address_components"].([]interface{})
	for _, component := range addressComponents {
		componentMap := component.(map[string]interface{})
		if types, exists := componentMap["types"].([]any); exists {
			if types[0] == "country" {
				if longName, exists := componentMap["long_name"].(string); exists {
					if shortName, exists := componentMap["short_name"].(string); exists {
						countryName = longName
						countryIso2 = shortName
						break
					}
				}
			}
		}
	}

	return countryName, countryIso2, nil
}
