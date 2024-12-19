package zImport

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"
	"strings"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

func ReadPropertyPT_Buy_ZomePT(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyPT_Buy_ZomePT: import property data`)()

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
		SourceURL    string
		PropertyName string
		Bedrooms     string
		Bathrooms    string
		Rooms        string
		ClosedGarage string
		PID          string
		BuyPrice     string
		TypePurpose  string
		Address      string
		Latitude     string
		Longitude    string
		Areas        string
		Attributes   string
		Areas2       string
		AgentName    string
		Title        string
		AgentNumber  string
		AgentEmail   string
		Description  string
		Image        string
	}

	var properties []propertyPT

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceURL := tsv.String()
		propName := tsv.String()
		bedrooms := tsv.String()
		bathrooms := tsv.String()
		rooms := tsv.String()
		closedGarage := tsv.String()
		pid := tsv.String()
		buyPrice := tsv.String()
		typePurpose := tsv.String()
		address := tsv.String()
		latitude := tsv.String()
		longitude := tsv.String()
		areas := tsv.String()
		attributes := tsv.String()
		areas2 := tsv.String()
		agentName := tsv.String()
		title := tsv.String()
		agentNumber := tsv.String()
		agentEmail := tsv.String()
		description := tsv.String()
		image := tsv.String()

		properties = append(properties, propertyPT{
			SourceURL:    S.Trim(sourceURL),
			PropertyName: S.Trim(propName),
			Bedrooms:     S.Trim(bedrooms),
			Bathrooms:    S.Trim(bathrooms),
			Rooms:        S.Trim(rooms),
			ClosedGarage: S.Trim(closedGarage),
			PID:          S.Trim(pid),
			BuyPrice:     S.Trim(buyPrice),
			TypePurpose:  S.Trim(typePurpose),
			Address:      S.Trim(address),
			Latitude:     S.Trim(latitude),
			Longitude:    S.Trim(longitude),
			Areas:        S.Trim(areas),
			Attributes:   S.Trim(attributes),
			Areas2:       S.Trim(areas2),
			AgentName:    S.Trim(agentName),
			Title:        S.Trim(title),
			AgentNumber:  S.Trim(agentNumber),
			AgentEmail:   S.Trim(agentEmail),
			Description:  S.Trim(description),
			Image:        S.Trim(image),
		})
	}

	if len(properties) > 1 {
		properties = properties[1:] // remove 1st element
	} else {
		panic(`Properties from ` + resourcePath + ` is empty`)
	}

	stat := &ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	rgxUsableAreaM2 := regexp.MustCompile(`(\d+m²) usable area`)

	for _, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyMutator(conn)
		property.SetBedroom(S.ToI(v.Bedrooms))
		property.SetBathroom(S.ToI(v.Bathrooms))
		property.SetUniqPropKey(v.PID + `_pt`)

		priceUSD := convertEURToUSD(v.BuyPrice)
		property.SetLastPrice(priceUSD)
		property.SetPurpose(mProperty.PropertyPurposeSale)
		property.SetAddress(v.Address)

		coord := [2]any{0, 0}
		if v.Latitude != `` {
			latitude := S.Replace(v.Latitude, `,`, `.`)
			coord[0] = S.ToF(latitude)
		}
		if v.Longitude != `` {
			longitude := S.Replace(v.Latitude, `,`, `.`)
			coord[1] = S.ToF(longitude)
		}
		property.SetCoord(coord[:])
		if v.Areas != `` {
			match := rgxUsableAreaM2.FindStringSubmatch(v.Areas)
			if len(match) > 1 {
				usableAreaM2 := S.Replace(match[1], `m²`, ``)
				property.SetSizeM2(usableAreaM2)
			}
		}

		jsonNote := getJsonStrPropertyNote(v.AgentEmail, v.AgentNumber, v.Description)
		property.SetNote(jsonNote)

		property.SetImages([]any{v.Image})

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())

		stat.Ok(property.DoUpsert())
	}
}

func getFloatPrice(str string) float64 {
	dotCount := S.Count(str, `.`)

	if dotCount > 1 {
		str = strings.ReplaceAll(str, `.`, ``)
		return S.ToF(str)
	}

	return S.ToF(str)
}

func convertEURToUSD(eurStrRaw string) string {
	eurStrRaw = S.Trim(eurStrRaw)
	if !S.Contains(eurStrRaw, `€`) {
		return "0.00"
	}

	eurStr := S.Trim(strings.ReplaceAll(eurStrRaw, `€`, ``))
	eurFloat := getFloatPrice(eurStr)

	var exchangeRateUSD float64 = 1.04 // possibly change

	usdAmount := (eurFloat * exchangeRateUSD)

	return fmt.Sprintf("%.2f", usdAmount)
}
