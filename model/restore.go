package model

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

func RestoreTable(tableName string) {
	switch tableName {
	case TablePropertyUS:
		err := restoreTablePropertyUS()
		if err != nil {
			log.Println("failed to restore propertyUS")
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

func restoreTablePropertyUS() error {
	backupFiles, err := getBackupFiles(TablePropertyUS)
	if err != nil {
		return err
	}

	if len(backupFiles) > 0 {
		file1 := fmt.Sprintf("%s/%s", backupDir, backupFiles[0])
		file, err := os.Open(file1)
		if err != nil {
			L.LOG.Error("failed to open file: ", err)
			return err
		}
		defer file.Close()

		reader := lz4.NewReader(file)

		byteData, err := io.ReadAll(reader)
		if err != nil {
			L.LOG.Error("Err io.ReadAll(reader): ", err)
			return err
		}

		fmt.Println("Byte Data", string(byteData))
	}

	return nil
}
