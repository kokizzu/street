package rqProperty

import (
	"fmt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/W2/example/conf"
)

func (rq *Property) FindPropertiesBySerialNumber(serialNumber string) (res []*Property) {
	query := `SELECT * from ` + rq.sqlTableName() + `WHERE ` + rq.SerialNumber + ` = serialNumber`
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}

func (rq *Property) FindPropertiesByUniqueKey(uniqueSerialAndSize string) (res []*Property) {
	query := `SELECT * from ` + rq.sqlTableName() +
		` WHERE "uniquePropertyKey" = '` + uniqueSerialAndSize + `'`

	fmt.Println("Query == ", query)
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}
