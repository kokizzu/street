package rqProperty

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"

	"street/conf"
	"street/model/zCrud"
)

// TODO: this is slow, use tarantool-go api instead

func (rq *Property) FindPropertiesBySerialNumber(serialNumber string) (res []*Property) {
	query := `SELECT ` + rq.SqlSelectAllFields() + ` 
FROM ` + rq.SqlTableName() + `
WHERE ` + rq.SqlSerialNumber() + ` = ` + S.Q(serialNumber)
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (rq *Property) FindAllProperties() (res []*Property) {
	query := `SELECT ` + rq.SqlSelectAllFields() + `FROM ` + rq.SqlTableName()
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (p *Property) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	whereAndSql := z3.WhereAndSql(z2.Filters, PropertyFieldTypeMap)

	queryCount := `-- Property) FindByPagination
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySql(z2.Order, PropertyFieldTypeMap)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := `-- Property) FindByPagination
SELECT ` + z.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (p *Property) FindByLatLong(lat float64, long float64, limit int, offset int, callback func(row []any)) bool {
	p.Coord = []any{lat, long}
	res, err := p.Adapter.Select(p.SpaceName(), p.SpatialIndexCoord(), uint32(offset), uint32(limit), tarantool.IterNeighbor, p.Coord)
	if L.IsError(err, `Property) FindByLatLong failed: `+p.SpaceName()) {
		return false
	}
	for _, row := range res.Tuples() {
		callback(row)
	}
	return true
}
func (p *PropertyHistory) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	whereAndSql := z3.WhereAndSql(z2.Filters, PropertyFieldTypeMap)

	queryCount := `-- PropertyHistory) FindByPagination
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySql(z2.Order, PropertyFieldTypeMap)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := `-- PropertyHistory) FindByPagination
SELECT ` + z.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}
