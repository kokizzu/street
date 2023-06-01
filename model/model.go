package model

import (
	"fmt"
	"street/model/mProperty"
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
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.PropOltp.MigrateTables(mProperty.TarantoolTables)
	m.PropOlap.MigrateTables(mProperty.ClickhouseTables)
}

func ImportExcelData(adapter *Tt.Adapter) {
	fmt.Println("Import excel data")
	pathData := "/Volumes/DATA/Projects/HapSTR/test-files/a_lvr_land_a.xlsx"
	f, err := excelize.OpenFile(pathData)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("House")
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, row := range rows {
		fmt.Println("Row => ", index)
		if index == 0 || index == 1 {
			continue
		}

		propertyMutator := wcProperty.NewPropertyMutator(adapter)
		for colIndex, colCell := range row {
			fmt.Println("col index => ", colIndex)
			fmt.Println(colCell, "\t")
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
		if len(existingProperties) > 0 {
			continue
		} else {
			propertyMutator.DoInsert()
		}
	}
}
