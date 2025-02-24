package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"strings"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"github.com/pierrec/lz4/v4"
)

func RestoreTable(conn *Tt.Adapter, tableName string) {
	switch tableName {
	case TableProperty:
		err := restoreTableProperty(conn)
		if err != nil {
			log.Println("failed to restore property")
		}
	case TablePropertyUS:
		err := restoreTablePropertyUS(conn)
		if err != nil {
			log.Println("failed to restore propertyUS")
		}
	case TablePropertyTW:
		err := restoreTablePropertyTW(conn)
		if err != nil {
			log.Println("failed to restore propertyTW")
		}
	default:
		fmt.Println("invalid table name, must be property/propertyUS/propertyTW")
		return
	}
}

func getBackupFiles(tableName string) (files []string, err error) {
	err = filepath.Walk(backupDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), (tableName+`_`)) {
			files = append(files, info.Name())
		}

		return nil
	})
	if err != nil {
		L.LOG.Error("error walking through directory:", err)
		return
	}

	return
}

func restoreTableProperty(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TablePropertyUS)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			L.LOG.Error("failed to open file: ", err)
			return err
		}
		defer file.Close()

		reader := lz4.NewReader(file)

		scanner := bufio.NewScanner(reader)

		var idxLine int = 0
		for scanner.Scan() {
			line := scanner.Text()

			var row []any
			err := json.Unmarshal([]byte(line), &row)
			if err != nil {
				L.LOG.Error(`Err json.Unmarshal([]byte(line), &row): `, err)
				continue
			}

			if len(row) < 48 {
				continue
			}

			prop := wcProperty.NewPropertyMutator(conn)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])
			serialNumber := X.ToS(row[prop.IdxSerialNumber()])
			sizeM2 := X.ToS(row[prop.IdxSizeM2()])
			mainUse := X.ToS(row[prop.IdxMainUse()])
			mainBuildingMaterial := X.ToS(row[prop.IdxMainBuildingMaterial()])
			constructCompleteDate := X.ToS(row[prop.IdxConstructCompletedDate()])
			numberOfFloors := X.ToS(row[prop.IdxNumberOfFloors()])
			buildingLamination := X.ToS(row[prop.IdxBuildingLamination()])
			address := X.ToS(row[prop.IdxAddress()])
			district := X.ToS(row[prop.IdxDistrict()])
			note := X.ToS(row[prop.IdxNote()])
			coord := X.ToArr(row[prop.IdxCoord()])
			createdAt := X.ToI(row[prop.IdxCreatedAt()])
			createdBy := X.ToU(row[prop.IdxCreatedBy()])
			updatedAt := X.ToI(row[prop.IdxUpdatedAt()])
			deletedAt := X.ToI(row[prop.IdxDeletedAt()])
			formattedAddress := X.ToS(row[prop.IdxFormattedAddress()])
			lastPrice := X.ToS(row[prop.IdxLastPrice()])
			priceHistoriesSell := X.ToArr(row[prop.IdxPriceHistoriesSell()])
			priceHistoriesRent := X.ToArr(row[prop.IdxPriceHistoriesRent()])
			purpose := X.ToS(row[prop.IdxPurpose()])
			houseType := X.ToS(row[prop.IdxHouseType()])
			images := X.ToArr(row[prop.IdxImages()])
			bedroom := X.ToI(row[prop.IdxBedroom()])
			bathroom := X.ToI(row[prop.IdxBathroom()])
			agencyFeePercent := X.ToF(row[prop.IdxAgencyFeePercent()])
			floorList := X.ToArr(row[prop.IdxFloorList()])
			version := X.ToS(row[prop.IdxVersion()])
			yearBuilt := X.ToI(row[prop.IdxYearBuilt()])
			yearRenovated := X.ToI(row[prop.IdxYearRenovated()])
			totalSqFt := X.ToF(row[prop.IdxTotalSqft()])
			countyName := X.ToS(row[prop.IdxCountyName()])
			street := X.ToS(row[prop.IdxStreet()])
			city := X.ToS(row[prop.IdxCity()])
			state := X.ToS(row[prop.IdxState()])
			zip := X.ToS(row[prop.IdxZip()])
			propertyLastUpdatedDate := X.ToI(row[prop.IdxPropertyLastUpdatedDate()])
			approvalState := X.ToS(row[prop.IdxApprovalState()])
			countryCode := X.ToS(row[prop.IdxCountryCode()])
			livingRoom := X.ToI(row[prop.IdxLivingroom()])
			altitude := X.ToF(row[prop.IdxAltitude()])
			parking := X.ToF(row[prop.IdxAltitude()])
			depositFee := X.ToF(row[prop.IdxDepositFee()])
			minimumDurationYear := X.ToF(row[prop.IdxMinimumDurationYear()])
			otherFees := X.ToArr(row[prop.IdxOtherFees()])
			imageLabels := X.ToArr(row[prop.IdxImageLabels()])

			var isFoundByUniqPropKey = true

			prop.UniqPropKey = uniqPropKey
			if !prop.FindByUniqPropKey() {
				prop.SetUniqPropKey(uniqPropKey)
				isFoundByUniqPropKey = false
			}
			prop.SetSerialNumber(serialNumber)
			prop.SetSizeM2(sizeM2)
			prop.SetMainUse(mainUse)
			prop.SetMainBuildingMaterial(mainBuildingMaterial)
			prop.SetConstructCompletedDate(constructCompleteDate)
			prop.SetNumberOfFloors(numberOfFloors)
			prop.SetBuildingLamination(buildingLamination)
			prop.SetAddress(address)
			prop.SetDistrict(district)
			prop.SetNote(note)
			prop.SetCoord(coord)
			prop.SetCreatedAt(createdAt)
			prop.SetCreatedBy(createdBy)
			prop.SetUpdatedAt(updatedAt)
			prop.SetDeletedAt(deletedAt)
			prop.SetFormattedAddress(formattedAddress)
			prop.SetLastPrice(lastPrice)
			prop.SetPriceHistoriesSell(priceHistoriesSell)
			prop.SetPriceHistoriesRent(priceHistoriesRent)
			prop.SetPurpose(purpose)
			prop.SetHouseType(houseType)
			prop.SetImages(images)
			prop.SetBedroom(bedroom)
			prop.SetBathroom(bathroom)
			prop.SetAgencyFeePercent(agencyFeePercent)
			prop.SetFloorList(floorList)
			prop.SetVersion(version)
			prop.SetYearBuilt(yearBuilt)
			prop.SetYearRenovated(yearRenovated)
			prop.SetTotalSqft(totalSqFt)
			prop.SetCountyName(countyName)
			prop.SetStreet(street)
			prop.SetCity(city)
			prop.SetState(state)
			prop.SetZip(zip)
			prop.SetPropertyLastUpdatedDate(propertyLastUpdatedDate)
			prop.SetApprovalState(approvalState)
			prop.SetCountryCode(countryCode)
			prop.SetLivingroom(livingRoom)
			prop.SetAltitude(altitude)
			prop.SetParking(parking)
			prop.SetDepositFee(depositFee)
			prop.SetMinimumDurationYear(minimumDurationYear)
			prop.SetOtherFees(otherFees)
			prop.SetImageLabels(imageLabels)
			if len(row) == 48 {
				var attribute = rqProperty.PropertyAttribute{
					Title:          "",
					AgentPhone:     "",
					AgentBio:       "",
					BuildingSizeM2: 0,
					YardSizeM2:     0,
					ContactLink:    "",
				}

				attrByte, err := json.Marshal(attribute)
				if err != nil {
					L.LOG.Error("failed to marshal property attribute "+uniqPropKey+" : ", err)
				} else {
					prop.SetAttribute(string(attrByte))
				}
			} else {
				attribute := X.ToS(row[prop.IdxAttribute()])
				prop.SetAttribute(attribute)
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite property by uniqPropKey: " + uniqPropKey)
					continue
				}
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert property " + uniqPropKey)
					continue
				}
			}

			idxLine++
		}

		rowsInsertedStr := color.GreenString(fmt.Sprintf("[ %d ] Rows Inserted from file %s", idxLine, filePath))
		fmt.Println(rowsInsertedStr)
	}

	return nil
}

func restoreTablePropertyUS(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TablePropertyUS)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			L.LOG.Error("failed to open file: ", err)
			return err
		}
		defer file.Close()

		reader := lz4.NewReader(file)

		scanner := bufio.NewScanner(reader)

		var idxLine int = 0
		for scanner.Scan() {
			line := scanner.Text()

			var row []any
			err := json.Unmarshal([]byte(line), &row)
			if err != nil {
				L.LOG.Error(`Err json.Unmarshal([]byte(line), &row): `, err)
				continue
			}

			if len(row) < 48 {
				continue
			}

			prop := wcProperty.NewPropertyUSMutator(conn)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])
			serialNumber := X.ToS(row[prop.IdxSerialNumber()])
			sizeM2 := X.ToS(row[prop.IdxSizeM2()])
			mainUse := X.ToS(row[prop.IdxMainUse()])
			mainBuildingMaterial := X.ToS(row[prop.IdxMainBuildingMaterial()])
			constructCompleteDate := X.ToS(row[prop.IdxConstructCompletedDate()])
			numberOfFloors := X.ToS(row[prop.IdxNumberOfFloors()])
			buildingLamination := X.ToS(row[prop.IdxBuildingLamination()])
			address := X.ToS(row[prop.IdxAddress()])
			district := X.ToS(row[prop.IdxDistrict()])
			note := X.ToS(row[prop.IdxNote()])
			coord := X.ToArr(row[prop.IdxCoord()])
			createdAt := X.ToI(row[prop.IdxCreatedAt()])
			createdBy := X.ToU(row[prop.IdxCreatedBy()])
			updatedAt := X.ToI(row[prop.IdxUpdatedAt()])
			deletedAt := X.ToI(row[prop.IdxDeletedAt()])
			formattedAddress := X.ToS(row[prop.IdxFormattedAddress()])
			lastPrice := X.ToS(row[prop.IdxLastPrice()])
			priceHistoriesSell := X.ToArr(row[prop.IdxPriceHistoriesSell()])
			priceHistoriesRent := X.ToArr(row[prop.IdxPriceHistoriesRent()])
			purpose := X.ToS(row[prop.IdxPurpose()])
			houseType := X.ToS(row[prop.IdxHouseType()])
			images := X.ToArr(row[prop.IdxImages()])
			bedroom := X.ToI(row[prop.IdxBedroom()])
			bathroom := X.ToI(row[prop.IdxBathroom()])
			agencyFeePercent := X.ToF(row[prop.IdxAgencyFeePercent()])
			floorList := X.ToArr(row[prop.IdxFloorList()])
			version := X.ToS(row[prop.IdxVersion()])
			yearBuilt := X.ToI(row[prop.IdxYearBuilt()])
			yearRenovated := X.ToI(row[prop.IdxYearRenovated()])
			totalSqFt := X.ToF(row[prop.IdxTotalSqft()])
			countyName := X.ToS(row[prop.IdxCountyName()])
			street := X.ToS(row[prop.IdxStreet()])
			city := X.ToS(row[prop.IdxCity()])
			state := X.ToS(row[prop.IdxState()])
			zip := X.ToS(row[prop.IdxZip()])
			propertyLastUpdatedDate := X.ToI(row[prop.IdxPropertyLastUpdatedDate()])
			approvalState := X.ToS(row[prop.IdxApprovalState()])
			countryCode := X.ToS(row[prop.IdxCountryCode()])
			livingRoom := X.ToI(row[prop.IdxLivingroom()])
			altitude := X.ToF(row[prop.IdxAltitude()])
			parking := X.ToF(row[prop.IdxAltitude()])
			depositFee := X.ToF(row[prop.IdxDepositFee()])
			minimumDurationYear := X.ToF(row[prop.IdxMinimumDurationYear()])
			otherFees := X.ToArr(row[prop.IdxOtherFees()])
			imageLabels := X.ToArr(row[prop.IdxImageLabels()])

			var isFoundByUniqPropKey = true

			prop.UniqPropKey = uniqPropKey
			if !prop.FindByUniqPropKey() {
				prop.SetUniqPropKey(uniqPropKey)
				isFoundByUniqPropKey = false
			}
			prop.SetSerialNumber(serialNumber)
			prop.SetSizeM2(sizeM2)
			prop.SetMainUse(mainUse)
			prop.SetMainBuildingMaterial(mainBuildingMaterial)
			prop.SetConstructCompletedDate(constructCompleteDate)
			prop.SetNumberOfFloors(numberOfFloors)
			prop.SetBuildingLamination(buildingLamination)
			prop.SetAddress(address)
			prop.SetDistrict(district)
			prop.SetNote(note)
			prop.SetCoord(coord)
			prop.SetCreatedAt(createdAt)
			prop.SetCreatedBy(createdBy)
			prop.SetUpdatedAt(updatedAt)
			prop.SetDeletedAt(deletedAt)
			prop.SetFormattedAddress(formattedAddress)
			prop.SetLastPrice(lastPrice)
			prop.SetPriceHistoriesSell(priceHistoriesSell)
			prop.SetPriceHistoriesRent(priceHistoriesRent)
			prop.SetPurpose(purpose)
			prop.SetHouseType(houseType)
			prop.SetImages(images)
			prop.SetBedroom(bedroom)
			prop.SetBathroom(bathroom)
			prop.SetAgencyFeePercent(agencyFeePercent)
			prop.SetFloorList(floorList)
			prop.SetVersion(version)
			prop.SetYearBuilt(yearBuilt)
			prop.SetYearRenovated(yearRenovated)
			prop.SetTotalSqft(totalSqFt)
			prop.SetCountyName(countyName)
			prop.SetStreet(street)
			prop.SetCity(city)
			prop.SetState(state)
			prop.SetZip(zip)
			prop.SetPropertyLastUpdatedDate(propertyLastUpdatedDate)
			prop.SetApprovalState(approvalState)
			prop.SetCountryCode(countryCode)
			prop.SetLivingroom(livingRoom)
			prop.SetAltitude(altitude)
			prop.SetParking(parking)
			prop.SetDepositFee(depositFee)
			prop.SetMinimumDurationYear(minimumDurationYear)
			prop.SetOtherFees(otherFees)
			prop.SetImageLabels(imageLabels)
			if len(row) == 48 {
				var attribute = rqProperty.PropertyAttribute{
					Title:          "",
					AgentPhone:     "",
					AgentBio:       "",
					BuildingSizeM2: 0,
					YardSizeM2:     0,
					ContactLink:    "",
				}

				attrByte, err := json.Marshal(attribute)
				if err != nil {
					L.LOG.Error("failed to marshal property attribute "+uniqPropKey+" : ", err)
				} else {
					prop.SetAttribute(string(attrByte))
				}
			} else {
				attribute := X.ToS(row[prop.IdxAttribute()])
				prop.SetAttribute(attribute)
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite property by uniqPropKey: " + uniqPropKey)
					continue
				}
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert property " + uniqPropKey)
					continue
				}
			}

			idxLine++
		}

		rowsInsertedStr := color.GreenString(fmt.Sprintf("[ %d ] Rows Inserted from file %s", idxLine, filePath))
		fmt.Println(rowsInsertedStr)
	}

	return nil
}

func restoreTablePropertyTW(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TablePropertyUS)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			L.LOG.Error("failed to open file: ", err)
			return err
		}
		defer file.Close()

		reader := lz4.NewReader(file)

		scanner := bufio.NewScanner(reader)

		var idxLine int = 0
		for scanner.Scan() {
			line := scanner.Text()

			var row []any
			err := json.Unmarshal([]byte(line), &row)
			if err != nil {
				L.LOG.Error(`Err json.Unmarshal([]byte(line), &row): `, err)
				continue
			}

			if len(row) < 48 {
				continue
			}

			prop := wcProperty.NewPropertyTWMutator(conn)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])
			serialNumber := X.ToS(row[prop.IdxSerialNumber()])
			sizeM2 := X.ToS(row[prop.IdxSizeM2()])
			mainUse := X.ToS(row[prop.IdxMainUse()])
			mainBuildingMaterial := X.ToS(row[prop.IdxMainBuildingMaterial()])
			constructCompleteDate := X.ToS(row[prop.IdxConstructCompletedDate()])
			numberOfFloors := X.ToS(row[prop.IdxNumberOfFloors()])
			buildingLamination := X.ToS(row[prop.IdxBuildingLamination()])
			address := X.ToS(row[prop.IdxAddress()])
			district := X.ToS(row[prop.IdxDistrict()])
			note := X.ToS(row[prop.IdxNote()])
			coord := X.ToArr(row[prop.IdxCoord()])
			createdAt := X.ToI(row[prop.IdxCreatedAt()])
			createdBy := X.ToU(row[prop.IdxCreatedBy()])
			updatedAt := X.ToI(row[prop.IdxUpdatedAt()])
			deletedAt := X.ToI(row[prop.IdxDeletedAt()])
			formattedAddress := X.ToS(row[prop.IdxFormattedAddress()])
			lastPrice := X.ToS(row[prop.IdxLastPrice()])
			priceHistoriesSell := X.ToArr(row[prop.IdxPriceHistoriesSell()])
			priceHistoriesRent := X.ToArr(row[prop.IdxPriceHistoriesRent()])
			purpose := X.ToS(row[prop.IdxPurpose()])
			houseType := X.ToS(row[prop.IdxHouseType()])
			images := X.ToArr(row[prop.IdxImages()])
			bedroom := X.ToI(row[prop.IdxBedroom()])
			bathroom := X.ToI(row[prop.IdxBathroom()])
			agencyFeePercent := X.ToF(row[prop.IdxAgencyFeePercent()])
			floorList := X.ToArr(row[prop.IdxFloorList()])
			version := X.ToS(row[prop.IdxVersion()])
			yearBuilt := X.ToI(row[prop.IdxYearBuilt()])
			yearRenovated := X.ToI(row[prop.IdxYearRenovated()])
			totalSqFt := X.ToF(row[prop.IdxTotalSqft()])
			countyName := X.ToS(row[prop.IdxCountyName()])
			street := X.ToS(row[prop.IdxStreet()])
			city := X.ToS(row[prop.IdxCity()])
			state := X.ToS(row[prop.IdxState()])
			zip := X.ToS(row[prop.IdxZip()])
			propertyLastUpdatedDate := X.ToI(row[prop.IdxPropertyLastUpdatedDate()])
			approvalState := X.ToS(row[prop.IdxApprovalState()])
			countryCode := X.ToS(row[prop.IdxCountryCode()])
			livingRoom := X.ToI(row[prop.IdxLivingroom()])
			altitude := X.ToF(row[prop.IdxAltitude()])
			parking := X.ToF(row[prop.IdxAltitude()])
			depositFee := X.ToF(row[prop.IdxDepositFee()])
			minimumDurationYear := X.ToF(row[prop.IdxMinimumDurationYear()])
			otherFees := X.ToArr(row[prop.IdxOtherFees()])
			imageLabels := X.ToArr(row[prop.IdxImageLabels()])

			var isFoundByUniqPropKey = true

			prop.UniqPropKey = uniqPropKey
			if !prop.FindByUniqPropKey() {
				prop.SetUniqPropKey(uniqPropKey)
				isFoundByUniqPropKey = false
			}
			prop.SetSerialNumber(serialNumber)
			prop.SetSizeM2(sizeM2)
			prop.SetMainUse(mainUse)
			prop.SetMainBuildingMaterial(mainBuildingMaterial)
			prop.SetConstructCompletedDate(constructCompleteDate)
			prop.SetNumberOfFloors(numberOfFloors)
			prop.SetBuildingLamination(buildingLamination)
			prop.SetAddress(address)
			prop.SetDistrict(district)
			prop.SetNote(note)
			prop.SetCoord(coord)
			prop.SetCreatedAt(createdAt)
			prop.SetCreatedBy(createdBy)
			prop.SetUpdatedAt(updatedAt)
			prop.SetDeletedAt(deletedAt)
			prop.SetFormattedAddress(formattedAddress)
			prop.SetLastPrice(lastPrice)
			prop.SetPriceHistoriesSell(priceHistoriesSell)
			prop.SetPriceHistoriesRent(priceHistoriesRent)
			prop.SetPurpose(purpose)
			prop.SetHouseType(houseType)
			prop.SetImages(images)
			prop.SetBedroom(bedroom)
			prop.SetBathroom(bathroom)
			prop.SetAgencyFeePercent(agencyFeePercent)
			prop.SetFloorList(floorList)
			prop.SetVersion(version)
			prop.SetYearBuilt(yearBuilt)
			prop.SetYearRenovated(yearRenovated)
			prop.SetTotalSqft(totalSqFt)
			prop.SetCountyName(countyName)
			prop.SetStreet(street)
			prop.SetCity(city)
			prop.SetState(state)
			prop.SetZip(zip)
			prop.SetPropertyLastUpdatedDate(propertyLastUpdatedDate)
			prop.SetApprovalState(approvalState)
			prop.SetCountryCode(countryCode)
			prop.SetLivingroom(livingRoom)
			prop.SetAltitude(altitude)
			prop.SetParking(parking)
			prop.SetDepositFee(depositFee)
			prop.SetMinimumDurationYear(minimumDurationYear)
			prop.SetOtherFees(otherFees)
			prop.SetImageLabels(imageLabels)
			if len(row) == 48 {
				var attribute = rqProperty.PropertyAttribute{
					Title:          "",
					AgentPhone:     "",
					AgentBio:       "",
					BuildingSizeM2: 0,
					YardSizeM2:     0,
					ContactLink:    "",
				}

				attrByte, err := json.Marshal(attribute)
				if err != nil {
					L.LOG.Error("failed to marshal property attribute "+uniqPropKey+" : ", err)
				} else {
					prop.SetAttribute(string(attrByte))
				}
			} else {
				attribute := X.ToS(row[prop.IdxAttribute()])
				prop.SetAttribute(attribute)
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite property by uniqPropKey: " + uniqPropKey)
					continue
				}
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert property " + uniqPropKey)
					continue
				}
			}

			idxLine++
		}

		rowsInsertedStr := color.GreenString(fmt.Sprintf("[ %d ] Rows Inserted from file %s", idxLine, filePath))
		fmt.Println(rowsInsertedStr)
	}

	return nil
}
