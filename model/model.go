package model

import (
	"fmt"
	"time"

	"street/model/mAuth/wcAuth"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"

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

		stat.Total += len(existingHouseData)
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
			dataMutator.Property = rqProperty.Property{
				Adapter:                *adapter,
				Id:                     house.Id,
				UniqPropKey:            house.UniqPropKey,
				SerialNumber:           house.SerialNumber,
				SizeM2:                 house.SizeM2,
				MainUse:                house.MainUse,
				MainBuildingMaterial:   house.MainBuildingMaterial,
				ConstructCompletedDate: house.ConstructCompletedDate,
				NumberOfFloors:         house.NumberOfFloors,
				BuildingLamination:     house.BuildingLamination,
				Address:                house.Address,
				District:               house.District,
				Note:                   house.Note,
				CreatedAt:              house.CreatedAt,
				CreatedBy:              house.CreatedBy,
				UpdatedAt:              time.Now().UnixMilli(),
				UpdatedBy:              house.UpdatedBy,
				DeletedAt:              house.DeletedAt,
			}

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

		stat.Total += len(existingHouseData)
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
			dataMutator.Property = rqProperty.Property{
				Adapter:                *adapter,
				Id:                     house.Id,
				UniqPropKey:            house.UniqPropKey,
				SerialNumber:           house.SerialNumber,
				SizeM2:                 house.SizeM2,
				MainUse:                house.MainUse,
				MainBuildingMaterial:   house.MainBuildingMaterial,
				ConstructCompletedDate: house.ConstructCompletedDate,
				NumberOfFloors:         house.NumberOfFloors,
				BuildingLamination:     house.BuildingLamination,
				Address:                house.Address,
				District:               house.District,
				Note:                   house.Note,
				CreatedAt:              house.CreatedAt,
				CreatedBy:              house.CreatedBy,
				UpdatedAt:              time.Now().UnixMilli(),
				UpdatedBy:              house.UpdatedBy,
				DeletedAt:              house.DeletedAt,
			}

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

		stat.Total += len(existingHouseData)
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
			dataMutator.Property = rqProperty.Property{
				Adapter:                *adapter,
				Id:                     house.Id,
				UniqPropKey:            house.UniqPropKey,
				SerialNumber:           house.SerialNumber,
				SizeM2:                 house.SizeM2,
				MainUse:                house.MainUse,
				MainBuildingMaterial:   house.MainBuildingMaterial,
				ConstructCompletedDate: house.ConstructCompletedDate,
				NumberOfFloors:         house.NumberOfFloors,
				BuildingLamination:     house.BuildingLamination,
				Address:                house.Address,
				District:               house.District,
				Note:                   house.Note,
				CreatedAt:              house.CreatedAt,
				CreatedBy:              house.CreatedBy,
				UpdatedAt:              time.Now().UnixMilli(),
				UpdatedBy:              house.UpdatedBy,
				DeletedAt:              house.DeletedAt,
			}

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

func ImportExcelData(adapter *Tt.Adapter, resourcePath string) {
	fmt.Println("[Start] Beginning of process data")

	ReadHouseDataSheet(adapter, resourcePath)
	GetHouseAddressInBuySellData(&adapter, resourcePath)
	GetHouseAddressInRentData1(&adapter, resourcePath)
	GetHouseAddressInRentData2(&adapter, resourcePath)

	fmt.Println("[End] End process of import house data")
}
