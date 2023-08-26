package rqAuth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"street/model/zCrud"
)

func (u *Users) CheckPassword(pass string) error {
	return S.CheckPassword(u.Password, pass)
}

func (s *Sessions) AllActiveSession(userId uint64, now int64) (res []*Sessions) {
	const comment = `-- Sessions) AllActiveSession`

	query := comment + `
SELECT ` + s.SqlSelectAllFields() + `
FROM ` + s.SqlTableName() + `
WHERE ` + s.SqlUserId() + ` = ` + I.UToS(userId) + `
	AND ` + s.SqlExpiredAt() + ` > ` + I.ToS(now)
	s.Adapter.QuerySql(query, func(row []any) {
		res = append(res, s.FromArray(row))
	})
	return
}

func (u *Users) CountUserRegisterToday() (res int64) {
	currentDate := time.Now()
	beginCurrentDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
	endCurrentDate := beginCurrentDate.AddDate(0, 0, 1).Add(-time.Second)

	beginDateUnix := strconv.FormatUint(uint64(beginCurrentDate.Unix()), 10)
	endDateUnix := strconv.FormatUint(uint64(endCurrentDate.Unix()), 10)

	queryCountRegisteredToday := `SELECT COUNT(*) FROM ` + u.SqlTableName() +
		` WHERE ` + u.SqlCreatedAt() + ` >= ` + beginDateUnix + ` and ` +
		u.SqlCreatedAt() + ` <= ` + endDateUnix

	fmt.Println(beginCurrentDate)
	fmt.Println(endCurrentDate)

	u.Adapter.QuerySql(queryCountRegisteredToday, func(row []any) {
		res = X.ToI(row[0])
	})

	return res
}

func (u *Users) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Users) FindByPagination`

	validFields := UsersFieldTypeMap
	whereAndSql := out.WhereAndSql(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + u.SqlTableName() + whereAndSql + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySql(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + u.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}
