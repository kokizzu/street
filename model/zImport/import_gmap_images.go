package zImport

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/conf"
	"street/domain"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
)

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

	for { // for all image
		props := p.FindAllPropertiesOffsetLimit(start, 1000)
		if len(props) == 0 {
			break
		}
		stat.Total += 1000
		start += 1000

		for _, prop := range props {

			stat.Print()

			if len(prop.Images) > 0 { // if already have image
				stat.Skip()
				continue
			}
			if len(prop.Coord) < 2 { // no lat, long
				stat.Skip()
				continue
			}
			lat := X.ToF(prop.Coord[0])
			long := X.ToF(prop.Coord[1])

			// find image
			reader := StreetViewImageFromLatLong(400, 400, lat, long, 85, 70, 3, conf.PublicApiKey)
			if reader == nil {
				stat.Fail(``)
				continue
			}

			// cache to local
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

			// save to property
			prop.Adapter = d.PropOltp
			updatedProp := wcProperty.PropertyMutator{
				Property: *prop,
			}
			updatedProp.SetImages([]any{out.UrlPattern})

			stat.Ok(updatedProp.DoUpdateById())
		}
	}
}

func StreetViewImageFromLatLong(width, height int, lat, lng float64, fov, heading, pitch int, apiKey string) io.ReadCloser {
	url := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/streetview?size=%dx%d&location=%.15f,%.15f&fov=%d&heading=%d&pitch=%d&key=%s",
		width, height, lat, lng, fov, heading, pitch, apiKey,
	)
	resp, err := http.Get(url)
	if L.IsError(err, `StreetViewImageFromLatLong`) {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		L.Print(`StreetViewImageFromLatLong.StatusCode.not200`, resp.StatusCode)
		return nil
	}
	return resp.Body
}
