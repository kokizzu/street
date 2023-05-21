package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"street/model/mAuth"
)

type Migrator struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
}

func RunMigration(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
) {
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
	}
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
}
