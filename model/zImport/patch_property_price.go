package zImport

import (
	"fmt"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"strconv"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"time"
)

func PatchPropertiesPrice(propOltp *Tt.Adapter) {
	// data was imported with unix millis, it should be unix
	start := time.Now()

	UpdatePriceToProperties(propOltp)

	L.TimeTrack(start, `Patch to properties price`)
}

func UpdatePriceToProperties(propOltp *Tt.Adapter) {

	propertyMutator := wcProperty.NewPropertyMutator(propOltp)
	propertyHistoryRq := rqProperty.NewPropertyHistory(propOltp)

	properties := propertyMutator.FindAllProperties()

	stat := &ImporterStat{Total: len(properties), PrintEvery: 10}

	for _, p := range properties {
		stat.Print()

		if p.LastPrice != "" && len(p.PriceHistories) > 0 {
			stat.Skip()
			continue
		}

		p.LastPrice = "0"
		p.PriceHistories = []any{}
		// Update price for property based on history sheet
		propertyHistoryList := propertyHistoryRq.FindByPropertyKey(p.UniqPropKey)

		if len(propertyHistoryList) > 0 {

			for _, ph := range propertyHistoryList {
				if ph.PriceNtd == 0 {
					stat.Skip()
					continue
				}
				fmt.Println("Property key => ", ph.PropertyKey)
				fmt.Println("Price NTD => ", ph.PriceNtd)

				p.PriceHistories = append(p.PriceHistories, ph.PriceNtd)
				currentPrice, _ := strconv.ParseInt(p.LastPrice, 10, 64)
				if p.LastPrice == "" {
					p.LastPrice = strconv.Itoa(int(ph.PriceNtd))
				} else if p.LastPrice != "" && ph.PriceNtd > currentPrice {
					p.LastPrice = strconv.Itoa(int(ph.PriceNtd))
				}
			}
			fmt.Println("Price History => ", p.PriceHistories)
			propertyMutator.Property = *p
			propertyMutator.Adapter = propOltp
			propertyMutator.UpdatedAt = time.Now().Unix()
			stat.Ok(propertyMutator.DoOverwriteById())

		} else {
			// If unable to find the property through propKey serialNumber#House => Get prices by serial number
			pHistoryList := propertyHistoryRq.FindBySerialNumber(p.SerialNumber)

			for _, ph := range pHistoryList {
				if ph.PriceNtd == 0 {
					stat.Skip()
					continue
				}
				p.PriceHistories = append(p.PriceHistories, ph.PriceNtd)

				currentPrice, _ := strconv.ParseInt(p.LastPrice, 10, 64)
				if p.LastPrice == "" {
					p.LastPrice = strconv.Itoa(int(ph.PriceNtd))
				} else if p.LastPrice != "" && ph.PriceNtd > currentPrice {
					p.LastPrice = strconv.Itoa(int(ph.PriceNtd))
				}
			}
			//fmt.Println("LAst price => ", p.LastPrice)
			fmt.Println("Price History => ", p.PriceHistories)
			propertyMutator.Property = *p
			propertyMutator.Adapter = propOltp
			propertyMutator.UpdatedAt = time.Now().Unix()
			stat.Ok(propertyMutator.DoOverwriteById())
		}
	}

}
