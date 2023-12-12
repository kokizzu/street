package zImport

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"street/model/mProperty/wcProperty"
	"street/model/xGmap"
	"time"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
)

type PropertyTWFullResponse struct {
	Id                  int         `json:"id"`
	Address             string      `json:"address"`
	BuildAreaV2         AreaInfo    `json:"build_area_v2"`
	BuildPurposeStr     string      `json:"build_purpose_str"`
	HistoryTransLog     []TransInfo `json:"history_trans_log"`
	Layout              string      `json:"layout"`
	Month               string      `json:"month"`
	RealParkAreaV2      AreaInfo    `json:"real_park_area_v2"`
	RealParkTotalPriceV PriceInfo   `json:"real_park_total_price_v"`
	ShiftFloor          string      `json:"shift_floor"`
	TotalFloor          string      `json:"total_floor"`
	TotalPrice          string      `json:"total_price"`
	UnitPrice           PriceInfo   `json:"unit_price"`
}

type AreaInfo struct {
	Area string
	Unit string
}

type PriceInfo struct {
	Price string
	Unit  string
}

type TransInfo struct {
	Info struct {
		Rate string
		Type int
	} `json:"info"`
	Number     int       `json:"number"`
	Price      PriceInfo `json:"price"`
	TotalPrice PriceInfo `json:"total_price"`
	TransDate  string    `json:"trans_date"`
}

const PROPERTYTW_FILE = `static/property_tw_data/props_tw.jsonl`

func retrievePropertyTwLatLong(propertyMutator *wcProperty.PropertyTWMutator, gmap xGmap.Gmap) error {
	propertyMutator.Coord = []any{0, 0}

	fullUrl := gmap.BuildFullLocationSearchUrl(propertyMutator.Address)
	locationResponse, err := http.Get(fullUrl)

	if L.IsError(err, `retrievePropertyTwLatLong: get location response`) {
		return err
	}

	responseData, err := io.ReadAll(locationResponse.Body)
	if L.IsError(err, `retrievePropertyTwLatLong: read response body`) {
		return err
	}

	propertyLocation := PlaceResponse{}

	err = json.Unmarshal(responseData, &propertyLocation)
	if L.IsError(err, `retrievePropertyTwLatLong: unmarshal response data`) {
		L.Print(propertyMutator.Address)
		L.Print(string(responseData))
		return err
	}
	if len(propertyLocation.Candidates) == 0 {
		return errors.New(`property location zero candidates`)
	}

	propertyMutator.FormattedAddress = propertyLocation.Candidates[0].FormattedAddress

	latitude := propertyLocation.Candidates[0].Geometry.Location.Lat
	longitude := propertyLocation.Candidates[0].Geometry.Location.Lng

	propertyMutator.Coord = []any{latitude, longitude}
	return nil
}

func parsePropertyTwData(propertyMutator *wcProperty.PropertyTWMutator, propertyResponseObject *PropertyTWFullResponse, stat *ImporterStat, gmap xGmap.Gmap) {
	propertyMutator.Address = propertyResponseObject.Address

	fmt.Sscanf(propertyResponseObject.Layout,
		"%d房%d廳%d衛",
		&propertyMutator.Bedroom,
		&propertyMutator.Livingroom,
		&propertyMutator.Bathroom)
	propertyMutator.NumberOfFloors = propertyResponseObject.TotalFloor
	propertyMutator.HouseType = propertyResponseObject.BuildPurposeStr

	propertyMutator.SizeM2 = propertyResponseObject.BuildAreaV2.Area
	fmt.Sscanf(propertyResponseObject.RealParkAreaV2.Area, "%f", &propertyMutator.Parking)

	propertyMutator.LastPrice = propertyResponseObject.TotalPrice
	for _, trans := range propertyResponseObject.HistoryTransLog {
		propertyMutator.PriceHistoriesSell = append(propertyMutator.PriceHistoriesSell, any(trans))
	}

	err := retrievePropertyTwLatLong(propertyMutator, gmap)
	if err != nil {
		stat.Warn(`no lat long`)
	}

	// -------- Metadata ---------
	currentTime := time.Now().Unix()
	propertyMutator.CreatedAt = currentTime
	propertyMutator.UpdatedAt = currentTime
}

func ImportPropertyTwData(adapter *Tt.Adapter, gmap xGmap.Gmap) {
	f, err := os.Open(PROPERTYTW_FILE)
	if L.IsError(err, `failed opening file %s`, PROPERTYTW_FILE) {
		L.Print(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var props []PropertyTWFullResponse
	for scanner.Scan() {
		var prop PropertyTWFullResponse
		line := scanner.Text()
		if line == "\"\"" {
			continue
		}
		err := json.Unmarshal([]byte(line), &prop)
		if !L.IsError(err, `failed reading PropertyTW`) {
			props = append(props, prop)
		}
	}

	stat := &ImporterStat{Total: len(props), PrintEvery: 11}
	defer stat.Print(`last`)

	for _, prop := range props {
		stat.Print()

		twPropKey := `tw` + strconv.Itoa(prop.Id)

		propertyMutator := wcProperty.NewPropertyTWMutator(adapter)
		propertyMutator.UniqPropKey = twPropKey
		if propertyMutator.FindByUniqPropKey() {
			stat.Skip()
			continue
		}

		parsePropertyTwData(propertyMutator, &prop, stat, gmap)
		stat.Ok(propertyMutator.DoInsert())
	}
}
