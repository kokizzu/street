package zImport

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"street/model"
	"street/model/mProperty/wcProperty"
	"street/model/xGmap"
)

// PlaceResponse A Place Response
type PlaceResponse struct {
	Candidates []Candidate `json:"candidates"`
	Status     string      `json:"status"`
}

// Location lat and long
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// ViewPort viewport
type ViewPort struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// Geometry geometry model
type Geometry struct {
	Location Location `json:"location"`
	ViewPort ViewPort `json:"viewport"`
}

// Candidate candidate
type Candidate struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

func retrieveLatLongFromAddress(adapter *Tt.Adapter, gmap xGmap.Gmap) {
	defer subTaskPrint(`retrieveLatLongFromAddress: retrieve lat/long`)()

	propertyMutator := wcProperty.NewPropertyMutator(adapter)
	properties := propertyMutator.FindAllProperties()

	stat := &model.ImporterStat{Total: len(properties), PrintEvery: 10}
	defer stat.Print(`last`)

	for _, p := range properties {
		stat.Print()

		if p.Address == "" {
			stat.Skip()
			continue
		}
		if p.FormattedAddress != "" {
			stat.Skip()
			//log.Println("Property has address and lat/long already")
			continue
		}

		fullUrl := gmap.BuildFullLocationSearchUrl(p.Address)
		locationResponse, err := http.Get(fullUrl)

		if L.IsError(err, `retrieveLatLongFromAddress: get location response`) {
			stat.Fail(`fail request`)
			continue
		}

		responseData, err := io.ReadAll(locationResponse.Body)
		if L.IsError(err, `retrieveLatLongFromAddress: read response body`) {
			stat.Fail(`no resp body`)
			continue
		}
		//fmt.Println("Response data => " + string(responseData))
		propertyLocation := PlaceResponse{}

		err = json.Unmarshal(responseData, &propertyLocation)
		if L.IsError(err, `retrieveLatLongFromAddress: unmarshal response data`) {
			stat.Fail(`fail json decoded`)
			L.Print(p.Id, p.UniqPropKey, p.Address)
			log.Println(string(responseData))
			continue
		}
		if len(propertyLocation.Candidates) == 0 {
			stat.Skip()
			stat.Warn(`empty location`)
			//fmt.Println("There is no available location for this address")
			continue
		}

		dataMutator := wcProperty.NewPropertyMutator(adapter)
		dataMutator.Property = *p
		dataMutator.Adapter = adapter
		dataMutator.FormattedAddress = propertyLocation.Candidates[0].FormattedAddress

		latitude := propertyLocation.Candidates[0].Geometry.Location.Lat
		longitude := propertyLocation.Candidates[0].Geometry.Location.Lng

		dataMutator.Coord = []any{latitude, longitude}
		dataMutator.UpdatedAt = time.Now().Unix()

		stat.Ok(dataMutator.DoOverwriteById())
	}
}

func ImportHouseLocation(adapter *Tt.Adapter, gmap xGmap.Gmap) {
	start := time.Now()
	retrieveLatLongFromAddress(adapter, gmap)
	L.TimeTrack(start, "ImportHouseLocation")
}
