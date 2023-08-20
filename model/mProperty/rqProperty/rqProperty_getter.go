package rqProperty

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"

	"street/conf"
	"street/model/mProperty"
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
		obj.NormalizeFloorList()
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
		obj.NormalizeFloorList()
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

func (p *Property) FindOwnedByPagination(ownerId uint64, in *zCrud.PagerIn, out *zCrud.PagerOut) (res []Property) {
	const comment = `-- Property) FindOwnedByPagination`

	validFields := PropertyFieldTypeMap
	if in.Filters == nil {
		in.Filters = map[string][]string{}
	}
	// override owner
	in.Filters[mProperty.CreatedBy] = []string{`=`, X.ToS(ownerId)}
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

	res = make([]Property, 0, out.PerPage)
	count := 0

	queryRows := comment + `
SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	L.Print(queryRows)
	p.Adapter.QuerySql(queryRows, func(row []any) {
		L.Describe(row)
		res = append(res, Property{})
		res[count].FromArray(row)
		res[count].NormalizeFloorList()
		count++
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

func (p *Property) NormalizeFloorList() {
	/*
		the structure is too deply nested just like in mongodb,
		this will be pain in the future

		[]interface {}{
		    uint64(0x48d4),
		    "1_11461081313297949017",
		    "",
		    "123",
		    "swimming pool, 2 car port",
		    "",
		    "",
		    "1",
		    "",
		    "",
		    "",
		    "near railroad",
		    []interface {}{
		        float64(13.5344844),
		        float64(144.8827904),
		    },
		    uint64(0x64d95f95),
		    uint64(0x1),
		    uint64(0x64d95f95),
		    uint64(0x1),
		    uint64(0x0),
		    "Gayuman Ave, Yigo, Guam",
		    "",
		    []interface {}{
		    },
		    []interface {}{
		    },
		    "rent",
		    "house",
		    []interface {}{
		        "/guest/files/E-___.png",
		    },
		    uint64(0x2),
		    uint64(0x1),
		    float64(1),
		    []interface {}{
		        map[interface {}]interface {}{ --> this cannot be converted to json
		            "baths":        float64(1),
		            "beds":         float64(2),
		            "floor":        float64(1),
		            "planImageUrl": "",
		            "rooms":        []interface {}{ --> this one also
		                map[interface {}]interface {}{
		                    "name":   "bedroom",
		                    "sizeM2": float64(12),
		                    "unit":   "m2",
		                },
		                map[interface {}]interface {}{
		                    "name":   "bedroom",
		                    "sizeM2": float64(13),
		                    "unit":   "m2",
		                },
		                map[interface {}]interface {}{
		                    "name":   "bathroom",
		                    "sizeM2": float64(8),
		                    "unit":   "m2",
		                },
		                map[interface {}]interface {}{
		                    "name":   "living room",
		                    "sizeM2": float64(11),
		                    "unit":   "m2",
		                },
		            },
		            "type": "floor",
		        },
		    },
		    "",
		}
	*/
	amsx := []any{}
	for _, floor := range p.FloorList {
		msx := M.SX{}
		maa, ok := floor.(map[any]any)
		if !ok {
			continue
		}
		for key, val := range maa {
			strKey, ok := key.(string)
			if !ok {
				continue
			}
			if strKey == `rooms` {
				rooms, ok := val.([]any)
				if !ok {
					continue
				}
				for _, room := range rooms {
					room, ok := room.(map[any]any)
					if !ok {
						continue
					}
					roomsx := map[string]any{}
					for key, val := range room {
						strKey, ok := key.(string)
						if !ok {
							continue
						}
						roomsx[strKey] = val
					}
					val = roomsx
				}
			}
			msx[strKey] = val
		}
		amsx = append(amsx, msx)
	}
	p.FloorList = amsx
}
