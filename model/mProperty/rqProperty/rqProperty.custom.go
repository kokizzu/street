package rqProperty

import (
	"street/conf"
)

func (rq *Property) FindPropertiesBySerialNumber(serialNumber string) (res []*Property) {
	query := `SELECT ` + rq.SqlSelectAllFields() + ` from ` + rq.SqlTableName() + `WHERE "serialNumber" = '` + serialNumber + `'`
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

func (rq *Property) FindPropertiesByUniqueKey(uniqueSerialAndSize string) (res []*Property) {
	query := `SELECT ` + rq.SqlSelectAllFields() + ` from ` + rq.SqlTableName() +
		` WHERE "uniquePropertyKey" = '` + uniqueSerialAndSize + `'`

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
