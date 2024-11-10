package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/rs/zerolog"

	"street/model/mAuth/wcAuth"
	"street/model/mBusiness"
	"street/model/mProperty"
	"street/model/mStorage"

	"street/model/mAuth"
)

type Migrator struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter
	BusinessOltp *Tt.Adapter
	StorOltp *Tt.Adapter
}

func RunMigration(logger *zerolog.Logger, authOltp *Tt.Adapter, authOlap *Ch.Adapter, propOltp *Tt.Adapter, propOlap *Ch.Adapter, businessOltp *Tt.Adapter, storOltp *Tt.Adapter) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
		PropOltp: propOltp,
		PropOlap: propOlap,
		BusinessOltp: businessOltp,
		StorOltp: storOltp,
	}
	mAuth.TarantoolTables[mAuth.TableUsers].PreUnique1MigrationHook = wcAuth.UniqueUsernameMigration
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.PropOltp.MigrateTables(mProperty.TarantoolTables)
	m.PropOlap.MigrateTables(mProperty.ClickhouseTables)
	m.BusinessOltp.MigrateTables(mBusiness.TarantoolTables)
	m.StorOltp.MigrateTables(mStorage.TarantoolTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	propOltp *Tt.Adapter,
	propOlap *Ch.Adapter,
	businessOltp *Tt.Adapter,
	storOltp *Tt.Adapter,
) {
	Ch.CheckClickhouseTables(authOlap, mAuth.ClickhouseTables)
	Ch.CheckClickhouseTables(propOlap, mProperty.ClickhouseTables)
	Tt.CheckTarantoolTables(authOltp, mAuth.TarantoolTables)
	Tt.CheckTarantoolTables(propOltp, mProperty.TarantoolTables)
	Tt.CheckTarantoolTables(businessOltp, mBusiness.TarantoolTables)
	Tt.CheckTarantoolTables(storOltp, mStorage.TarantoolTables)
}
