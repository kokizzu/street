package zImport

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"street/model"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

func ReadPropertyPT_RightmoveCoUk(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyPT_RightmoveCoUk: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()
	type propertyPT struct {
		SourceURL      string
		PropertyTitle  string
		Location       string
		PhoneNumber    string
		PrimaryPrice   string
		SecondaryPrice string
		PropertyType   string
		Bedrooms       string
		Bathrooms      string
		Size           string
		AgentName      string
		FullAddress    string
		Lat            string
		Long           string
		Images         string
		Description    string
		KeyFeatures    string
	}

	var properties []propertyPT

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceURL := tsv.String()
		propertyTitle := tsv.String()
		location := tsv.String()
		phoneNumber := tsv.String()
		primaryPrice := tsv.String()
		secondaryPrice := tsv.String()
		propertyType := tsv.String()
		bedrooms := tsv.String()
		bathrooms := tsv.String()
		size := tsv.String()
		agentName := tsv.String()
		fullAddress := tsv.String()
		lat := tsv.String()
		long := tsv.String()
		images := tsv.String()
		description := tsv.String()
		keyFeatures := tsv.String()

		properties = append(properties, propertyPT{
			SourceURL:      S.Trim(sourceURL),
			PropertyTitle:  S.Trim(propertyTitle),
			Location:       S.Trim(location),
			PhoneNumber:    S.Trim(phoneNumber),
			PrimaryPrice:   S.Trim(primaryPrice),
			SecondaryPrice: S.Trim(secondaryPrice),
			PropertyType:   S.Trim(propertyType),
			Bedrooms:       S.Trim(bedrooms),
			Bathrooms:      S.Trim(bathrooms),
			Size:           S.Trim(size),
			AgentName:      S.Trim(agentName),
			FullAddress:    S.Trim(fullAddress),
			Lat:            S.Trim(lat),
			Long:           S.Trim(long),
			Images:         S.Trim(images),
			Description:    S.Trim(description),
			KeyFeatures:    S.Trim(keyFeatures),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:] // remove 1st element
	} else {
		panic(`Properties from ` + resourcePath + ` is empty`)
	}

	stat := &model.ImporterStat{Total: len(properties)}
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

		property.UniqPropKey = propKey + `_pt`
		if property.FindByUniqPropKey() {
			fmt.Println(`Override property with uniq prop key (PT): `, property.UniqPropKey)
		} else {
			property.SetUniqPropKey(propKey + `_pt`)
		}

		purpose := getPropertyPurpose(v.PropertyTitle)
		property.SetPurpose(purpose)
		property.SetCity(`Lisbon`)

		priceUSD := convertGBPToUSD(v.PrimaryPrice)
		property.SetLastPrice(priceUSD)
		property.SetHouseType(v.PropertyType)
		property.SetBedroom(S.ToI(v.Bedrooms))
		property.SetBathroom(S.ToI(v.Bathrooms))

		sizeSqFt := getSizeSqFt(v.Size)
		property.SetTotalSqft(sizeSqFt)
		property.SetAddress(v.FullAddress)

		images := S.Split(v.Images, `|`)
		imagesAny := convertArrStrToArrAny(images)
		property.SetImages(imagesAny)

		jsonNote := getJsonStrPropertyNote(``, v.PhoneNumber, v.Description)

		property.SetNote(jsonNote)

		coord := getLatLong(v.Lat, v.Long)
		property.SetCoord(coord[:])

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		stat.Ok(property.DoUpsert())
	}
}

func ReadPropertyPT_RightmoveCoUk2(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyPT_RightmoveCoUk2: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()
	type propertyPT struct {
		SourceURL      string
		PropertyTitle  string
		Location       string
		PhoneNumber    string
		PrimaryPrice   string
		SecondaryPrice string
		PropertyType   string
		Bedrooms       string
		Bathrooms      string
		Size           string
		AgentName      string
		FullAddress    string
		Lat            string
		Long           string
		Description    string
		KeyFeatures    string
		Image          string
	}

	var properties []propertyPT

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceURL := tsv.String()
		propertyTitle := tsv.String()
		location := tsv.String()
		phoneNumber := tsv.String()
		primaryPrice := tsv.String()
		secondaryPrice := tsv.String()
		propertyType := tsv.String()
		bedrooms := tsv.String()
		bathrooms := tsv.String()
		size := tsv.String()
		agentName := tsv.String()
		fullAddress := tsv.String()
		lat := tsv.String()
		long := tsv.String()
		description := tsv.String()
		keyFeatures := tsv.String()
		images := tsv.String()

		properties = append(properties, propertyPT{
			SourceURL:      S.Trim(sourceURL),
			PropertyTitle:  S.Trim(propertyTitle),
			Location:       S.Trim(location),
			PhoneNumber:    S.Trim(phoneNumber),
			PrimaryPrice:   S.Trim(primaryPrice),
			SecondaryPrice: S.Trim(secondaryPrice),
			PropertyType:   S.Trim(propertyType),
			Bedrooms:       S.Trim(bedrooms),
			Bathrooms:      S.Trim(bathrooms),
			Size:           S.Trim(size),
			AgentName:      S.Trim(agentName),
			FullAddress:    S.Trim(fullAddress),
			Lat:            S.Trim(lat),
			Long:           S.Trim(long),
			Description:    S.Trim(description),
			KeyFeatures:    S.Trim(keyFeatures),
			Image:          S.Trim(images),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:] // remove 1st element
	} else {
		panic(`Properties from ` + resourcePath + ` is empty`)
	}

	stat := &model.ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	for _, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyMutator(conn)
		if isValidURL(v.SourceURL) {
			propKey, err := getSheetPropertyPTUniqPropKey(v.SourceURL)
			if err != nil {
				stat.Skip()
				continue
			}
			property.SetUniqPropKey(propKey + `_pt`)
			property.FindByUniqPropKey()
		} else {
			stat.Skip()
			continue
		}
		purpose := getPropertyPurpose(v.PropertyTitle)
		property.SetPurpose(purpose)
		property.SetCity(`Lisbon`)

		priceUSD := convertGBPToUSD(v.PrimaryPrice)
		property.SetLastPrice(priceUSD)
		property.SetHouseType(v.PropertyType)
		property.SetBedroom(S.ToI(v.Bedrooms))
		property.SetBathroom(S.ToI(v.Bathrooms))

		sizeSqFt := getSizeSqFt(v.Size)
		property.SetTotalSqft(sizeSqFt)
		property.SetAddress(v.FullAddress)

		images := S.Split(v.Image, `|`)
		imagesAny := convertArrStrToArrAny(images)
		property.SetImages(imagesAny)

		jsonNote := getJsonStrPropertyNote(``, v.PhoneNumber, v.Description)

		property.SetNote(jsonNote)

		coord := getLatLong(v.Lat, v.Long)
		property.SetCoord(coord[:])

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		stat.Ok(property.DoUpsert())
	}
}

func getSheetPropertyPTUniqPropKey(sourceURL string) (string, error) {
	parsedUrl, err := url.Parse(sourceURL)
	if err != nil {
		return ``, errors.New(`invalid URL`)
	}

	pathSegments := S.Split(parsedUrl.Path, `/`)
	if len(pathSegments) < 3 {
		return ``, errors.New(`cannot find unique property key`)
	}

	propKey := pathSegments[2]

	return propKey, nil
}

func getPropertyPurpose(str string) string {
	if idx := S.IndexOf(str, mProperty.PropertyPurposeSale); idx != -1 {
		return str[idx:]
	} else {
		if idx := S.IndexOf(str, mProperty.PropertyPurposeRent); idx != -1 {
			return str[idx:]
		} else {
			return mProperty.PropertyPurposeSale
		}
	}
}

func getJsonStrPropertyNote(email, phone, about string) string {
	pNote := rqProperty.PropertyNote{
		ContactEmail: email,
		ContactPhone: phone,
		About:        about,
	}
	jsonByte, err := json.Marshal(pNote)
	if err != nil {
		return `{"contactEmail":"` + email + `","contactPhone":"` + phone + `","about":"` + about + `"}`
	}

	return string(jsonByte)
}

func convertGBPToUSD(gbpStrRaw string) string {
	gbpStrRaw = strings.Replace(gbpStrRaw, ` pcm`, ``, 1)
	gbpStrRaw = strings.Replace(gbpStrRaw, `£`, ``, 1)
	gbpStrRaw = strings.Replace(gbpStrRaw, `,`, ``, 1)

	gbpFloat := S.ToF(gbpStrRaw)

	var exchangeRateUSD float64 = 1.25 // possibly change

	usdAmount := (gbpFloat * exchangeRateUSD)

	return fmt.Sprintf("%.2f", usdAmount)
}

func getSizeSqFt(strSize string) float64 {
	if S.Contains(strSize, `sq ft`) {
		sqFtArr := S.Split(strSize, ` `)
		sqFtStr := sqFtArr[0]
		sqFtStr = S.Replace(sqFtStr, ",", "")

		return S.ToF(sqFtStr)
	}

	return 0
}

func convertArrStrToArrAny(arrStr []string) []any {
	var arrAny []any
	for _, str := range arrStr {
		arrAny = append(arrAny, S.Trim(str))
	}

	return arrAny
}

func getLatLong(latRawStr, lngRawStr string) [2]any {
	coord := [2]any{0, 0}
	if !(latRawStr == `` && lngRawStr == ``) {
		latitude := S.Replace(latRawStr, `,`, `.`)
		coord[0] = S.ToF(latitude)

		longitude := S.Replace(lngRawStr, `,`, `.`)
		coord[1] = S.ToF(longitude)
	}

	return coord
}
