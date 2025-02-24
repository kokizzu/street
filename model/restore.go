package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"street/model/mProperty"
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

			rowsLength := len(row)
			tableLength := len(mProperty.TarantoolTables[mProperty.TableProperty].Fields)
			for {
				if rowsLength < tableLength {
					rowType := mProperty.TarantoolTables[mProperty.TableProperty].Fields[rowsLength].Type
					rowName := mProperty.TarantoolTables[mProperty.TableProperty].Fields[rowsLength].Name

					var rowDataToAppend any = ""

					switch rowType {
					case Tt.String:
						rowDataToAppend = ""
					case Tt.Boolean:
						rowDataToAppend = false
					case Tt.Double, Tt.Integer, Tt.Unsigned:
						rowDataToAppend = 0
					case Tt.Array:
						rowDataToAppend = []any{0, 0}
					default:
						rowDataToAppend = ""
					}

					switch rowName {
					case mProperty.Attribute, mProperty.Note:
						rowDataToAppend = "{}"
					}

					row = append(row, rowDataToAppend)
					rowsLength++
					continue
				}

				break
			}

			prop := wcProperty.NewPropertyUSMutator(conn)
			prop.FromUncensoredArray(row)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])

			var isFoundByUniqPropKey = true
			if !prop.FindByUniqPropKey() {
				isFoundByUniqPropKey = false
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite " + TableProperty + " by uniqPropKey: " + uniqPropKey)
					continue
				}
				rowOverrideStr := color.BlueString(fmt.Sprintf("[ %d ] Updated %s '%s'", idxLine, TableProperty, uniqPropKey))
				fmt.Println(rowOverrideStr)
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert " + TableProperty + " " + uniqPropKey)
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

			rowsLength := len(row)
			tableLength := len(mProperty.TarantoolTables[mProperty.TablePropertyUS].Fields)
			for {
				if rowsLength < tableLength {
					rowType := mProperty.TarantoolTables[mProperty.TablePropertyUS].Fields[rowsLength].Type
					rowName := mProperty.TarantoolTables[mProperty.TablePropertyUS].Fields[rowsLength].Name

					var rowDataToAppend any = ""

					switch rowType {
					case Tt.String:
						rowDataToAppend = ""
					case Tt.Boolean:
						rowDataToAppend = false
					case Tt.Double, Tt.Integer, Tt.Unsigned:
						rowDataToAppend = 0
					case Tt.Array:
						rowDataToAppend = []any{0, 0}
					default:
						rowDataToAppend = ""
					}

					switch rowName {
					case mProperty.Attribute, mProperty.Note:
						rowDataToAppend = "{}"
					}

					row = append(row, rowDataToAppend)
					rowsLength++
					continue
				}

				break
			}

			prop := wcProperty.NewPropertyUSMutator(conn)
			prop.FromUncensoredArray(row)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])

			var isFoundByUniqPropKey = true
			if !prop.FindByUniqPropKey() {
				isFoundByUniqPropKey = false
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite " + TablePropertyUS + " by uniqPropKey: " + uniqPropKey)
					continue
				}
				rowOverrideStr := color.BlueString(fmt.Sprintf("[ %d ] Updated %s '%s'", idxLine, TablePropertyUS, uniqPropKey))
				fmt.Println(rowOverrideStr)
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert " + TablePropertyUS + " " + uniqPropKey)
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

			rowsLength := len(row)
			tableLength := len(mProperty.TarantoolTables[mProperty.TablePropertyTW].Fields)
			for {
				if rowsLength < tableLength {
					rowType := mProperty.TarantoolTables[mProperty.TablePropertyTW].Fields[rowsLength].Type
					rowName := mProperty.TarantoolTables[mProperty.TablePropertyTW].Fields[rowsLength].Name

					var rowDataToAppend any = ""

					switch rowType {
					case Tt.String:
						rowDataToAppend = ""
					case Tt.Boolean:
						rowDataToAppend = false
					case Tt.Double, Tt.Integer, Tt.Unsigned:
						rowDataToAppend = 0
					case Tt.Array:
						rowDataToAppend = []any{0, 0}
					default:
						rowDataToAppend = ""
					}

					switch rowName {
					case mProperty.Attribute, mProperty.Note:
						rowDataToAppend = "{}"
					}

					row = append(row, rowDataToAppend)
					rowsLength++
					continue
				}

				break
			}

			prop := wcProperty.NewPropertyTWMutator(conn)
			prop.FromUncensoredArray(row)

			uniqPropKey := X.ToS(row[prop.IdxUniqPropKey()])

			var isFoundByUniqPropKey = true
			if !prop.FindByUniqPropKey() {
				isFoundByUniqPropKey = false
			}

			if isFoundByUniqPropKey {
				if !prop.DoOverwriteByUniqPropKey() {
					L.Print("failed to overwrite " + TablePropertyTW + " by uniqPropKey: " + uniqPropKey)
					continue
				}
				rowOverrideStr := color.BlueString(fmt.Sprintf("[ %d ] Updated %s '%s'", idxLine, TablePropertyTW, uniqPropKey))
				fmt.Println(rowOverrideStr)
			} else {
				if !prop.DoInsert() {
					L.Print("failed to insert " + TablePropertyTW + " " + uniqPropKey)
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
