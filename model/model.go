package model

import (
	"fmt"
	"strconv"
	"street/model/mPropertyHistory/wcPropertyHistory"
	"time"

	"street/model/mAuth/wcAuth"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"
	"street/model/mPropertyHistory"

	"github.com/xuri/excelize/v2"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"street/model/mAuth"
)

type Migrator struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter
}

func RunMigration(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	propOltp *Tt.Adapter,
	propOlap *Ch.Adapter,
) {
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
		PropOltp: propOltp,
		PropOlap: propOlap,
	}
	mAuth.TarantoolTables[mAuth.TableUsers].PreUnique1MigrationHook = wcAuth.UniqueUsernameMigration
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.PropOltp.MigrateTables(mProperty.TarantoolTables)
	m.PropOlap.MigrateTables(mProperty.ClickhouseTables)
	m.PropOltp.MigrateTables(mPropertyHistory.TarantoolTables)
	m.PropOlap.MigrateTables(mPropertyHistory.ClickhouseTables)
}

func GetHouseAddressInBuySellData(adapter **Tt.Adapter, resourceFile string) {
	fmt.Println("[Scan Buy-Sell sheet] Begin the process update address to house")
	pathData := resourceFile
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}
	propertyMutator := wcProperty.NewPropertyMutator(*adapter)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("不動產買賣BuyandSell")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat := &ImporterStat{Total: len(rows)}
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
		existingHouseData := propertyMutator.FindPropertiesBySerialNumber(serialPropertyNumber)

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

	fmt.Println("\n[Scan Buy-Sell sheet] End the process update address to house")

}

func GetHouseAddressInRentData1(adapter **Tt.Adapter, resourceFile string) {
	fmt.Println("[Scan Rent sheet] Begin the process update address to house")
	pathData := resourceFile
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}
	propertyMutator := wcProperty.NewPropertyMutator(*adapter)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("不動產租賃Rent")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat := &ImporterStat{Total: len(rows)}
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
		existingHouseData := propertyMutator.FindPropertiesBySerialNumber(serialPropertyNumber)

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

	fmt.Println("\n[Scan Rent sheet] End the process update address to house")
}

func GetHouseAddressInRentData2(adapter **Tt.Adapter, resourceFile string) {
	fmt.Println("[Scan Rent sheet2] Begin the process update address to house")
	pathData := resourceFile
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}
	propertyMutator := wcProperty.NewPropertyMutator(*adapter)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("不動產租賃Rent")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat := &ImporterStat{Total: len(rows)}
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
		existingHouseData := propertyMutator.FindPropertiesBySerialNumber(serialPropertyNumber)

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

	fmt.Println("\n[Scan Rent sheet2] End the process update address to house")
}

func ReadHouseDataSheet(adapter *Tt.Adapter, resourcePath string) {
	pathData := resourcePath
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("建物HouseDetail")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat := &ImporterStat{Total: len(rows)}
	fmt.Println("Begin the process of import house data")
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
			currentTime := time.Now().UnixMilli()
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
	fmt.Println("\nEnd process of import house data")
}

func ImportHouseHistoryInBuySellSheet(adapter **Tt.Adapter, resourcePath string) {
	fmt.Println("[Start] Beginning of import house history data in buy-sell sheet")
	pathData := resourcePath
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}
	propertyHistoryMutator := wcPropertyHistory.NewPropertyHistoryMutator(*adapter)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("不動產買賣BuyandSell")
	stat := &ImporterStat{Total: len(rows)}
	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 1 {
			stat.Skip()
			continue
		}

		propertyHistoryMutator := wcPropertyHistory.NewPropertyHistoryMutator(*adapter)

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
					fmt.Println("Error during conversion for total-price [" + colCell + "]")
					propertyHistoryMutator.PriceNtd = 0
				} else {
					propertyHistoryMutator.PriceNtd = int64(totalPriceData)
				}

			} else if colIndex == 22 {
				pricePerUnitData, err := strconv.Atoi(colCell)
				if err != nil {
					fmt.Println("Error during conversion for price per unit data [" + colCell + "]")
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

		currentTime := time.Now().UnixMilli()
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
	fmt.Println(propertyHistoryMutator)
	fmt.Println("[End] End of import house history data  in buy-sell sheet")
}

func ImportHouseHistoryInRentSheet(adapter *Tt.Adapter, resourcePath string) {
	fmt.Println("[Start] Beginning of import house history data in rent sheet")
	pathData := resourcePath
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}
	propertyHistoryMutator := wcPropertyHistory.NewPropertyHistoryMutator(adapter)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("不動產租賃Rent")
	stat := &ImporterStat{Total: len(rows)}
	for index, row := range rows {
		stat.Print()

		if index >= 0 && index <= 2 {
			stat.Skip()
			continue
		}

		propertyHistoryMutator := wcPropertyHistory.NewPropertyHistoryMutator(adapter)

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
					fmt.Println("Error during conversion for total-price [" + colCell + "]")
					propertyHistoryMutator.PriceNtd = 0
				} else {
					propertyHistoryMutator.PriceNtd = int64(totalPriceData)
				}

			} else if colIndex == 23 {
				pricePerUnitData, err := strconv.Atoi(colCell)
				if err != nil {
					fmt.Println("Error during conversion for price per unit data [" + colCell + "]")
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

		currentTime := time.Now().UnixMilli()
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
	fmt.Println(propertyHistoryMutator)
	fmt.Println("[End] End of import house history data  in rent sheet")
}

func ImportExcelData(adapter *Tt.Adapter, resourcePath string) {
	fmt.Println("[Start] Beginning of process data")

	ReadHouseDataSheet(adapter, resourcePath)
	GetHouseAddressInBuySellData(&adapter, resourcePath)
	GetHouseAddressInRentData1(&adapter, resourcePath)
	GetHouseAddressInRentData2(&adapter, resourcePath)

	fmt.Println("[End] End process of import house data")

	fmt.Println("=========")

	fmt.Println("[Start] Import history of house data")
	ImportHouseHistoryInBuySellSheet(&adapter, resourcePath)
	ImportHouseHistoryInRentSheet(adapter, resourcePath)
	fmt.Println("[End] Import history of house data")
}
