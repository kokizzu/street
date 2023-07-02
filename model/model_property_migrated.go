package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"street/model/mProperty/wcProperty"
	"time"

	"github.com/kokizzu/gotro/D/Tt"
)

// A Place Response
type PlaceResponse struct {
	Candidates []Candidate `json:"candidates"`
	Status     string      `json:"status"`
}

// Location
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// View Port
type ViewPort struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// Geometry model
type Geometry struct {
	Location Location `json:"location"`
	ViewPort ViewPort `json:"viewport"`
}

type Candidate struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

func buildFullPlaceSearchUrl(apiKey string, googleApiUrl string, address string) string {
	fields := "formatted_address%2Cname%2Crating%2Copening_hours%2Cgeometry"
	addressStr := address
	inputType := "textquery"
	googleApiKey := apiKey
	placeApiUrl := googleApiUrl + "/json?fields=" + fields + "&input=" + addressStr + "&inputtype=" + inputType + "&key=" + googleApiKey

	return placeApiUrl
}

func retrieveLatLongFromAddress(adapter *Tt.Adapter, apiKey string, googleApiUrl string) {

	fmt.Println("[Start] Beginning of process data to retrieve lat/long")
	propertyMutator := wcProperty.NewPropertyMutator(adapter)

	properties := propertyMutator.FindAllProperties()

	stat := &ImporterStat{Total: len(properties)}
	for index, p := range properties {
		stat.Print()

		if p.Address == "" {
			stat.Skip()
			continue
		}
		if p.FormattedAddress != "" {
			stat.Skip()
			fmt.Println("Property has address and lat/long already")
			continue
		}

		fullUrl := buildFullPlaceSearchUrl(apiKey, googleApiUrl, p.Address)
		fmt.Println("Full URL => ", fullUrl)
		locationResponse, err := http.Get(fullUrl)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(locationResponse.Body)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println("Response data => " + string(responseData))
		propertyLocation := PlaceResponse{}

		json.Unmarshal([]byte(responseData), &propertyLocation)
		fmt.Println("Response => ", propertyLocation.Candidates[0].FormattedAddress)
		fmt.Println("Response 2 => ", propertyLocation.Candidates[0].Geometry)
		if len(propertyLocation.Candidates) == 0 {
			stat.Skip()
			fmt.Println("There is no available location for this address")
			break
		}

		dataMutator := wcProperty.NewPropertyMutator(adapter)
		dataMutator.Property = *p
		dataMutator.Adapter = adapter
		dataMutator.FormattedAddress = propertyLocation.Candidates[0].FormattedAddress

		latitude := propertyLocation.Candidates[0].Geometry.Location.Lat
		longitude := propertyLocation.Candidates[0].Geometry.Location.Lng

		dataMutator.Coord = []any{latitude, longitude}
		dataMutator.UpdatedAt = time.Now().Unix()

		fmt.Println("PropertyKey = " + p.UniqPropKey)
		stat.Ok(dataMutator.DoOverwriteById())

		if index >= 10 {
			break
		}
	}

	fmt.Println("Total length of properties => ", len(properties))

	fmt.Println("[End] Beginning of process data to retrieve lat/long")
}

func MigratedPropertyLocation(adapter *Tt.Adapter) {

	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	if googleApiKey == "" {
		fmt.Println("Require google api key to execute this operation")
		return
	}

	googleApiUrl := os.Getenv("GOOGLE_API_URL")
	if googleApiUrl == "" {
		fmt.Println("Require google api url to execute this operation")
		return
	}

	retrieveLatLongFromAddress(adapter, googleApiKey, googleApiUrl)
}
