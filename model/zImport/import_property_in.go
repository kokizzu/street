package zImport

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"street/model"
	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/valyala/tsvreader"
)

func randChar(length int) string {
	const letterBytes = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func ReadPropertyIN(conn *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyIN: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(`failed to open file "`+path+`" : `, err)
		os.Exit(1)
	}
	defer file.Close()

	type propertyIn struct {
		PostName  string
		Type      string
		Address   string
		City      string
		State     string
		ZipCode   string
		Lat       string
		Lng       string
		MainPrice string
		Bed       string
		Bath      string
		SqFt      string
	}

	var properties []propertyIn

	tsv := tsvreader.New(file)
	for tsv.Next() {
		var prop propertyIn

		prop.PostName = S.Trim(tsv.String())
		prop.Type = S.Trim(tsv.String())
		prop.Address = S.Trim(tsv.String())
		prop.City = S.Trim(tsv.String())
		prop.State = S.Trim(tsv.String())
		prop.ZipCode = S.Trim(tsv.String())
		prop.Lat = S.Trim(tsv.String())
		prop.Lng = S.Trim(tsv.String())
		prop.MainPrice = S.Trim(tsv.String())
		prop.Bed = S.Trim(tsv.String())
		prop.Bath = S.Trim(tsv.String())
		prop.SqFt = S.Trim(tsv.String())

		properties = append(properties, prop)
	}

	properties = properties[1:]

	if len(properties) == 0 {
		fmt.Println(`no data found`)
		os.Exit(1)
	}

	stat := &model.ImporterStat{
		Total: len(properties),
	}
	defer stat.Print(`last`)

	for _, v := range properties {
		stat.Print()

		prop := wcProperty.NewPropertyMutator(conn)

		uniqPropKey := randChar(10) + `_in`

		prop.UniqPropKey = uniqPropKey
		if !prop.FindByUniqPropKey() {
			prop.SetUniqPropKey(uniqPropKey)
		}

		prop.SetHouseType(v.Type)
		prop.SetAddress(v.Address)
		prop.SetCountryCode(`IN`)
		prop.SetCity(v.City)
		prop.SetState(v.State)
		prop.SetZip(v.ZipCode)
		prop.SetCoord([]any{
			X.ToF(v.Lat),
			X.ToF(v.Lng),
		})
		priceArr := S.Split(v.MainPrice, ` `)
		if len(priceArr) > 0 {
			prop.SetLastPrice(convertINRToUSD(priceArr[0]))
		}
		prop.SetBedroom(S.ToI(v.Bed))
		prop.SetBathroom(S.ToI(v.Bath))
		prop.SetTotalSqft(S.ToF(v.SqFt))

		stat.Ok(prop.DoUpsert())
	}
}

func convertINRToUSD(inrStrt string) string {
	inrStrt = S.Trim(inrStrt)
	inrStrt = S.Replace(inrStrt, `₹`, ``)

	inrFloat := S.ToF(inrStrt)

	var exchangeRate float64 = 0.011

	usdAmount := (inrFloat * exchangeRate)

	return fmt.Sprintf("%.2f", usdAmount)
}
