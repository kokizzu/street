package zImport

import (
	"strconv"
	"strings"
	"time"

	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

func ImportPropertyHistories(propertyResponseObject *PropertyFullResponse, redFinId string) []rqProperty.PropertyHistoryUS {

	if propertyResponseObject.PropertyHistoryInfo.HasPropertyHistory == false {
		return []rqProperty.PropertyHistoryUS{}
	}

	propHistories := []rqProperty.PropertyHistoryUS{}

	for _, historyItem := range propertyResponseObject.PropertyHistoryInfo.Events {

		// Tracking only history have price
		if historyItem.Price == 0 && strings.ContainsAny(historyItem.EventDescription, "Pending") {
			continue
		}

		data := rqProperty.PropertyHistoryUS{}

		data.TransactionTime = strconv.FormatInt(historyItem.EventDate, 10)
		data.TransactionKey = X.ToS(historyItem.EventDate)
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

func SavePropertyHistories(adapter *Tt.Adapter, propList []rqProperty.PropertyHistoryUS, stat *ImporterStat) {

	for _, pHistory := range propList {

		propertyHistoryMutator := wcProperty.NewPropertyHistoryUSMutator(adapter)
		propertyHistoryMutator.PropertyHistoryUS = pHistory
		propertyHistoryMutator.Adapter = adapter

		propertyHistoryMutator.CreatedAt = time.Now().Unix()
		propertyHistoryMutator.UpdatedAt = time.Now().Unix()

		if propertyHistoryMutator.FindByTransactionKey() {
			continue
		}
		stat.Ok(propertyHistoryMutator.DoInsert())
	}
}

func ImportPropertyHistoryUsData(adapter *Tt.Adapter, baseUrl string, minPropertyId int, maxPropertyId int) {
	// const baseUrl = "https://www.redfin.com/stingray/api/home/details/belowTheFold"
	// const minPropertyId = 1
	// const maxPropertyId = 10000000

	stat := &ImporterStat{Total: maxPropertyId - minPropertyId, PrintEvery: 11}
	defer stat.Print(`last`)

	for i := minPropertyId; i <= maxPropertyId; i++ {
		stat.Print()

		redfinKey := `rf` + strconv.Itoa(i)

		data, err := fetchPropertyUSByPropID(baseUrl, i)
		L.PanicIf(err, "Error: fetchPropertyUSByPropID "+redfinKey)

		propertyResponse, ok := data["payload"].(map[string]any)
		if !ok {
			stat.Skip()
			continue
		}
		// Property into the struct
		propertyResponseObject, err := parseMapIntoStruct(propertyResponse)
		L.PanicIf(err, "Parse property response error")

		propertyHistories := ImportPropertyHistories(propertyResponseObject, redfinKey)

		stat.Print()

		// Store property history
		SavePropertyHistories(adapter, propertyHistories, stat)

	}
}
