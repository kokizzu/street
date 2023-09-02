package zImport

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"golang.org/x/exp/rand"

	"street/conf"
	"street/domain"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
)

//var streetViewCache = map[string]string{}

func ImportStreetViewImage(d *domain.Domain, conf conf.GmapConf) {
	p := rqProperty.NewProperty(d.PropOltp)
	start := 0

	const UA = `ServiceAccount/1.0`
	session, _ := d.CreateSession(1, `admin@localhost`, UA)
	if !session.DoInsert() {
		L.Print(`failed inserting session`)
		return
	}

	rc := domain.NewLocalRequestCommon(session.SessionToken, UA)

	stat := &ImporterStat{}
	defer stat.Print(`last`)

	// all possible degree to 360, increment by 15
	degrees := []int{0, 15, 30, 45, 60, 75, 90, 105, 120, 135, 150, 160, 165, 180, 195, 210, 225, 240, 255, 270, 285, 300, 315, 330, 345}

	for { // for all image
		props := p.FindAllPropertiesOffsetLimit(start, 1000)
		if len(props) == 0 {
			break
		}
		stat.Total += 1000
		start += 1000

		for _, prop := range props {

			stat.Print()
			if len(prop.Coord) < 2 { // no lat, long
				stat.Skip()
				continue
			}
			lat := X.ToF(prop.Coord[0])
			long := X.ToF(prop.Coord[1])

			//cacheKey := fmt.Sprintf("%.15f,%.15f", lat, long)

			if len(prop.Images) > 0 { // if already have image
				stat.Skip()
				//streetViewCache[cacheKey] = X.ToS(prop.Images[0]) // save to cache
				continue
			}

			//if _, ok := streetViewCache[cacheKey]; !ok { // if not on cache

			// find image, from random id angle
			rand.Shuffle(len(degrees), func(i, j int) { degrees[i], degrees[j] = degrees[j], degrees[i] })

			var reader io.ReadCloser
			for _, degree := range degrees {
				reader = StreetViewImageFromLatLong(400, 400, lat, long, 90, degree, 3, conf.PublicApiKey)
				if reader == nil {
					continue
				}
				break
			}
			if reader == nil {
				stat.Fail(`StreetViewImageFromLatLong.noPossibleDegree`)
				continue
			}

			// save file to local
			rc.Action = domain.UserUploadFileAction
			rc.RawFile = domain.NewLocalRawFileFromReader(`streetview.jpg`, reader)
			out := d.UserUploadFile(&domain.UserUploadFileIn{
				RequestCommon: rc,
				Purpose:       domain.UserUploadFile_PurposeProperty,
			})

			if out.HasError() {
				stat.Fail(`UserUploadFile`)
				continue
			}

			//streetViewCache[cacheKey] = out.UrlPattern // cache after save
			//}

			// save to property
			prop.Adapter = d.PropOltp
			updatedProp := wcProperty.PropertyMutator{
				Property: *prop,
			}
			//updatedProp.SetImages([]any{streetViewCache[cacheKey]})
			updatedProp.SetImages([]any{out.UrlPattern})

			stat.Ok(updatedProp.DoUpdateById())
		}
	}
}

func StreetViewImageFromLatLong(width, height int, lat, lng float64, fov, heading, pitch int, apiKey string) io.ReadCloser {
	// https://developers.google.com/maps/documentation/streetview/request-streetview
	// heading indicates the compass heading of the camera. Accepted values are from 0 to 360 (both values indicating North, with 90 indicating East, and 180 South). If no heading is specified, a value will be calculated that directs the camera towards the specified location, from the point at which the closest photograph was taken.
	// fov (default is 90) determines the horizontal field of view of the image. The field of view is expressed in degrees, with a maximum allowed value of 120. When dealing with a fixed-size viewport, as with a Street View image of a set size, field of view in essence represents zoom, with smaller numbers indicating a higher level of zoom.
	// pitch (default is 0) specifies the up or down angle of the camera relative to the Street View vehicle. This is often, but not always, flat horizontal. Positive values angle the camera up (with 90 degrees indicating straight up); negative values angle the camera down (with -90 indicating straight down).
	// radius (default is 50) sets a radius, specified in meters, in which to search for a panorama, centered on the given latitude and longitude. Valid values are non-negative integers.
	url := fmt.Sprintf(
		`https://maps.googleapis.com/maps/api/streetview?size=%dx%d&location=%.15f,%.15f&fov=%d&heading=%d&pitch=%d&key=%s&return_error_code=true`,
		width, height, lat, lng, fov, heading, pitch, apiKey,
	)
	resp, err := http.Get(url)
	if L.IsError(err, `StreetViewImageFromLatLong`) {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	return resp.Body
}
