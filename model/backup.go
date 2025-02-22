package model

import (
	"fmt"
	"log"
	"os"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"strings"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/pierrec/lz4/v4"
)

func getBackupPropertyFileOutput(outputDir, tableName string, offset, limit int) string {
	return fmt.Sprintf(
		"./%s%s_%d_%d.jsonline.lz4",
		outputDir, tableName, offset, limit,
	)
}

func BackupProperty(conn *Tt.Adapter, outputDir, tableName string) {
	if strings.HasPrefix(outputDir, `/`) {
		outputDir = strings.TrimLeft(outputDir, `/`)
	}

	if !strings.HasSuffix(outputDir, `/`) && outputDir != `` {
		outputDir += `/`
	}

	switch tableName {
	case string(mProperty.TableProperty):
		backupPropertyDefault(conn, outputDir)
	default:
		fmt.Println("Invalid table name")
		return
	}

	fmt.Println("Backup completed successfully ")
}

func backupPropertyDefault(conn *Tt.Adapter, outputDir string) {
	prop := rqProperty.NewProperty(conn)

	totalAllRows := prop.CountTotalAllRows()
	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		data := prop.GetRows(uint32(i), uint32(batchSize))

		outputFile := getBackupPropertyFileOutput(outputDir, string(mProperty.TableProperty), i, batchSize)
		file, err := os.Create(outputFile)
		if err != nil {
			log.Println(`Err os.Create(outputFile): `, err)
			return
		}
		defer file.Close()

		lz4Writer := lz4.NewWriter(file)
		defer lz4Writer.Close()

		for _, row := range data {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				log.Println(`Err json.Marshal(row): `, err)
				continue
			}

			_, err = lz4Writer.Write(append(jsonRow, '\n'))
			if err != nil {
				log.Println(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
				continue
			}
		}
	}
}
