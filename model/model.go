package model

import (
	"log"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"street/model/mAuth/wcAuth"
	"street/model/mProperty"

	"street/model/mAuth"
)

type Migrator struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter
}

func RunMigration(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	propOltp *Tt.Adapter,
	propOlap *Ch.Adapter,
) {
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
		PropOltp: propOltp,
		PropOlap: propOlap,
	}
	mAuth.TarantoolTables[mAuth.TableUsers].PreUnique1MigrationHook = wcAuth.UniqueUsernameMigration
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.PropOltp.MigrateTables(mProperty.TarantoolTables)
	m.PropOlap.MigrateTables(mProperty.ClickhouseTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(authOltp *Tt.Adapter, authOlap *Ch.Adapter, propOltp *Tt.Adapter, propOlap *Ch.Adapter) {
	checkClickhouseTables(authOlap, mAuth.ClickhouseTables)
	checkClickhouseTables(propOlap, mProperty.ClickhouseTables)
	checkTarantoolTables(authOltp, mAuth.TarantoolTables)
	checkTarantoolTables(propOltp, mProperty.TarantoolTables)
}

func checkClickhouseTables(cConn *Ch.Adapter, tables map[Ch.TableName]*Ch.TableProp) {
	for tableName, props := range tables {
		tableName := string(tableName)
		resp, err := cConn.Query(`SELECT column_name, data_type FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = ` + S.Z(tableName))
		if err != nil {
			log.Fatalf(`please run TABLE migration on %v for clickhouse: %v`, tableName, err)
		}
		defer resp.Close()

		// fetch all columns
		columnNameTypes := map[string]string{}
		for resp.Next() {
			var columnName, dataType string
			err := resp.Scan(&columnName, &dataType)
			if err != nil {
				log.Fatalf(`please run COLUMN migration on %v for clickhouse: %v`, tableName, err)
			}
			columnNameTypes[columnName] = dataType
		}

		// check data type match
		for _, field := range props.Fields {
			existingType := columnNameTypes[field.Name]
			if existingType != string(field.Type) {
				log.Fatalf(`please run COLUMN_TYPE %v migration on %v for clickhouse: %v <> %v`, field.Name, tableName, existingType, field.Type)
			}
		}
	}
}

func checkTarantoolTables(tConn *Tt.Adapter, tables map[Tt.TableName]*Tt.TableProp) {
	for tableName := range tables {
		tableName := string(tableName)
		res, err := tConn.Call(`box.space.`+tableName+`:format`, []interface{}{})
		if err != nil {
			log.Fatalf(`please run TABLE migration on %v for tarantool: %v`, tableName, err)
		}
		rows := res.Tuples()
		if len(rows) != 1 {
			log.Fatalf(`please run FIELDS migration on %v for tarantool`, tableName)
		}
		columnNameTypes := map[string]string{}
		for _, row := range rows[0] {
			m, ok := row.(map[interface{}]interface{})
			if !ok {
				log.Fatalf(`please run FIELD migration on %v for tarantool: %v`, tableName, row)
			}
			columnNameTypes[X.ToS(m[`name`])] = X.ToS(m[`type`])
		}
		for _, field := range tables[Tt.TableName(tableName)].Fields {
			existingType := columnNameTypes[field.Name]
			if existingType != string(field.Type) {
				log.Fatalf(`please run FIELD_TYPE %v migration on %v for tarantool: %v <> %v`, field.Name, tableName, existingType, field.Type)
			}
		}
	}
}
