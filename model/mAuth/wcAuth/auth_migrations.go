package wcAuth

import (
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"street/model/mAuth/rqAuth"
)

func UniqueUsernameMigration(a *Tt.Adapter) {
	r := rqAuth.Users{}
	w := &UsersMutator{}
	query := `-- ` + L.CallerInfo().String() + `
SELECT ` + r.SqlId() + `
, ` + r.SqlEmail() + `
FROM ` + r.SqlTableName() + `
WHERE ` + r.SqlUserName() + ` = ''
	OR ` + r.SqlUserName() + ` IS NULL
`
	okCount := 0
	total := 0
	// find all empty username
	_ = a.QuerySql(query, func(row []any) {
		total++
		id := X.ToI(row[0])
		email := X.ToS(row[1])

		// generate username from email
		newUsername := w.SetGenUniqueUsernameByEmail(email, id)
		res := a.ExecSql(`-- ` + L.CallerInfo().String() + `
UPDATE ` + r.SqlTableName() + `
SET ` + r.SqlUserName() + ` = ` + S.Q(newUsername) + `
WHERE ` + r.SqlId() + ` = ` + I.ToS(id))
		if X.ToI(res[`row_count`]) == 1 {
			okCount++
		}
		if X.ToS(res[`error`]) != `` {
			L.Print(res)
		}
	})
	if total > 0 {
		L.Print(`success`, okCount, `of`, total)
	}
}
