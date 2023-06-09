package rqAuth

import (
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"street/model/zCrud"
)

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

func (u *Users) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	out.PerPage = in.Limit()

	whereAndSql, orderBySql := out.WhereOrderSql(in.Filters, in.Order, map[string]Tt.DataType{}) // TODO: u.FieldsTypeMap())

	total := 0
	queryCount := `-- Users) FindByPagination
SELECT COUNT(1)
FROM ` + u.SqlTableName() + whereAndSql + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		total = int(X.ToI(row[0]))
		out.CalculatePages(total)
	})

	limitOffsetSql := out.LimitOffsetSql(in, total)

	queryRows := `-- Users) FindByPagination
SELECT ` + u.SqlSelectAllUncensoredFields() + `
FROM ` + u.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	u.Adapter.QuerySql(queryRows, func(row []any) {
		res = append(res, row)
	})

	return
}
