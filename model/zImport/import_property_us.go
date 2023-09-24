package zImport

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kokizzu/gotro/X"

	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/mitchellh/mapstructure"
)

type PropertyFullResponse struct {
	AmenitiesInfo           AmenitiesInfo
	PublicRecordsInfo       PublicRecordsInfo
	PropertyHistoryInfo     PropertyHistoryInfo
	ZoningDataInfo          ZoningDataInfo
	SchoolsAndDistrictsInfo map[string]any
	RiskFactorData          map[string]any
}

type AmenitiesInfo struct {
	MlsDisclaimerInfo MlsDisclaimerInfo
	SuperGroups       []AmenitySuperGroups
}

type MlsDisclaimerInfo struct {
	ListingBrokerName   string
	ListingBrokerNumber string
	ListingAgentName    string
	ListingAgentNumber  string
}

type AmenitySuperGroups struct {
	Types         []int64
	TitleString   string
	AmenityGroups []AmenityGroups
}

type AmenityGroups struct {
	GroupTitle     string
	ReferenceName  string
	AmenityEntries []AmenityEntries
}

type AmenityEntries struct {
	AmenityName   string
	ReferenceName string
	AmenityValues []string
}

type PublicRecordsInfo struct {
	BasicInfo          BasicInfo
	AddressInfo        AddressInfo
	TaxInfo            TaxInfo
	CountyUrl          string
	CountyName         string
	CountyIsActive     bool
	SectionPreviewText string
}

type ZoneType struct {
	ZoneType    string
	ZoneSubType string
	Display     []string
}

type ZoningDataInfo struct {
	ZoneName            string
	ZoneType            ZoneType
	ZoneCode            string
	PermittedLandUse    []string
	NotPermittedLandUse []string
}

type AddressInfo struct {
	Street      string
	City        string
	State       string
	Zip         string
	CountryCode string
}

type BasicInfo struct {
	Beds                    int64
	Baths                   int64
	PropertyTypeName        string
	YearBuilt               int64
	YearRenovated           int64
	TotalSqft               int64
	Apn                     string
	PropertyLastUpdatedDate int64
	DisplayTimeZone         string
}

type TaxInfo struct {
	TaxableLandValue        int64
	TaxableImprovementValue int64
	RollYear                int64
	TaxesDue                float64
}

type PropertyHistoryInfo struct {
	HasPropertyHistory         bool
	MediaBrowserInfoBySourceId map[string]any
	PriceEstimates             PriceEstimates
	SectionPreviewText         string
	Events                     []PropertyHistoryEvent
}

type PropertyHistoryEvent struct {
	Price            int64
	EventDescription string
	SourceId         string // suppose it's transaction ID
	EventDate        int64
	EventDateString  string
}

type PriceEstimates struct {
	PriceHomeUrl string
}

type StreetView struct {
	LatLong             LatLong
	StreetViewUrl       string
	StreetViewAvailable bool
}

type LatLong struct {
	latitude  float64
	longitude float64
}

const (
	SELL    = "SELL"
	LISTED  = "LISTED"
	RENT    = "RENT"
	UNKNOWN = "UNKNOWN"
)

func fetchPropertyUSByPropID(baseUrl string, propertyIdNum int) (map[any]any, error) {

	// URL to fetch
	// url := "https://www.redfin.com/stingray/api/home/details/belowTheFold?propertyId=10000000&accessLevel=1"

	propertyId := strconv.FormatInt(int64(propertyIdNum), 10)

	url := fmt.Sprintf("%s?propertyId=%s&accessLevel=1", baseUrl, propertyId)

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making HTTP GET request: %v\n", err)
		return nil, errors.New("Error making HTTP GET request")
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("HTTP GET request returned status: %v\n", response.Status)
		return nil, errors.New("Some error for fetching this property ID " + propertyId)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, errors.New("Some error for fetching this property ID " + propertyId)
	}
	responseString := string(body)
	formattedResponseStr := strings.Replace(responseString, "{}&&", "", -1)

	// Define a struct to parse the JSON response into
	var data map[any]any

	// Parse the JSON response for us property
	if err := json.Unmarshal([]byte(formattedResponseStr), &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return nil, errors.New("Unexpected response json format for propertyID: " + propertyId)
	}

	return data, nil
}

func parseMapIntoStruct(inputMap map[string]any) (*PropertyFullResponse, error) {
	var result PropertyFullResponse

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &result,
	})

	if err != nil {
		return nil, err
	}

	// Decode the map into the struct
	if err := decoder.Decode(inputMap); err != nil {
		return nil, err
	}

	return &result, nil
}

func parsePropertyExtraData(propertyMutator *wcProperty.PropertyExtraUSMutator, propertyResponseObject *PropertyFullResponse) {

	// -------- Tax info --------
	taxInfoJson, err := json.Marshal(propertyResponseObject.PublicRecordsInfo.TaxInfo)
	if err != nil {
		L.Print("Failed to convert tax info !")
	}
	propertyMutator.TaxInfo = string(taxInfoJson)
	propertyMutator.TaxNote = propertyResponseObject.PublicRecordsInfo.SectionPreviewText

	// -------- County --------
	propertyMutator.CountyUrl = propertyResponseObject.PublicRecordsInfo.CountyUrl
	propertyMutator.CountyIsActive = propertyResponseObject.PublicRecordsInfo.CountyIsActive

	// -------- Build amenity list --------
	amenityGroupString, err := json.Marshal(propertyResponseObject.AmenitiesInfo.SuperGroups)
	if err != nil {
		L.Print("Failed to convert amenities group !")
	}

	propertyMutator.AmenitySuperGroups = string(amenityGroupString)

	// -------- Build source photo (media-source) --------
	mediaSourceJson, err := json.Marshal(propertyResponseObject.PropertyHistoryInfo.MediaBrowserInfoBySourceId)
	if err != nil {
		L.Print("Failed to convert media sources !")
	}
	propertyMutator.MediaSourceJson = string(mediaSourceJson)

	// -------- Zone data info & Permitted land to use --------
	zoneDataInfo, err := json.Marshal(propertyResponseObject.ZoningDataInfo)
	if err != nil {
		L.Print("Can't parse Zone Data Info")
	}
	propertyMutator.ZoneDataInfo = string(zoneDataInfo)

	// // -------- Store agent & broker --------
	mlsDisclaimerInfo, err := json.Marshal(propertyResponseObject.AmenitiesInfo.MlsDisclaimerInfo)
	if err != nil {
		L.Print("Can't parse Zone Data Info")
	}
	propertyMutator.MlsDisclaimerInfo = string(mlsDisclaimerInfo)

	// // --------Facility Info --------
	facilityInfo := propertyResponseObject.SchoolsAndDistrictsInfo
	facilityInfoJson, err := json.Marshal(facilityInfo)
	if err != nil {
		L.Print("Can't parse Zone Data Info")
	}
	propertyMutator.FacilityInfo = string(facilityInfoJson)

	// // --------Risk Info --------
	riskInfo := propertyResponseObject.RiskFactorData
	riskInfoJson, err := json.Marshal(riskInfo)
	if err != nil {
		L.Print("Can't parse Zone Data Info")
	}
	propertyMutator.RiskInfo = string(riskInfoJson)
}

func parsePropertyData(propertyMutator *wcProperty.PropertyUSMutator, propertyResponseObject *PropertyFullResponse, stat *ImporterStat) {

	// // -------- Basic info --------
	propertyMutator.Street = propertyResponseObject.PublicRecordsInfo.AddressInfo.Street
	propertyMutator.City = propertyResponseObject.PublicRecordsInfo.AddressInfo.City
	propertyMutator.State = propertyResponseObject.PublicRecordsInfo.AddressInfo.State
	propertyMutator.Zip = propertyResponseObject.PublicRecordsInfo.AddressInfo.Zip

	propertyMutator.Bathroom = propertyResponseObject.PublicRecordsInfo.BasicInfo.Baths
	propertyMutator.Bedroom = propertyResponseObject.PublicRecordsInfo.BasicInfo.Beds
	propertyMutator.HouseType = propertyResponseObject.PublicRecordsInfo.BasicInfo.PropertyTypeName
	propertyMutator.YearBuilt = propertyResponseObject.PublicRecordsInfo.BasicInfo.YearBuilt // Same with complete construct date

	propertyMutator.YearRenovated = propertyResponseObject.PublicRecordsInfo.BasicInfo.YearRenovated
	propertyMutator.TotalSqft = float64(propertyResponseObject.PublicRecordsInfo.BasicInfo.TotalSqft)
	propertyMutator.SizeM2 = fmt.Sprintf("%.3f", convertSqftToM2(propertyMutator.TotalSqft))
	propertyMutator.SerialNumber = propertyResponseObject.PublicRecordsInfo.BasicInfo.Apn // Store as serial number
	propertyMutator.PropertyLastUpdatedDate = propertyResponseObject.PublicRecordsInfo.BasicInfo.PropertyLastUpdatedDate

	propertyMutator.CountyName = propertyResponseObject.PublicRecordsInfo.CountyName

	// Address
	propertyMutator.Address = propertyMutator.Street + "," + propertyMutator.City + "," + propertyMutator.CountyName + "," + propertyMutator.State
	propertyMutator.FormattedAddress = propertyMutator.Street + "," + propertyMutator.City + "," + propertyMutator.CountyName +
		"," + propertyMutator.State + ", Zip: " + propertyMutator.Zip + ", US"

	// Get latest coord from media resources
	propertyMutator.Coord = []any{0, 0}
	mediaSources := propertyResponseObject.PropertyHistoryInfo.MediaBrowserInfoBySourceId

	for _, value := range mediaSources {

		var lat float64
		var long float64

		mediaObject, ok := value.(map[string]any)

		if ok {
			streetViewData, convertedOk1 := mediaObject["streetView"].(map[string]any)
			if convertedOk1 {
				latLong, convertedOk2 := streetViewData["latLong"].(map[string]any)

				if convertedOk2 {
					lat = latLong["latitude"].(float64)
					long = latLong["longitude"].(float64)

					// Store latest lat/long for property
					if propertyMutator.Coord[0] == 0 && propertyMutator.Coord[1] == 0 {
						propertyMutator.Coord = []any{lat, long}
					}
				} else {
					stat.Warn(`no lat long`)
				}
			}
		}
	}

	// -------- Metadata ---------
	currentTime := time.Now().Unix()
	propertyMutator.CreatedAt = currentTime
	propertyMutator.UpdatedAt = currentTime
}

func ImportPropertyHistories(propertyResponseObject *PropertyFullResponse, redFinId string) []rqProperty.PropertyHistory {

	if propertyResponseObject.PropertyHistoryInfo.HasPropertyHistory == false {
		return []rqProperty.PropertyHistory{}
	}

	propHistories := []rqProperty.PropertyHistory{}

	for _, historyItem := range propertyResponseObject.PropertyHistoryInfo.Events {

		// Tracking only history have price
		if historyItem.Price == 0 && strings.ContainsAny(historyItem.EventDescription, "Pending") {
			continue
		}

		data := rqProperty.PropertyHistory{}

		data.TransactionTime = strconv.FormatInt(historyItem.EventDate, 10)
		data.TransactionKey = historyItem.SourceId
		data.Price = historyItem.Price
		data.PropertyKey = redFinId
		data.TransactionDescription = historyItem.EventDescription
		data.TransactionDateNormal = historyItem.EventDateString
		data.SerialNumber = propertyResponseObject.PublicRecordsInfo.BasicInfo.Apn // Serial number of house

		// Check sell transaction
		if strings.ContainsAny(historyItem.EventDescription, "Sold") || strings.ContainsAny(historyItem.EventDescription, "Listed") {
			data.TransactionType = SELL
		} else {
			data.TransactionType = UNKNOWN
		}

		propHistories = append(propHistories, data)
	}

	return propHistories

}

func convertSqftToM2(sqft float64) float64 {
	return sqft * 0.092903
}

func SavePropertyHistories(adapter *Tt.Adapter, propList []rqProperty.PropertyHistory) {

	stat := &ImporterStat{Total: len(propList)}
	defer stat.Print(`last`)

	for _, pHistory := range propList {

		stat.Print()

		propertyHistoryMutator := wcProperty.NewPropertyHistoryMutator(adapter)
		propertyHistoryMutator.Adapter = adapter
		propertyHistoryMutator.PropertyHistory = pHistory
		propertyHistoryMutator.UpdatedAt = time.Now().Unix()
		stat.Ok(propertyHistoryMutator.DoInsert())
	}
}

func ImportPropertyUsData(adapter *Tt.Adapter) {
	const baseUrl = "https://www.redfin.com/stingray/api/home/details/belowTheFold"
	const minPropertyId = 1
	const maxPropertyId = 10000000

	stat := &ImporterStat{Total: maxPropertyId * 2, PrintEvery: 10}
	defer stat.Print(`last`)

	for i := minPropertyId; i <= maxPropertyId; i++ {
		stat.Print()

		redfinKey := `rf` + strconv.Itoa(i)

		// check first whether exists
		propertyExtraMutator := wcProperty.NewPropertyExtraUSMutator(adapter)
		propertyExtraMutator.PropertyKey = redfinKey
		if propertyExtraMutator.FindByPropertyKey() {
			stat.Skip()
			continue
		}

		propertyMutator := wcProperty.NewPropertyUSMutator(adapter)
		propertyMutator.UniqPropKey = redfinKey
		propertyMutator.FindByUniqPropKey()

		data, err := fetchPropertyUSByPropID(baseUrl, i)
		L.PanicIf(err, "Error: fetchPropertyUSByPropID "+redfinKey)

		propertyResponse := data["payload"].(map[string]any)
		propertyVersion := X.ToS(data["version"])

		// Property into the struct
		propertyResponseObject, err := parseMapIntoStruct(propertyResponse)
		L.PanicIf(err, "Parse property response error")

		parsePropertyData(propertyMutator, propertyResponseObject, stat)

		propertyMutator.Version = propertyVersion

		propertyHistories := ImportPropertyHistories(propertyResponseObject, strconv.Itoa(i))

		if len(propertyHistories) == 0 {
			stat.Warn("no prop history")
			propertyMutator.LastPrice = `0`
		} else {
			// Update last price
			propertyMutator.LastPrice = fmt.Sprint(propertyHistories[0].Price)
		}
		stat.Ok(propertyMutator.DoInsert())

		// Store extra
		parsePropertyExtraData(propertyExtraMutator, propertyResponseObject)
		stat.Ok(propertyExtraMutator.DoInsert())

		// TODO: save property history
		//SavePropertyHistories(adapter, propertyHistories)

	}

}
