package zUpgrade

import (
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"street/model/mAuth"
	"street/model/mAuth/wcAuth"
)

func UserSessionToMemtx(t *Tt.Adapter) {
	users := [][]any{}
	sessions := [][]any{}
	t.QuerySql(`SELECT * 
FROM "users"`, func(row []any) {
		users = append(users, row)
	})

	t.QuerySql(`SELECT *
FROM "sessions"`, func(row []any) {
		sessions = append(sessions, row)
	})
	t.ExecSql(`DROP TABLE "users"`)
	t.ExecSql(`DROP TABLE "sessions"`)
	t.MigrateTables(mAuth.TarantoolTables)
	t = &Tt.Adapter{
		Reconnect:  t.Reconnect,
		Connection: t.Reconnect(),
	}
	inserted, failed := 0, 0
	for _, row := range users {
		m := wcAuth.NewUsersMutator(t)
		m.FromArray(row)
		if m.DoInsert() {
			inserted++
		} else {
			failed++
		}
	}
	for _, row := range sessions {
		m := wcAuth.NewSessionsMutator(t)
		m.FromArray(row)
		if m.DoInsert() {
			inserted++
		} else {
			failed++
		}
	}
	L.Print(`inserted:`, inserted, `failed:`, failed)
}
