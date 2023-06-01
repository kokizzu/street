package rqProperty

import (
	"fmt"
	"street/conf"

	"github.com/kokizzu/gotro/L"
)

func (rq *Property) FindPropertiesBySerialNumber(serialNumber string) (res []*Property) {
	query := `SELECT ` + rq.sqlSelectAllFields() + ` from ` + rq.sqlTableName() + `WHERE "serialNumber" = '` + serialNumber + `'`
	if conf.IsDebug() {
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
	query := `SELECT ` + rq.sqlSelectAllFields() + ` from ` + rq.sqlTableName() +
		` WHERE "uniquePropertyKey" = '` + uniqueSerialAndSize + `'`

	fmt.Println("Query == ", query)
	if conf.IsDebug() {
		L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}
