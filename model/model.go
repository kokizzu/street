package model

import (
	"fmt"
	"street/model/mAuth/wcAuth"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"time"

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
	mAuth.TarantoolTables[mAuth.TableUsers].PreReformatMigrationHook = wcAuth.UniqueUsernameMigration
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.PropOltp.MigrateTables(mProperty.TarantoolTables)
	m.PropOlap.MigrateTables(mProperty.ClickhouseTables)
}

func UpdateAddressesToHouse(adapter **Tt.Adapter, resourceFile string) {
	fmt.Println("Begin the process update address to house")
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

	for index, row := range rows {
		if index == 0 || index == 1 {
			continue
		}

		var district string = ""
		var address string = ""
		var serialPropertyNumber string = ""

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
			// Get list of property based on house serial number
			existingHouseData := propertyMutator.FindPropertiesBySerialNumber(serialPropertyNumber)

			for _, house := range existingHouseData {
				fmt.Println("House data -> ", house)

				if house.Address != "" && house.District != "" {
					continue
				} else {
					house.Address = address
					house.District = district

					// Updated house
					dataMutator := wcProperty.NewPropertyMutator(*adapter)
					dataMutator.Property = rqProperty.Property{
						Adapter:                *adapter,
						Id:                     house.Id,
						UniquePropertyKey:      house.UniquePropertyKey,
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
					dataMutator.DoOverwriteById()
				}
			}
		}
	}

	fmt.Println("End the process update address to house")

}

func ImportExcelData(adapter *Tt.Adapter, resourcePath string) {
	fmt.Println("Import excel data")
	fmt.Println("Resource path => " + resourcePath)
	// TODO path files could be changed
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
	fmt.Println("Begin the process of import house data")
	for index, row := range rows {
		if index == 0 || index == 1 {
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
			fmt.Println(propertyMutator)
		}

		// Build unique key with serial number and size
		uniqueSerialNumber := propertyMutator.SerialNumber +
			"#" + propertyMutator.SizeM2
		propertyMutator.UniquePropertyKey = uniqueSerialNumber

		// Check if unique property key is existed
		existingProperties := propertyMutator.FindPropertiesByUniqueKey(uniqueSerialNumber)
		if len(existingProperties) > 0 || uniqueSerialNumber == "#" {
			continue
		} else {
			propertyMutator.DoInsert()
		}
	}
	fmt.Println("End process of import house data")
	UpdateAddressesToHouse(&adapter, resourcePath)
}
