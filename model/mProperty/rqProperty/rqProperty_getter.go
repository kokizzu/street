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
	const comment = `-- Property) FindPropertiesBySerialNumber`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + ` 
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
	const comment = `-- Property) FindAllProperties`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName()
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

func (rq *PropertyHistory) FindAllPropertyHistories() (res []*PropertyHistory) {
	const comment = `-- PropertyHistory) FindAllPropertyHistories`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + `FROM ` + rq.SqlTableName()
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyHistory{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (rq *PropertyHistory) FindAllPropertyHistoriesWithBlankSerial() (res []*PropertyHistory) {
	const comment = `-- PropertyHistory) FindAllPropertyHistoriesWithBlankSerial`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + ` 
WHERE ` + rq.SqlSerialNumber() + ` = ''`
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyHistory{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (p *Property) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Property) FindByPagination`

	validFields := PropertyFieldTypeMap
	whereAndSql := out.WhereAndSql(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySql(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (p *Property) FindByLatLong(lat float64, long float64, limit int, offset int, callback func(row []any) bool) bool {
	const prefix = `Property) FindByLatLong`
	p.Coord = []any{lat, long}
	res, err := p.Adapter.Select(p.SpaceName(), p.SpatialIndexCoord(), uint32(offset), uint32(limit), tarantool.IterNeighbor, p.Coord)
	if L.IsError(err, prefix+` failed: `+p.SpaceName()) {
		return false
	}
	for _, row := range res.Tuples() {
		if !callback(row) {
			break
		}
	}
	return true
}
func (p *PropertyHistory) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- PropertyHistory) FindByPagination`

	validFields := PropertyHistoryFieldTypeMap
	whereAndSql := out.WhereAndSql(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySql(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (rq *PropertyHistory) FindByPropertyKey(key string) (res []*PropertyHistory) {
	const comment = `-- PropertyHistory) FindByPropertyKey`

	query := comment + `
SELECT` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + `
WHERE ` + rq.SqlPropertyKey() + ` = ` + S.Z(key)
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyHistory{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}

func (rq *PropertyHistory) FindBySerialNumber(key string) (res []*PropertyHistory) {
	const comment = `-- PropertyHistory) FindBySerialNumber`

	query := comment + `
SELECT
` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + `
WHERE ` + rq.SqlSerialNumber() + `=` + S.Z(key) + ``
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyHistory{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}
