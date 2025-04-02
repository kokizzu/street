package zImport

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"street/model"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/valyala/tsvreader"
)

func ReadPropertyJP(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyJP: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}
	defer file.Close()

	type propertyJp struct {
		sellingPrice   string
		buildingName   string
		city           string
		district       string
		address        string
		roomType       string
		sizeM2         string
		floor          string
		year           string
		month          string
		closestStation string
		minToStation   string
		status         string
	}

	var properties []propertyJp

	tsv := tsvreader.New(file)
	for tsv.Next() {
		property := propertyJp{}
		property.sellingPrice = tsv.String()
		property.buildingName = tsv.String()
		property.city = tsv.String()
		property.district = tsv.String()
		property.address = tsv.String()
		property.roomType = tsv.String()
		property.sizeM2 = tsv.String()
		property.floor = tsv.String()
		property.year = tsv.String()
		property.month = tsv.String()
		property.closestStation = tsv.String()
		property.minToStation = tsv.String()
		property.status = tsv.String()

		properties = append(properties, property)
	}

	if len(properties) > 0 {
		properties = properties[1:]
	} else {
		L.Print(`ReadPropertyJP: no data`)
		os.Exit(1)
	}

	stat := &model.ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	for _, v := range properties {
		stat.Print()

		prop := wcProperty.NewPropertyMutator(conn)

		uniqPropKey := randChar(15) + `_jp`
		prop.SetUniqPropKey(uniqPropKey)
		prop.SetCity(v.city)
		prop.SetDistrict(v.district)
		prop.SetAddress(v.address)
		prop.SetSizeM2(v.sizeM2)

		floorSlice := S.Split(v.floor, `/`)

		if len(floorSlice) == 2 {
			prop.SetNumberOfFloors(floorSlice[1])
		} else {
			prop.SetNumberOfFloors(floorSlice[0])
		}

		prop.SetYearBuilt(S.ToI(v.year))

		attrByt, err := json.Marshal(mProperty.PropertyAttribute{
			ClosestStation:   v.closestStation,
			MinutesToStation: S.ToI(v.minToStation),
		})
		if err != nil {
			stat.Skip()
		}
		prop.SetAttribute(string(attrByt))

		price := convertJPYToUSD(v.sellingPrice)
		prop.SetLastPrice(price)
		prop.SetCoord([]any{0, 0})
		prop.SetCountryCode(`JP`)

		stat.Ok(prop.DoUpsert())
	}
}

func convertJPYToUSD(jpyStr string) string {
	jpyStr = S.Trim(jpyStr)

	var exchangeRate float64 = 0.0067

	usdAmount := (S.ToF(jpyStr) * exchangeRate)

	return fmt.Sprintf("%.2f", usdAmount)
}

func randChar(length int) string {
	const letterBytes = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
