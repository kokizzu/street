package model

import (
	"fmt"
	"log"
	"os"
	"street/model/mProperty/rqProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/rasky/go-lzo"
)

func BackupProperty(conn *Tt.Adapter, outputFile string) {
	prop := rqProperty.NewProperty(conn)
	data := prop.GetRows(1000)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	compressedData := lzo.Compress1X(jsonData)

	err = os.WriteFile(outputFile, compressedData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Backup completed successfully:", outputFile)
}
