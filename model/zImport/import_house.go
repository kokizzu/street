package zImport

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/xuri/excelize/v2"

	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
)

func readExcelWorksheet(pathData string, workspace string) (res [][]string, closerFunc func()) {
	f, err := excelize.OpenFile(pathData)
	L.PanicIf(err, `excelize.OpenFile`, pathData)

	closerFunc = func() {
		if err := f.Close(); err != nil {
			log.Println(`f.Close`, err)
		}
	}

	rows, err := f.GetRows(workspace)
	L.PanicIf(err, `f.GetRows`, workspace)

	return rows, closerFunc
}

func GetHouseAddressInBuySellData(adapter **Tt.Adapter, resourceFile string) {
	defer subTaskPrint(`GetHouseAddressInBuySellData: Buy-Sell sheet, update address to house`)()

	rows, closer := readExcelWorksheet(resourceFile, "不動產買賣BuyandSell")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {

		stat.Print()

		if index == 0 || index == 1 {
			stat.Skip()
			continue
		}

		var district = ""
		var address = ""
		var serialPropertyNumber = ""

		for colIndex, colCell := range row {
			// Col 0 - District
			// Col 2 - Address
			// Col 27 - House Serial Number

			if colIndex == 0 {
				district = colCell
			}
			if colIndex == 2 {
				address = colCell
			}
			if colIndex == 27 {
				serialPropertyNumber = colCell
			}
		}

		if district == "" || address == "" || serialPropertyNumber == "" {
			stat.Skip()
			continue
		}

		// Get list of property based on house serial number
		prop := rqProperty.NewProperty(*adapter)
		existingHouseData := prop.FindPropertiesBySerialNumber(serialPropertyNumber)

		stat.Total += len(existingHouseData) - 1
		for _, house := range existingHouseData {
			//fmt.Println("House data -> ", house)

			if house.Address != "" && house.District != "" {
				stat.Skip()
				continue
			}
			house.Address = address
			house.District = district

			// Updated house
			dataMutator := wcProperty.NewPropertyMutator(*adapter)
			dataMutator.Property = *house
			dataMutator.Adapter = *adapter
			dataMutator.UpdatedAt = time.Now().Unix()

			// Update
			stat.Ok(dataMutator.DoOverwriteById())
		}
	}

}

func GetHouseAddressInRentData1(adapter **Tt.Adapter, resourceFile string) {
	defer subTaskPrint(`GetHouseAddressInRentData1: Rent sheet, update address to house`)()

	rows, closer := readExcelWorksheet(resourceFile, "不動產租賃Rent")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 2 {
			stat.Skip()
			continue
		}

		var district = ""
		var address = ""
		var serialPropertyNumber = ""

		for colIndex, colCell := range row {
			// Col 0 - District
			// Col 2 - Address
			// Col 27/28 - House Serial Number

			if colIndex == 0 {
				district = colCell
			}
			if colIndex == 2 {
				address = colCell
			}
			if colIndex == 27 {
				serialPropertyNumber = colCell
			}
		}

		//fmt.Println("Row index => ", index)
		if district == "" || address == "" || serialPropertyNumber == "" {
			stat.Skip()
			continue
		}

		// Get list of property based on house serial number
		prop := rqProperty.NewProperty(*adapter)
		existingHouseData := prop.FindPropertiesBySerialNumber(serialPropertyNumber)

		stat.Total += len(existingHouseData) - 1
		for _, house := range existingHouseData {
			//fmt.Println("House data -> ", house)

			if house.Address != "" && house.District != "" {
				stat.Skip()
				continue
			}
			house.Address = address
			house.District = district

			// Updated house
			dataMutator := wcProperty.NewPropertyMutator(*adapter)
			dataMutator.Property = *house
			dataMutator.Adapter = *adapter
			dataMutator.UpdatedAt = time.Now().Unix()

			// Update
			stat.Ok(dataMutator.DoOverwriteById())
		}

	}
}

func GetHouseAddressInRentData2(adapter **Tt.Adapter, resourceFile string) {
	defer subTaskPrint(`GetHouseAddressInRentData2: Rent sheet2, update address to house`)()

	rows, closer := readExcelWorksheet(resourceFile, "不動產租賃Rent")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 2 {
			stat.Skip()
			continue
		}

		var district = ""
		var address = ""
		var serialPropertyNumber = ""

		for colIndex, colCell := range row {
			// Col 0 - District
			// Col 2 - Address
			// Col 27/28 - House Serial Number

			if colIndex == 0 {
				district = colCell
			}
			if colIndex == 2 {
				address = colCell
			}
			if colIndex == 28 {
				serialPropertyNumber = colCell
			}
		}

		//fmt.Println("Row index => ", index)
		if district == "" || address == "" || serialPropertyNumber == "" {
			stat.Skip()
			continue
		}

		// Get list of property based on house serial number
		prop := rqProperty.NewProperty(*adapter)
		existingHouseData := prop.FindPropertiesBySerialNumber(serialPropertyNumber)

		stat.Total += len(existingHouseData) - 1
		for _, house := range existingHouseData {
			//fmt.Println("House data -> ", house)

			if house.Address != "" && house.District != "" {
				stat.Skip()
				continue
			}
			house.Address = address
			house.District = district

			// Updated house
			dataMutator := wcProperty.NewPropertyMutator(*adapter)
			dataMutator.Property = *house
			dataMutator.Adapter = *adapter
			dataMutator.UpdatedAt = time.Now().Unix()

			// Update
			stat.Ok(dataMutator.DoOverwriteById())
		}

	}
}

func ReadHouseDataSheet(adapter *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadHouseDataSheet: import house data`)()

	rows, closer := readExcelWorksheet(resourcePath, "建物HouseDetail")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {
		stat.Print()

		if index == 0 || index == 1 {
			stat.Skip()
			continue
		}

		propertyMutator := wcProperty.NewPropertyMutator(adapter)
		propertyMutator.Coord = []any{0, 0}
		for colIndex, colCell := range row {
			propertyMutator.Address = ""
			if colIndex == 0 {
				propertyMutator.SerialNumber = colCell
			} else if colIndex == 2 {
				propertyMutator.SizeM2 = colCell
			} else if colIndex == 3 {
				propertyMutator.MainUse = colCell
			} else if colIndex == 4 {
				propertyMutator.MainBuildingMaterial = colCell
			} else if colIndex == 5 {
				propertyMutator.ConstructCompletedDate = colCell
			} else if colIndex == 6 {
				propertyMutator.NumberOfFloors = colCell
			} else if colIndex == 7 {
				propertyMutator.BuildingLamination = colCell
			}
			// Update time
			currentTime := time.Now().Unix()
			propertyMutator.CreatedAt = currentTime
			propertyMutator.UpdatedAt = currentTime
			//fmt.Println(propertyMutator)
		}

		// Build unique key with serial number and size
		uniqueSerialNumber := propertyMutator.SerialNumber +
			"#" + propertyMutator.SizeM2
		if uniqueSerialNumber == "#" {
			stat.Skip()
			continue
		}
		propertyMutator.UniqPropKey = uniqueSerialNumber

		// Check if unique property key is existed
		if propertyMutator.FindByUniqPropKey() {
			stat.Skip()
			continue
		}

		stat.Ok(propertyMutator.DoInsert())
	}
}
func ImportHouseHistoryInBuySellSheet(adapter **Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ImportHouseHistoryInBuySellSheet: import house history data in buy-sell sheet`)()

	rows, closer := readExcelWorksheet(resourcePath, "不動產買賣BuyandSell")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 1 {
			stat.Skip()
			continue
		}

		propertyHistoryMutator := wcProperty.NewPropertyHistoryMutator(*adapter)

		var propertySize = ""
		var serialNumber = ""
		var uniquePropertyKey = ""
		var transactionKey = "" // Combination of serialNumber#size#transactionTime

		for colIndex, colCell := range row {

			// col 0 : District
			// col 1 : TransactionSign
			// col 2 : Address
			// col 3 : Size of property
			// col 7 : Transaction time
			// col 8 : Transaction pen number

			// col 21 : total price NTD
			// col 22 : Unit price per square
			// col 26 : transaction note
			// col 27 : Serial number

			if colIndex == 0 {
				propertyHistoryMutator.District = colCell
			} else if colIndex == 1 {
				propertyHistoryMutator.TransactionSign = colCell
			} else if colIndex == 2 {
				propertyHistoryMutator.Address = colCell
			} else if colIndex == 3 {
				propertySize = colCell
			} else if colIndex == 7 {
				propertyHistoryMutator.TransactionTime = colCell
			} else if colIndex == 8 {
				propertyHistoryMutator.TransactionNumber = colCell
			} else if colIndex == 21 {
				totalPriceData, err := strconv.Atoi(colCell)
				if err != nil {
					//fmt.Println("Error during conversion for total-price [" + colCell + "]")
					stat.Warn(`invalid PriceNtd`)
					propertyHistoryMutator.PriceNtd = 0
				} else {
					propertyHistoryMutator.PriceNtd = int64(totalPriceData)
				}

			} else if colIndex == 22 {
				pricePerUnitData, err := strconv.Atoi(colCell)
				if err != nil {
					//fmt.Println("Error during conversion for price per unit data [" + colCell + "]")
					stat.Warn(`invalid PricePerUnit`)
					propertyHistoryMutator.PricePerUnit = 0
				} else {
					propertyHistoryMutator.PricePerUnit = int64(pricePerUnitData)
				}
			} else if colIndex == 26 {
				propertyHistoryMutator.Note = colCell
			} else if colIndex == 27 {
				serialNumber = colCell
			}
		}
		uniquePropertyKey = serialNumber + "#" + propertySize
		transactionKey = serialNumber + "#" + propertySize + "#" + propertyHistoryMutator.TransactionTime

		currentTime := time.Now().Unix()
		propertyHistoryMutator.CreatedAt = currentTime
		propertyHistoryMutator.UpdatedAt = currentTime
		propertyHistoryMutator.PropertyKey = uniquePropertyKey
		propertyHistoryMutator.TransactionKey = transactionKey
		propertyHistoryMutator.TransactionType = "BUY_SELL"

		if propertyHistoryMutator.FindByTransactionKey() {
			stat.Skip()
			continue
		}
		stat.Ok(propertyHistoryMutator.DoInsert())

	}
}

func ImportHouseHistoryInRentSheet(adapter *Tt.Adapter, resourcePath string) {
	defer subTaskPrint("ImportHouseHistoryInRentSheet: import house history data in rent sheet")()

	rows, closer := readExcelWorksheet(resourcePath, "不動產租賃Rent")
	defer closer()

	stat := &ImporterStat{Total: len(rows)}
	defer stat.Print()

	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 2 {
			stat.Skip()
			continue
		}

		propertyHistoryMutator := wcProperty.NewPropertyHistoryMutator(adapter)

		var propertySize = ""
		var serialNumber = ""
		var uniquePropertyKey = ""
		var transactionKey = "" // Combination of serialNumber#size#transactionTime

		for colIndex, colCell := range row {

			// col 0 : District
			// col 1 : TransactionSign
			// col 2 : Address
			// col 3 : Size of property
			// col 7 : Transaction time
			// col 8 : Transaction pen number

			// col 22 : total price NTD
			// col 23 : Unit price per square
			// col 27 : transaction note
			// col 28 : Serial number

			if colIndex == 0 {
				propertyHistoryMutator.District = colCell
			} else if colIndex == 1 {
				propertyHistoryMutator.TransactionSign = colCell
			} else if colIndex == 2 {
				propertyHistoryMutator.Address = colCell
			} else if colIndex == 3 {
				propertySize = colCell
			} else if colIndex == 7 {
				propertyHistoryMutator.TransactionTime = colCell
			} else if colIndex == 8 {
				propertyHistoryMutator.TransactionNumber = colCell
			} else if colIndex == 22 {
				totalPriceData, err := strconv.Atoi(colCell)
				if err != nil {
					//fmt.Println("Error during conversion for total-price [" + colCell + "]")
					stat.Warn(`invalid PriceNtd`)
					propertyHistoryMutator.PriceNtd = 0
				} else {
					propertyHistoryMutator.PriceNtd = int64(totalPriceData)
				}

			} else if colIndex == 23 {
				pricePerUnitData, err := strconv.Atoi(colCell)
				if err != nil {
					stat.Warn(`invalid PricePerUnit`)
					//fmt.Println("Error during conversion for price per unit data [" + colCell + "]")
					propertyHistoryMutator.PricePerUnit = 0
				} else {
					propertyHistoryMutator.PricePerUnit = int64(pricePerUnitData)
				}
			} else if colIndex == 27 {
				propertyHistoryMutator.Note = colCell
			} else if colIndex == 28 {
				serialNumber = colCell
			}
		}
		uniquePropertyKey = serialNumber + "#" + propertySize
		transactionKey = serialNumber + "#" + propertySize + "#" + propertyHistoryMutator.TransactionTime

		currentTime := time.Now().Unix()
		propertyHistoryMutator.CreatedAt = currentTime
		propertyHistoryMutator.UpdatedAt = currentTime
		propertyHistoryMutator.PropertyKey = uniquePropertyKey
		propertyHistoryMutator.TransactionKey = transactionKey
		propertyHistoryMutator.TransactionType = "RENT"

		if propertyHistoryMutator.FindByTransactionKey() {
			stat.Skip()
			continue
		}
		stat.Ok(propertyHistoryMutator.DoInsert())

	}
}

func ImportExcelData(adapter *Tt.Adapter, resourcePath string) {
	start := time.Now()

	fmt.Println("[Start] House")
	ReadHouseDataSheet(adapter, resourcePath)
	GetHouseAddressInBuySellData(&adapter, resourcePath)
	GetHouseAddressInRentData1(&adapter, resourcePath)
	GetHouseAddressInRentData2(&adapter, resourcePath)
	fmt.Println("[End] House")

	fmt.Println("[Start] House Trx History")
	ImportHouseHistoryInBuySellSheet(&adapter, resourcePath)
	ImportHouseHistoryInRentSheet(adapter, resourcePath)
	fmt.Println("[End] House Trx History")

	L.TimeTrack(start, "ImportExcelData")
}

func subTaskPrint(str string) func() {
	fmt.Println(`  [Start] ` + str)
	return func() {
		fmt.Println("\n" + `  [End] ` + str)
	}
}
