package zImport

import (
	"fmt"
	"os"
	"path/filepath"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

func ReadPropertyUK_Rent_RightmoveCoUk(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyUK_Rent_RightmoveCoUk: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}
	defer file.Close()

	type propertyUK struct {
		SourceURL       string
		PropertyTitle   string
		Location        string
		Description     string
		PhoneNumber     string
		PropertyPurpose string
		PrimaryPrice    string
		SecondaryPrice  string
		Date            string
		PropertyType    string
		Bedrooms        string
		Bathrooms       string
		SizeSqFt        string
		AgentName       string
		DisplayAddress  string
		Lat             string
		Long            string
		Bio             string
		KeyFeatures     string
		Image           string
	}

	var properties []propertyUK

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceURL := tsv.String()
		propertyTitle := tsv.String()
		location := tsv.String()
		description := tsv.String()
		phoneNumber := tsv.String()
		purpose := tsv.String()
		primaryPrice := tsv.String()
		secondaryPrice := tsv.String()
		date := tsv.String()
		propertyType := tsv.String()
		bedrooms := tsv.String()
		bathrooms := tsv.String()
		sizeSqFt := tsv.String()
		agentName := tsv.String()
		displayAddress := tsv.String()
		latitude := tsv.String()
		longitude := tsv.String()
		bio := tsv.String()
		keyFeatures := tsv.String()
		image := tsv.String()

		properties = append(properties, propertyUK{
			SourceURL:       S.Trim(sourceURL),
			PropertyTitle:   S.Trim(propertyTitle),
			Location:        S.Trim(location),
			Description:     S.Trim(description),
			PhoneNumber:     S.Trim(phoneNumber),
			PropertyPurpose: S.Trim(purpose),
			PrimaryPrice:    S.Trim(primaryPrice),
			SecondaryPrice:  S.Trim(secondaryPrice),
			Date:            S.Trim(date),
			PropertyType:    S.Trim(propertyType),
			Bedrooms:        S.Trim(bedrooms),
			Bathrooms:       S.Trim(bathrooms),
			SizeSqFt:        S.Trim(sizeSqFt),
			AgentName:       S.Trim(agentName),
			DisplayAddress:  S.Trim(displayAddress),
			Lat:             S.Trim(latitude),
			Long:            S.Trim(longitude),
			Bio:             S.Trim(bio),
			KeyFeatures:     S.Trim(keyFeatures),
			Image:           S.Trim(image),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:]
	} else {
		panic(`Properties from ` + resourcePath + ` is empty`)
	}

	stat := &ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	for idx, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyMutator(conn)
		if !isValidURL(v.SourceURL) {
			stat.Skip()
			continue
		}

		property.SetPurpose(mProperty.PropertyPurposeRent)
		property.SetAddress(v.Location)
		property.SetCity(`London`)
		property.SetCountryCode(`UK`)

		jsonNote := getJsonStrPropertyNote(``, v.PhoneNumber, v.Description)
		property.SetNote(jsonNote)

		priceUSD := convertGBPToUSD(v.PrimaryPrice)
		property.SetLastPrice(priceUSD)
		property.SetHouseType(v.PropertyType)
		property.SetBedroom(S.ToI(v.Bedrooms))
		property.SetBathroom(S.ToI(v.Bathrooms))

		sizeSqFt := v.SizeSqFt
		sizeSqFt = S.Replace(sizeSqFt, `sq ft`, ``)
		sizeSqFt = S.Replace(sizeSqFt, `,`, ``)
		property.SetTotalSqft(S.ToF(sizeSqFt))

		property.SetFormattedAddress(v.DisplayAddress)

		// Set coordinate
		coord := getLatLong(v.Lat, v.Long)
		property.SetCoord(coord[:])

		// Set images url
		imageRaw := v.Image
		imageRaw = S.Replace(imageRaw, `"}`, ``)
		images := S.Split(imageRaw, `|`)
		imagesAny := convertArrStrToArrAny(images)
		property.SetImages(imagesAny)

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		propKey, err := getSheetPropertyPTUniqPropKey(v.SourceURL)
		if err != nil {
			stat.Skip()
			fmt.Printf("\033[32m[%d] Skipped (uniqPropKey not found from source url)\033[0m\n", idx+1)
			continue
		}

		// Find by uniq prop key
		// Override from '_pt' to '_uk' if exists
		property.UniqPropKey = propKey + `_pt`
		if property.FindByUniqPropKey() {
			fmt.Printf("\033[34m[%d] Override (uniqPropKey '%s_pt' to '%s_uk')\033[0m\n", idx+1, propKey, propKey)
			property.SetUniqPropKey(propKey + `_uk`)
			// Update by ID because uniqPropKey will be modified from '_pt' to '_uk'
			if !property.DoUpdateById() {
				stat.Skip()
				fmt.Printf("\033[31m[%d] Skipped (property not updated by id)\033[0m\n", idx+1)
				continue
			}
			stat.Ok(true)
			continue
		}

		// Find by uniq prop key
		// Skip if exists
		property.UniqPropKey = propKey + `_uk`
		if property.FindByUniqPropKey() {
			fmt.Printf("\033[32m[%d] Override (uniqPropKey %s already exists)\033[0m\n", idx+1, property.UniqPropKey)
			if !property.DoUpdateByUniqPropKey() {
				stat.Skip()
				fmt.Printf("\033[31m[%d] Skipped (property not updated by uniqPropKey)\033[0m\n", idx+1)
				continue
			}
			stat.Ok(true)
			continue
		}

		property.SetUniqPropKey(propKey + `_uk`)

		if !property.DoInsert() {
			stat.Skip()
			fmt.Printf("\033[31m[%d] Skipped (property not inserted)\033[0m\n", idx+1)
			continue
		}

		stat.Ok(true)
	}
}

func ReadPropertyUK_Sale_RightmoveCoUk(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyUK_Sale_RightmoveCoUk: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}
	defer file.Close()

	type propertyUK struct {
		SourceURL       string
		PropertyTitle   string
		Location        string
		Description     string
		PhoneNumber     string
		PropertyPurpose string
		PrimaryPrice    string
		Date            string
		PropertyType    string
		Bedrooms        string
		Bathrooms       string
		SizeSqFt        string
		Tenure          string
		AgentName       string
		DisplayAddress  string
		Lat             string
		Long            string
		Bio             string
		KeyFeatures     string
		Image           string
	}

	var properties []propertyUK

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceURL := tsv.String()
		propertyTitle := tsv.String()
		location := tsv.String()
		description := tsv.String()
		phoneNumber := tsv.String()
		purpose := tsv.String()
		primaryPrice := tsv.String()
		date := tsv.String()
		propertyType := tsv.String()
		bedrooms := tsv.String()
		bathrooms := tsv.String()
		sizeSqFt := tsv.String()
		tenure := tsv.String()
		agentName := tsv.String()
		displayAddress := tsv.String()
		latitude := tsv.String()
		longitude := tsv.String()
		bio := tsv.String()
		keyFeatures := tsv.String()
		image := tsv.String()

		properties = append(properties, propertyUK{
			SourceURL:       S.Trim(sourceURL),
			PropertyTitle:   S.Trim(propertyTitle),
			Location:        S.Trim(location),
			Description:     S.Trim(description),
			PhoneNumber:     S.Trim(phoneNumber),
			PropertyPurpose: S.Trim(purpose),
			PrimaryPrice:    S.Trim(primaryPrice),
			Date:            S.Trim(date),
			PropertyType:    S.Trim(propertyType),
			Bedrooms:        S.Trim(bedrooms),
			Bathrooms:       S.Trim(bathrooms),
			SizeSqFt:        S.Trim(sizeSqFt),
			Tenure:          S.Trim(tenure),
			AgentName:       S.Trim(agentName),
			DisplayAddress:  S.Trim(displayAddress),
			Lat:             S.Trim(latitude),
			Long:            S.Trim(longitude),
			Bio:             S.Trim(bio),
			KeyFeatures:     S.Trim(keyFeatures),
			Image:           S.Trim(image),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:]
	} else {
		panic(`Properties from ` + resourcePath + ` is empty`)
	}

	stat := &ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	for idx, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyMutator(conn)
		if !isValidURL(v.SourceURL) {
			stat.Skip()
			fmt.Printf("\033[32m[%d] Skipped (invalid url) ...\033[0m\n", idx+1)
			continue
		}

		property.SetPurpose(mProperty.PropertyPurposeSale)
		property.SetAddress(v.Location)
		property.SetCity(`London`)
		property.SetCountryCode(`UK`)

		jsonNote := getJsonStrPropertyNote(``, v.PhoneNumber, v.Description)
		property.SetNote(jsonNote)

		priceUSD := convertGBPToUSD(v.PrimaryPrice)
		property.SetLastPrice(priceUSD)
		property.SetHouseType(v.PropertyType)
		property.SetBedroom(S.ToI(v.Bedrooms))
		property.SetBathroom(S.ToI(v.Bathrooms))

		sizeSqFt := v.SizeSqFt
		sizeSqFt = S.Replace(sizeSqFt, `sq ft`, ``)
		sizeSqFt = S.Replace(sizeSqFt, `,`, ``)
		property.SetTotalSqft(S.ToF(sizeSqFt))

		property.SetFormattedAddress(v.DisplayAddress)

		// Set coordinate
		coord := getLatLong(v.Lat, v.Long)
		property.SetCoord(coord[:])

		// Set images url
		imageRaw := v.Image
		imageRaw = S.Replace(imageRaw, `"}`, ``)
		images := S.Split(imageRaw, `|`)
		imagesAny := convertArrStrToArrAny(images)
		property.SetImages(imagesAny)

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		propKey, err := getSheetPropertyPTUniqPropKey(v.SourceURL)
		if err != nil {
			stat.Skip()
			fmt.Printf("\033[32m[%d] Skipped (uniqPropKey not found from source url)\033[0m\n", idx+1)
			continue
		}

		// Find by uniq prop key
		// Override from '_pt' to '_uk' if exists
		property.UniqPropKey = propKey + `_pt`
		if property.FindByUniqPropKey() {
			fmt.Printf("\033[34m[%d] Override (uniqPropKey '%s_pt' to '%s_uk')\033[0m\n", idx+1, propKey, propKey)
			property.SetUniqPropKey(propKey + `_uk`)
			// Update by ID because uniqPropKey will be modified from '_pt' to '_uk'
			if !property.DoUpdateById() {
				stat.Skip()
				fmt.Printf("\033[31m[%d] Skipped (property not updated by id)\033[0m\n", idx+1)
				continue
			}
			stat.Ok(true)
			continue
		}

		// Find by uniq prop key
		// Skip if exists
		property.UniqPropKey = propKey + `_uk`
		if property.FindByUniqPropKey() {
			fmt.Printf("\033[32m[%d] Override (uniqPropKey %s already exists)\033[0m\n", idx+1, property.UniqPropKey)
			if !property.DoUpdateByUniqPropKey() {
				stat.Skip()
				fmt.Printf("\033[31m[%d] Skipped (property not updated by uniqPropKey)\033[0m\n", idx+1)
				continue
			}
			stat.Ok(true)
			continue
		}

		property.SetUniqPropKey(propKey + `_uk`)

		if !property.DoInsert() {
			stat.Skip()
			fmt.Printf("\033[31m[%d] Skipped (property not inserted)\033[0m\n", idx+1)
			continue
		}

		stat.Ok(true)
	}
}
