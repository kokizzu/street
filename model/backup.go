package model

import (
	"fmt"
	"log"
	"os"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/rasky/go-lzo"
)

func getBackupPropertyFileOutput(outputDir, tableName string, offset, limit uint64) string {
	return fmt.Sprintf(
		"./%s/%s_%d_%d.jsonline.lzo",
		outputDir, tableName, limit, offset,
	)
}

func BackupProperty(conn *Tt.Adapter, outputDir, tableName string) {
	// TODO: get total all rows
	switch tableName {
	case string(mProperty.TableProperty):
		backupPropertyDefault(conn, outputDir)
		break
	default:
		fmt.Println("Invalid table name")
		return
	}

	fmt.Println("Backup completed successfully ")
}

func backupPropertyDefault(conn *Tt.Adapter, outputDir string) {
	fmt.Println(`TODO: create excel file and then compress with LZO`)

	prop := rqProperty.NewProperty(conn)

	outputFile := getBackupPropertyFileOutput(outputDir, string(mProperty.TableProperty), 0, 10000)
	data := prop.GetRows(0, 10000)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	compressedData := lzo.Compress1X(jsonData)

	err = os.WriteFile(outputFile, compressedData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
