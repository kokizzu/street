package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"street/model/mProperty"

	"street/model/mAuth"
)

type Migrator struct {
	UserOltp *Tt.Adapter
	UserOlap *Ch.Adapter
}

func RunMigration(
	userOltp *Tt.Adapter,
	userOlap *Ch.Adapter,
) {
	L.Print(`run migration..`)
	m := Migrator{
		UserOltp: userOltp,
		UserOlap: userOlap,
	}
	m.UserOltp.MigrateTables(mAuth.TarantoolTables)
	m.UserOlap.MigrateTables(mAuth.ClickhouseTables)
	m.UserOltp.MigrateTables(mProperty.TarantoolTables)
	m.UserOlap.MigrateTables(mProperty.ClickhouseTables)
}
