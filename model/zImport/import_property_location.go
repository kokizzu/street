package zImport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/L"

	"github.com/kokizzu/gotro/D/Tt"
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

func buildFullLocationSearchUrl(apiKey string, address string) string {
	const googleApiUrl = `https://maps.googleapis.com/maps/api/place/findplacefromtext/json`

	listFields := "formatted_address,name,rating,opening_hours,geometry"
	inputType := "textquery"
	req, err := http.NewRequest("GET", googleApiUrl, nil)
	if err != nil {
		L.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("fields", listFields)
	q.Add("input", address)
	q.Add("inputtype", inputType)
	q.Add("key", apiKey)
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func retrieveLatLongFromAddress(adapter *Tt.Adapter, apiKey string) {
	defer subTaskPrint(`retrieveLatLongFromAddress: retrieve lat/long`)()

	propertyMutator := wcProperty.NewPropertyMutator(adapter)
	properties := propertyMutator.FindAllProperties()

	stat := &ImporterStat{Total: len(properties), PrintEvery: 10}
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

		fullUrl := buildFullLocationSearchUrl(apiKey, p.Address)
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

func ImportHouseLocation(adapter *Tt.Adapter) {
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		fmt.Println("Require google api key to execute this operation")
		return
	}

	start := time.Now()
	retrieveLatLongFromAddress(adapter, googleApiKey)
	L.TimeTrack(start, "ImportHouseLocation")
}
