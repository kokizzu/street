package zImport

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"street/model/mProperty/wcProperty"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/mitchellh/mapstructure"
)

type PropertyFullResponse struct {
	AmenitiesInfo       AmenitiesInfo
	PublicRecordsInfo   PublicRecordsInfo
	PropertyHistoryInfo PropertyHistoryInfo
	ZoningDataInfo      ZoningDataInfo
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
	MediaBrowserInfoBySourceId map[string]interface{}
	PriceEstimates             PriceEstimates
	SectionPreviewText         string
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

func fetchPropertyUSByPropID(baseUrl string, propertyIdNum int) (map[any]interface{}, error) {

	// URL to fetch
	// url := "https://www.redfin.com/stingray/api/home/details/belowTheFold?propertyId=10000000&accessLevel=1"

	propertyId := strconv.FormatInt(int64(propertyIdNum), 10)

	url := fmt.Sprintf("%s?propertyId=%s&accessLevel=1", baseUrl, propertyId)
	// fmt.Println("Full URL => ", url)

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
	var data map[interface{}]interface{}

	// Parse the JSON response for us property
	if err := json.Unmarshal([]byte(formattedResponseStr), &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return nil, errors.New("Unexpected response json format for propertyID: " + propertyId)
	}

	return data, nil
}

func parseMapIntoStruct(inputMap map[string]interface{}) (*PropertyFullResponse, error) {
	var result PropertyFullResponse

	// Create a new mapstructure decoder
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &result, // Pass a pointer to the struct where you want to store the data
		// Pass some configuration
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

func ParsePropertyData(propertyMutator wcProperty.PropertyUsaMutator, propertyResponse map[string]interface{}) (res *wcProperty.PropertyUsaMutator) {

	// propertyMutator := wcProperty.NewPropertyUsaMutator(adapter)

	propertyResponseObject, err := parseMapIntoStruct(propertyResponse)
	if err != nil {
		fmt.Println("Error => ", err)
	}

	// -------- Basic info --------
	propertyMutator.Street = propertyResponseObject.PublicRecordsInfo.AddressInfo.Street
	propertyMutator.City = propertyResponseObject.PublicRecordsInfo.AddressInfo.City
	propertyMutator.State = propertyResponseObject.PublicRecordsInfo.AddressInfo.State
	propertyMutator.Zip = propertyResponseObject.PublicRecordsInfo.AddressInfo.Zip
	propertyMutator.CountryCode = propertyResponseObject.PublicRecordsInfo.AddressInfo.CountryCode

	propertyMutator.Bathroom = propertyResponseObject.PublicRecordsInfo.BasicInfo.Baths
	propertyMutator.Bedroom = propertyResponseObject.PublicRecordsInfo.BasicInfo.Beds
	propertyMutator.PropertyTypeName = propertyResponseObject.PublicRecordsInfo.BasicInfo.PropertyTypeName
	propertyMutator.YearBuilt = propertyResponseObject.PublicRecordsInfo.BasicInfo.YearBuilt
	propertyMutator.YearRenovated = propertyResponseObject.PublicRecordsInfo.BasicInfo.YearRenovated
	propertyMutator.TotalSqft = float64(propertyResponseObject.PublicRecordsInfo.BasicInfo.TotalSqft)
	propertyMutator.Apn = propertyResponseObject.PublicRecordsInfo.BasicInfo.Apn
	propertyMutator.PropertyLastUpdatedDate = propertyResponseObject.PublicRecordsInfo.BasicInfo.PropertyLastUpdatedDate
	propertyMutator.DisplayTimeZone = propertyResponseObject.PublicRecordsInfo.BasicInfo.DisplayTimeZone

	// -------- Tax info --------
	propertyMutator.TaxableLandValue = propertyResponseObject.PublicRecordsInfo.TaxInfo.TaxableLandValue
	propertyMutator.TaxableImprovementValue = propertyResponseObject.PublicRecordsInfo.TaxInfo.TaxableImprovementValue
	propertyMutator.RollYear = propertyResponseObject.PublicRecordsInfo.TaxInfo.RollYear
	propertyMutator.TaxesDue = propertyResponseObject.PublicRecordsInfo.TaxInfo.TaxesDue
	propertyMutator.TaxNote = propertyResponseObject.PublicRecordsInfo.SectionPreviewText

	// -------- County --------
	propertyMutator.CountyUrl = propertyResponseObject.PublicRecordsInfo.CountyUrl
	propertyMutator.CountyName = propertyResponseObject.PublicRecordsInfo.CountyName
	propertyMutator.CountyIsActive = propertyResponseObject.PublicRecordsInfo.CountyIsActive

	// -------- Build amenity list --------
	amenityGroupString, err := json.Marshal(propertyResponseObject.AmenitiesInfo.SuperGroups)

	if err != nil {
		fmt.Println("Failed to convert amenities group !")
	}

	propertyMutator.AmenitySuperGroups = string(amenityGroupString)
	propertyMutator.Coord = []any{0, 0}

	// -------- Build source photo (media-source) --------
	mediaSources := propertyResponseObject.PropertyHistoryInfo.MediaBrowserInfoBySourceId
	var mediaSourceList []interface{}

	var count int64 = 0 // Get latest lat/long from media sources
	for key, value := range mediaSources {
		// fmt.Println("Key: %s, Value: %v\n", key, value)

		var lat float64
		var long float64
		var streetViewUrl string
		var streetViewAvailable bool
		var altTextForImage string
		var assembledAddress string

		mediaObject, ok := value.(map[string]interface{})

		if ok {
			streetViewData, convertedOk1 := mediaObject["streetView"].(map[string]interface{})
			if convertedOk1 {
				latLong, convertedOk2 := streetViewData["latLong"].(map[string]interface{})

				if convertedOk2 {
					lat = latLong["latitude"].(float64)
					long = latLong["longitude"].(float64)

					// Store latest lat/long for property
					if count == 0 {
						propertyMutator.Coord = []any{lat, long}
					}

				} else {
					fmt.Println("\nCant parse lat/long")
				}

				streetViewAvailableData, ok := streetViewData["streetViewAvailable"].(bool)
				if ok {
					streetViewAvailable = streetViewAvailableData
				} else {
					fmt.Println("Unable to parse street view available flag !")
				}

				if streetViewAvailable == true {
					streetViewData, ok := streetViewData["streetViewUrl"].(string)
					if ok {
						streetViewUrl = streetViewData
					}
				}

			} else {
				fmt.Println("Cant parse street view url")
			}
			altTextForImageData, ok := mediaObject["altTextForImage"].(string)
			if ok {
				altTextForImage = altTextForImageData
			}

			assembledAddressData, ok := mediaObject["assembledAddress"].(string)
			if ok {
				assembledAddress = assembledAddressData
			}

		} else {
			fmt.Println("No media data for this property")
		}

		mediaItem := map[string]interface{}{
			"id":                  key,
			"latitude":            lat,
			"longitude":           long,
			"streetViewUrl":       streetViewUrl,
			"streetViewAvailable": streetViewAvailable,
			"altTextForImage":     altTextForImage,
			"assembledAddress":    assembledAddress,
		}
		mediaSourceList = append(mediaSourceList, mediaItem)
		count++
	}

	if mediaSourceList == nil || len(mediaSourceList) > 0 {
		propertyMutator.MediaSource = []any{}
	} else {
		propertyMutator.MediaSource = mediaSourceList
	}

	// -------- Zone data info --------
	propertyMutator.ZoneName = propertyResponseObject.ZoningDataInfo.ZoneName
	propertyMutator.ZoneType = propertyResponseObject.ZoningDataInfo.ZoneType.ZoneType
	propertyMutator.ZoneSubType = propertyResponseObject.ZoningDataInfo.ZoneType.ZoneSubType
	propertyMutator.ZoneCode = propertyResponseObject.ZoningDataInfo.ZoneCode

	// --------  Build permitted land use --------
	permittedLandUse, err := json.Marshal(propertyResponseObject.ZoningDataInfo.PermittedLandUse)

	if err != nil {
		fmt.Println("Failed to convert permittedLandUse !")
	}
	propertyMutator.PermittedLandUse = string(permittedLandUse)

	notPermittedLandUseArr, err := json.Marshal(propertyResponseObject.ZoningDataInfo.NotPermittedLandUse)

	if err != nil {
		fmt.Println("Failed to convert notPermittedLandUseArr !")
	}
	propertyMutator.NotPermittedLandUse = string(notPermittedLandUseArr)

	// -------- Store agent & broker --------
	propertyMutator.ListingBrokerName = propertyResponseObject.AmenitiesInfo.MlsDisclaimerInfo.ListingBrokerName
	propertyMutator.ListingBrokerNumber = propertyResponseObject.AmenitiesInfo.MlsDisclaimerInfo.ListingBrokerNumber
	propertyMutator.ListingAgentName = propertyResponseObject.AmenitiesInfo.MlsDisclaimerInfo.ListingAgentName
	propertyMutator.ListingAgentNumber = propertyResponseObject.AmenitiesInfo.MlsDisclaimerInfo.ListingAgentNumber

	// -------- Metadata ---------
	currentTime := time.Now().Unix()
	propertyMutator.CreatedAt = currentTime
	propertyMutator.UpdatedAt = currentTime

	return &propertyMutator
}

func ImportPropertyUsData(adapter **Tt.Adapter, baseUrl string, maxPropertyID int) {

	stat := &ImporterStat{Total: maxPropertyID}
	defer stat.Print(`last`)

	for i := 1; i <= maxPropertyID; i++ {
		stat.Print()

		fmt.Println("Property ID => ", i)

		propertyMutator := wcProperty.NewPropertyUsaMutator(*adapter)
		propertyMutator.PropertyId = uint64(i)

		if propertyMutator.FindByPropertyId() {
			stat.Skip()
			continue
		}

		data, err := fetchPropertyUSByPropID(baseUrl, i)
		if err != nil {
			fmt.Println("Error has happen for property ID -> ", i)
		}

		propertyResponse := data["payload"].(map[string]interface{})
		propertyVersion := data["version"].(float64)

		propertyMutator = ParsePropertyData(*propertyMutator, propertyResponse)
		propertyMutator.Version = propertyVersion

		stat.Ok(propertyMutator.DoInsert())
	}

}
