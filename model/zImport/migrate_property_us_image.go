package zImport

import (
	"strconv"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
)

type PhotoUrls struct {
	FullScreenPhotoUrl              string `json:"fullScreenPhotoUrl"`
	LightboxListUrl                 string `json:"lightboxListUrl"`
	NonFullScreenPhotoUrl           string `json:"nonFullScreenPhotoUrl"`
	NonFullScreenPhotoUrlCompressed string `json:"nonFullScreenPhotoUrlCompressed"`
}

type Photo struct {
	FileName  string    `json:"fileName"`
	PhotoUrls PhotoUrls `json:"photoUrls"`
}

type Property struct {
	Photos []Photo `json:"photos"`
}

type PropertyUSMedia struct {
	PropertyMap map[string]Property `json:""`
}

func MigratePropertyUSImage(adapter *Tt.Adapter, minPropertyId int, maxPropertyId int) {

	rqPropertyUSExtra := rqProperty.NewPropertyExtraUS(adapter)
	rqPropertyUSExtra.FindByPropertyKey()

	propertyUSMutator := wcProperty.NewPropertyUSMutator(adapter)

	stat := &ImporterStat{Total: maxPropertyId - minPropertyId}
	defer stat.Print(`last`)

	for i := minPropertyId; i <= maxPropertyId; i++ {
		stat.Print()
		propKey := `rf` + strconv.Itoa(i)
		rqPropertyUSExtra.PropertyKey = propKey
		if !rqPropertyUSExtra.FindByPropertyKey() {
			stat.Skip()
			continue
		}

		propertyUSMutator.UniqPropKey = propKey
		if !propertyUSMutator.FindByUniqPropKey() {
			L.Print("There are no prop with unique prop key", propKey)
			stat.Skip()
			continue
		}

		jsonData := []byte(rqPropertyUSExtra.MediaSourceJson)

		var propertyUSMedia PropertyUSMedia
		err := json.Unmarshal(jsonData, &propertyUSMedia.PropertyMap)
		if err != nil {
			L.Print("Error unmarshaling JSON:", err)
			return
		}

		if len(propertyUSMedia.PropertyMap) == 0 {
			L.Print("There are no images for this property")
			stat.Skip()
			continue
		}

		var images []interface{}

		// Extract and print photo URLs
		for _, property := range propertyUSMedia.PropertyMap {
			for _, photo := range property.Photos {
				images = append(images, photo.PhotoUrls.FullScreenPhotoUrl)
			}
		}

		propertyUSMutator.Images = images
		stat.Ok(propertyUSMutator.DoOverwriteByUniqPropKey())
	}
}
