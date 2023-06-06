package rqAuth

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

func (u *Users) CensorFields() {
	u.Password = ``
}

func (u *Users) CheckPassword(pass string) error {
	return S.CheckPassword(u.Password, pass)
}

func (s *Sessions) AllActiveSession(userId uint64, now int64) (res []*Sessions) {
	query := `-- ` + L.CallerInfo().String() + `
SELECT ` + s.SqlSelectAllFields() + `
FROM ` + s.SqlTableName() + `
WHERE ` + s.SqlUserId() + ` = ` + I.UToS(userId) + `
	AND ` + s.SqlExpiredAt() + ` > ` + I.ToS(now)
	s.Adapter.QuerySql(query, func(row []any) {
		res = append(res, s.FromArray(row))
	})
	return
}
