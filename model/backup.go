package model

import (
	"errors"
	"fmt"
	"os"
	"street/model/mProperty/rqProperty"

	"github.com/fatih/color"
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
		err := backupProperty(conn, rqProperty.NewProperty, tableName)
		if err != nil {
			L.LOG.Error("failed to backup property :", err)
			return
		}
	case TablePropertyUS:
		err := backupProperty(conn, rqProperty.NewPropertyUS, tableName)
		if err != nil {
			L.LOG.Error("failed to backup propertyUS :", err)
			return
		}
	case TablePropertyTW:
		err := backupProperty(conn, rqProperty.NewPropertyTW, tableName)
		if err != nil {
			L.LOG.Error("failed to backup propertyTW :", err)
			return
		}
	default:
		fmt.Println("invalid table name, must be property/propertyUS/propertyTW")
		return
	}

	fmt.Println("Backup completed successfully ")
}

type newPropFunc interface {
	CountTotalAllRows() uint64
	GetRows(offset, limit uint32) [][]any
}

func backupProperty[T newPropFunc](conn *Tt.Adapter, newFunc func(tt *Tt.Adapter) T, tableName string) error {
	prop := newFunc(conn)

	totalAllRows := prop.CountTotalAllRows()

	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		err := func() error {
			data := prop.GetRows(uint32(i), uint32(batchSize))
			if len(data) == 0 {
				return errors.New("no data")
			}

			outputFile := getBackupPropertyFileOutput(tableName, i, batchSize)
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

			fmt.Println(color.GreenString("[OK] Backed up to file " + outputFile))

			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
