package model

import (
	"fmt"
	"os"
	"street/model/mProperty/rqProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

const (
	TableProperty   = `property`
	TablePropertyUS = `propertyUS`
	TablePropertyTW = `propertyTW`

	backupDir = `./upload`
)

func getBackupPropertyFileOutput(tableName string, offset, limit int) string {
	return fmt.Sprintf(
		"%s/%s_%d_%d.jsonline.lz4",
		backupDir, tableName, offset, limit,
	)
}

func BackupTruncateProperty(conn *Tt.Adapter, tableName string) {
	switch tableName {
	case TableProperty:
		err := backupPropertyDefault(conn)
		if err != nil {
			return
		}
		prop := rqProperty.NewProperty(conn)
		if !prop.Truncate() {
			L.LOG.Error(`failed to truncate table ` + prop.SqlTableName())
		}
	case TablePropertyUS:
		err := backupPropertyUS(conn)
		if err != nil {
			return
		}
		prop := rqProperty.NewPropertyUS(conn)
		if !prop.Truncate() {
			L.LOG.Error(`failed to truncate table ` + prop.SqlTableName())
		}
	case TablePropertyTW:
		err := backupPropertyTW(conn)
		if err != nil {
			return
		}
		prop := rqProperty.NewPropertyTW(conn)
		if !prop.Truncate() {
			L.LOG.Error(`failed to truncate table ` + prop.SqlTableName())
		}
	default:
		fmt.Println("invalid table name, must be property/propertyUS/propertyTW")
		return
	}

	fmt.Println("Backup completed successfully ")
}

func backupPropertyDefault(conn *Tt.Adapter) error {
	prop := rqProperty.NewProperty(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(TableProperty, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			L.LOG.Error(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				L.LOG.Error(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				L.LOG.Error(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}

func backupPropertyUS(conn *Tt.Adapter) error {
	prop := rqProperty.NewPropertyUS(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(TablePropertyUS, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			L.LOG.Error(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				L.LOG.Error(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				L.LOG.Error(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}

func backupPropertyTW(conn *Tt.Adapter) error {
	prop := rqProperty.NewPropertyTW(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(TablePropertyTW, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			L.LOG.Error(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				L.LOG.Error(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				L.LOG.Error(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}
