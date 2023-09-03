package zImport

import (
	"io"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"golang.org/x/exp/rand"

	"street/domain"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/xGmap"
)

//var streetViewCache = map[string]string{}

func ImportStreetViewImage(d *domain.Domain, gmap xGmap.Gmap) {
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
				reader = gmap.StreetViewImageFromLatLong(400, 400, lat, long, 90, degree, 3)
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
