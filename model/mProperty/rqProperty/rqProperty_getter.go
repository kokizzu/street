package rqProperty

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"
	"github.com/tidwall/gjson"

	"street/conf"
	"street/model/mProperty"
	"street/model/zCrud"
)

// TODO: this is slow, use tarantool-go api instead

type PropertyWithNote struct {
	*Property
	ContactEmail string `json:"contactEmail" form:"contactEmail" query:"contactEmail" long:"contactEmail" msg:"contactEmail"`
	ContactPhone string `json:"contactPhone" form:"contactPhone" query:"contactPhone" long:"contactPhone" msg:"contactPhone"`
	About        string `json:"about" form:"about" query:"about" long:"about" msg:"about"`
	Image3dUrl   string `json:"image3dUrl" form:"image3dUrl" query:"image3dUrl" long:"image3dUrl" msg:"image3dUrl"`
}

func (pwn *PropertyWithNote) FromArray(row []any) {
	pwn.Property.FromArray(row)
	pwn.NormalizeFloorList()
	var nt mProperty.PropertyNote
	if json.Unmarshal([]byte(pwn.Property.Note), &nt) != nil {
		pwn.About = nt.About
		pwn.ContactEmail = nt.ContactEmail
		pwn.ContactPhone = nt.ContactPhone
	}
}

func (rq *Property) ToPropertyWithNote() PropertyWithNote {
	return PropertyWithNote{
		Property:     rq,
		ContactEmail: gjson.Get(rq.Note, `contactEmail`).String(),
		ContactPhone: gjson.Get(rq.Note, `contactPhone`).String(),
		About:        gjson.Get(rq.Note, `about`).String(),
		Image3dUrl:   ``,
	}
}

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

func (rq *Property) FindAllPropertiesOffsetLimit(offset, limit int) (res []*Property) {
	const comment = `-- Property) FindAllProperties`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + `
ORDER BY "id"
LIMIT ` + X.ToS(offset) + `,` + X.ToS(limit)
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
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
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

func (p *Property) FindByPaginationWithNote(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Property) FindByPaginationWithNote`

	validFields := PropertyFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	L.Print(`Query:`, queryRows)
	p.Adapter.QuerySql(queryRows, func(row []any) {
		var nt mProperty.PropertyNote
		row[0] = X.ToS(row[0]) // ensure id is string
		if nil != json.Unmarshal([]byte(X.ToS(row[p.IdxNote()])), &nt) {
			nt.About = X.ToS(row[p.IdxNote()])
		}

		idxAttr := meta.GetIdxByName(mProperty.Attribute)
		var attr mProperty.PropertyAttribute
		if nil != json.Unmarshal([]byte(X.ToS(row[idxAttr])), &attr) {
			row[idxAttr] = attr
		}

		row[p.IdxNote()] = nt
		res = append(res, row)
	})

	return
}

func (p *PropertyUS) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- PropertyUS) FindByPagination`

	validFields := PropertyFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
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

func (p *Property) FindOwnedByPagination(ownerId uint64, in *zCrud.PagerIn, out *zCrud.PagerOut) (res []PropertyWithNote) {
	const comment = `-- Property) FindOwnedByPagination`

	validFields := PropertyFieldTypeMap
	if in.Filters == nil {
		in.Filters = map[string][]string{}
	}
	// override owner
	in.Filters[mProperty.CreatedBy] = []string{`=`, X.ToS(ownerId)}
	in.Filters[mProperty.DeletedAt] = []string{`=`, `0`}
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	res = make([]PropertyWithNote, 0, out.PerPage)
	count := 0

	queryRows := comment + `
SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		res = append(res, PropertyWithNote{Property: &Property{}})
		res[count].FromArray(row)
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

func (p *Property) FindByLatLongAndCountry(propAdapter *Tt.Adapter, countryCode string, lat float64, long float64, limit int, offset int, callback func(row []any) bool) bool {

	// Default is taiwan property data
	prefix := `Property) FindByLatLongAndCountry`

	// Switch case if the country code is US
	if countryCode == "US" {
		prefix = `PropertyUS) FindByLatLong`
		prop := NewPropertyUS(propAdapter)
		return prop.FindPropUSByLatLong(lat, long, limit, offset, callback)
	}

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

func (p *PropertyUS) FindPropUSByLatLong(lat float64, long float64, limit int, offset int, callback func(row []any) bool) bool {
	const prefix = `PropertyUS) FindByLatLong`
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
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
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

func ConvertTxTime(input string) string {

	// In case if format time is different, return the original time
	if len(input) != 7 {
		return input
	}

	year := input[:3]
	month := input[3:5]
	day := input[5:7]

	// Construct the date string in YYY/MM/DD format
	standardChineseTime := fmt.Sprintf("%s/%s/%s", year, month, day)

	return standardChineseTime
}

func (rq *PropertyHistory) FindBySerialNumber(key string) (res []*PropertyHistory) {
	const comment = `-- PropertyHistory) FindBySerialNumber`

	query := comment + `
SELECT
` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + `
WHERE ` + rq.SqlSerialNumber() + `=` + S.Z(key) + `
  OR ` + rq.SqlPropertyKey() + `=` + S.Z(key)
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyHistory{}
		obj.FromArray(row)
		obj.TransactionTime = ConvertTxTime(obj.TransactionTime)
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
	floorsArr := []any{}
	for _, floor := range p.FloorList {
		floorObj := M.SX{}
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
				roomsArr := []M.SX{}
				for _, room := range rooms {
					room, ok := room.(map[any]any)
					if !ok {
						continue
					}
					roomObj := M.SX{}
					for key, val := range room {
						strKey, ok := key.(string)
						if !ok {
							continue
						}
						roomObj[strKey] = val
					}
					roomsArr = append(roomsArr, roomObj)
				}
				val = roomsArr
			}
			floorObj[strKey] = val
		}
		floorsArr = append(floorsArr, floorObj)
	}
	p.FloorList = floorsArr
}

func (p *PropertyUS) NormalizeFloorList() {
	floorsArr := []any{}
	for _, floor := range p.FloorList {
		floorObj := M.SX{}
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
				roomsArr := []M.SX{}
				for _, room := range rooms {
					room, ok := room.(map[any]any)
					if !ok {
						continue
					}
					roomObj := M.SX{}
					for key, val := range room {
						strKey, ok := key.(string)
						if !ok {
							continue
						}
						roomObj[strKey] = val
					}
					roomsArr = append(roomsArr, roomObj)
				}
				val = roomsArr
			}
			floorObj[strKey] = val
		}
		floorsArr = append(floorsArr, floorObj)
	}
	p.FloorList = floorsArr
}

func (p *UserPropLikes) LikedMap(propIds []uint64) (res map[uint64]bool) {
	query := `-- UserPropLikes) LikedMap
SELECT ` + p.SqlPropId() + `
FROM ` + p.SqlTableName() + `
WHERE ` + p.SqlUserId() + ` = ` + X.ToS(p.UserId) + `
  AND ` + p.SqlPropId() + ` IN (` + A.UIntJoin(propIds, `,`) + `)`

	res = map[uint64]bool{}
	p.Adapter.QuerySql(query, func(row []any) {
		propId := X.ToU(row[0])
		res[propId] = true
	})

	return
}

func (p *PropLikeCount) CountMap(propIds []uint64) (res map[uint64]int64) {
	query := `-- PropLikeCount) CountMap
SELECT ` + p.SqlPropId() + `, ` + p.SqlCount() + `
FROM ` + p.SqlTableName() + `
WHERE ` + p.SqlPropId() + ` IN (` + A.UIntJoin(propIds, `,`) + `)`

	res = map[uint64]int64{}
	p.Adapter.QuerySql(query, func(row []any) {
		propId := X.ToU(row[0])
		count := X.ToI(row[1])
		res[propId] = count
	})

	return
}

func (p *PropertyExtraUS) Pagination(offset, limit int) []PropertyExtraUS {
	query := `-- PropertyExtraUS) Pagination`
	query += `
SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() + `
ORDER BY ` + p.SqlId() + `
LIMIT ` + X.ToS(limit) + ` OFFSET ` + X.ToS(offset)

	res := make([]PropertyExtraUS, 0, limit)
	p.Adapter.QuerySql(query, func(row []any) {
		obj := PropertyExtraUS{}
		obj.FromArray(row)
		res = append(res, obj)
	})

	return res
}

func (p *PropertyUS) ToProperty() *Property {
	out := &Property{}
	backupAdapter := p.Adapter
	p.Adapter = nil
	M.FastestCopyStruct(p, out)
	p.Adapter = backupAdapter
	return out
}

func (p *PropertyExtraUS) ToPropertyExtra() *PropertyExtraUS {
	out := &PropertyExtraUS{}
	backupAdapter := p.Adapter
	p.Adapter = nil
	M.FastestCopyStruct(p, out)
	p.Adapter = backupAdapter
	return out
}

func (p *PropertyTW) ToProperty() *Property {
	out := &Property{}
	backupAdapter := p.Adapter

	M.FastestCopyStruct(p, out)
	p.Adapter = backupAdapter
	return out
}

func (rq *PropertyTW) FindAllPropertiesOffsetLimit(offset, limit int) (res []*PropertyTW) {
	const comment = `-- PropertyTW) FindAllProperties`

	query := comment + `
SELECT ` + rq.SqlSelectAllFields() + `
FROM ` + rq.SqlTableName() + `
ORDER BY "id"
LIMIT ` + X.ToS(offset) + `,` + X.ToS(limit)
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &PropertyTW{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (p *PropertyTW) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- PropertyTW) FindByPagination`

	validFields := PropertyFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
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

func (p *Property) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + p.SqlTableName() + `
	LIMIT 1`

	p.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (p *Property) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (p *Property) Truncate() bool {
	return p.Adapter.ExecBoxSpace(`property:truncate`, A.X{})
}

func (p *PropertyUS) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + p.SqlTableName() + `
	LIMIT 1`

	p.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (p *PropertyUS) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (p *PropertyUS) Truncate() bool {
	return p.Adapter.ExecBoxSpace(`propertyUS:truncate`, A.X{})
}

func (p *PropertyTW) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + p.SqlTableName() + `
	LIMIT 1`

	p.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (p *PropertyTW) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (p *PropertyTW) Truncate() bool {
	return p.Adapter.ExecBoxSpace(`propertyTW:truncate`, A.X{})
}

func (p *Property) RemovePriceHistory() bool {
	p.Adapter.ExecSql(`UPDATE ` + p.SqlTableName() + ` SET ` + p.SqlPriceHistoriesRent() + ` = [], ` + p.SqlPriceHistoriesSell() + ` = []`)
	return true
}

func (p *PropertyTW) RemovePriceHistory() bool {
	p.Adapter.ExecSql(`UPDATE ` + p.SqlTableName() + ` SET ` + p.SqlPriceHistoriesRent() + ` = [], ` + p.SqlPriceHistoriesSell() + ` = []`)
	return true
}

func (p *PropertyUS) RemovePriceHistory() bool {
	p.Adapter.ExecSql(`UPDATE ` + p.SqlTableName() + ` SET ` + p.SqlPriceHistoriesRent() + ` = [], ` + p.SqlPriceHistoriesSell() + ` = []`)
	return true
}

func (p *Property) DeletePropertyJP() bool {
	p.Adapter.ExecSql(`DELETE FROM ` + p.SqlTableName() + ` WHERE ` + p.SqlCountryCode() + ` = 'JP'`)
	return true
}
