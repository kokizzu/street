package rqAuth

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"

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

func (u *Users) FindByPagination(in *zCrud.PaginationIn, out *zCrud.PaginationOut) (res [][]any) {
	limitOffsetSql := out.LimitOffsetSql(in)

	whereAndSql, orderBySql := out.WhereOrderSql(in.Filters, in.Order, map[string]string{}) // TODO: u.FieldsTypeMap())

	query := `-- ` + L.CallerInfo().String() + `
SELECT ` + u.SqlSelectAllUncensoredFields() + `
FROM ` + u.SqlTableName() + `
` + whereAndSql + `
` + orderBySql + `
` + limitOffsetSql
	u.Adapter.QuerySql(query, func(row []any) {
		res = append(res, row)
	})

	out.CalculatePages(len(res))
	return
}
