package model

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zImport"
	"strings"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

func RestoreTable(conn *Tt.Adapter, tableName string) {
	switch tableName {
	case TableProperty:
		if err := restoreTableProperty(conn); err != nil {
			L.LOG.Error("failed to restore property:", err)
		}
	case TablePropertyUS:
		if err := restoreTablePropertyUS(conn); err != nil {
			L.LOG.Error("failed to restore propertyUS:", err)
		}
	case TablePropertyTW:
		if err := restoreTablePropertyTW(conn); err != nil {
			L.LOG.Error("failed to restore propertyTW:", err)
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

	if len(files) == 0 {
		err = errors.New("no backup files available")
		return
	}

	return
}

func restoreTableProperty(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TableProperty)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		err := func() error {
			defer subTaskPrint("Restore Table " + TableProperty + " from " + backupDir + "/" + fileName)()

			filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &zImport.ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()

				line := scanner.Text()
				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
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

					prop := wcProperty.NewPropertyMutator(conn)
					prop.FromUncensoredArray(row)

					if prop.FindByUniqPropKey() {
						if !prop.DoOverwriteByUniqPropKey() {
							errMsg := "error while overwrite " + TableProperty + " by uniqPropKey: " + prop.UniqPropKey
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !prop.DoUpsert() {
						errMsg := "error while upsert " + TableProperty + " " + prop.UniqPropKey
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}

				idxLine++
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func restoreTablePropertyUS(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TablePropertyUS)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		err := func() error {
			defer subTaskPrint("Restore Table " + TablePropertyUS + " from " + backupDir + "/" + fileName)()

			filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &zImport.ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()

				line := scanner.Text()
				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
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

					if prop.FindByUniqPropKey() {
						if !prop.DoOverwriteByUniqPropKey() {
							errMsg := "error while overwrite " + TablePropertyUS + " by uniqPropKey: " + prop.UniqPropKey
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !prop.DoUpsert() {
						errMsg := "error while upsert " + TablePropertyUS + " " + prop.UniqPropKey
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}

				idxLine++
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func restoreTablePropertyTW(conn *Tt.Adapter) error {
	backupFiles, err := getBackupFiles(TablePropertyTW)
	if err != nil {
		return err
	}

	for _, fileName := range backupFiles {
		err := func() error {
			defer subTaskPrint("Restore Table " + TablePropertyTW + " from " + backupDir + "/" + fileName)()

			filePath := fmt.Sprintf("%s/%s", backupDir, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &zImport.ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()

				line := scanner.Text()
				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
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

					if prop.FindByUniqPropKey() {
						if !prop.DoOverwriteByUniqPropKey() {
							errMsg := "error while overwrite " + TablePropertyTW + " by uniqPropKey: " + prop.UniqPropKey
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !prop.DoUpsert() {
						errMsg := "error while upsert " + TablePropertyTW + " " + prop.UniqPropKey
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}

				idxLine++
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func subTaskPrint(str string) func() {
	startBlueStat := color.BlueString(`  [Start] ` + str)
	endBlueStat := color.BlueString(`  [End] ` + str)
	fmt.Println(startBlueStat)
	return func() {
		fmt.Println(endBlueStat)
	}
}
