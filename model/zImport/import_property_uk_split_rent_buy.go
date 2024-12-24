package zImport

import (
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

	for _, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyMutator(conn)
		if !isValidURL(v.SourceURL) {
			stat.Skip()
			continue
		}

		propKey, err := getSheetPropertyPTUniqPropKey(v.SourceURL)
		if err != nil {
			stat.Skip()
			continue
		}

		// Find by uniq prop key
		property.SetUniqPropKey(propKey + `_pt`)
		property.FindByUniqPropKey()

		// Update to UK
		property.SetUniqPropKey(propKey + `_uk`)
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
		images := S.Split(v.Image, `|`)
		imagesAny := convertArrStrToArrAny(images)
		property.SetImages(imagesAny)

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		stat.Ok(property.DoUpsert())
	}
}
