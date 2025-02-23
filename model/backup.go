package model

import (
	"fmt"
	"log"
	"os"
	"street/model/mProperty/rqProperty"
	"strings"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/pierrec/lz4/v4"
)

const (
	TableProperty   = `property`
	TablePropertyUS = `propertyUS`
	TablePropertyTW = `propertyTW`
)

func getBackupPropertyFileOutput(outputDir, tableName string, offset, limit int) string {
	return fmt.Sprintf(
		"./%s%s_%d_%d.jsonline.lz4",
		outputDir, tableName, offset, limit,
	)
}

func BackupTruncateProperty(conn *Tt.Adapter, outputDir, tableName string) {
	if strings.HasPrefix(outputDir, `/`) {
		outputDir = strings.TrimLeft(outputDir, `/`)
	}

	if !strings.HasSuffix(outputDir, `/`) && outputDir != `` {
		outputDir += `/`
	}

	switch tableName {
	case TableProperty:
		err := backupPropertyDefault(conn, outputDir)
		if err != nil {
			return
		}
		prop := rqProperty.NewProperty(conn)
		if !prop.Truncate() {
			log.Println(`failed to truncate table ` + prop.SqlTableName())
		}
	case TablePropertyUS:
		err := backupPropertyUS(conn, outputDir)
		if err != nil {
			return
		}
		prop := rqProperty.NewPropertyUS(conn)
		if !prop.Truncate() {
			log.Println(`failed to truncate table ` + prop.SqlTableName())
		}
	case TablePropertyTW:
		err := backupPropertyTW(conn, outputDir)
		if err != nil {
			return
		}
		prop := rqProperty.NewPropertyTW(conn)
		if !prop.Truncate() {
			log.Println(`failed to truncate table ` + prop.SqlTableName())
		}
	default:
		fmt.Println("invalid table name, must be property/propertyUS/propertyTW")
		return
	}

	fmt.Println("Backup completed successfully ")
}

func backupPropertyDefault(conn *Tt.Adapter, outputDir string) error {
	prop := rqProperty.NewProperty(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(outputDir, TableProperty, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			log.Println(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				log.Println(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				log.Println(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}

func backupPropertyUS(conn *Tt.Adapter, outputDir string) error {
	prop := rqProperty.NewPropertyUS(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(outputDir, TablePropertyUS, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			log.Println(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				log.Println(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				log.Println(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}

func backupPropertyTW(conn *Tt.Adapter, outputDir string) error {
	prop := rqProperty.NewPropertyTW(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(outputDir, TablePropertyTW, i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			log.Println(`Err os.Create(outputFile): `, err)
			return err
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				log.Println(`Err json.Marshal(row): `, err)
				return err
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				log.Println(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				return err
			}
		}
	}

	return nil
}
