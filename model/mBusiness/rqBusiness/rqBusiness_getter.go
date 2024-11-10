package rqBusiness

import (
	"street/model/mBusiness"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

func (s *Sales) FindRevenuesByRealtorId() (revenues []*mBusiness.RealtorRevenue) {
	const comment = `-- Sales) FindRevenuesByRealtorId`

	whereAndSql := ` WHERE ` + s.SqlRealtorId() + ` = ` + I.UToS(s.RealtorId)

	query := comment + `
SELECT ` + s.SqlPrice() + `, ` + s.SqlPropertyId() + `, ` + s.SqlSalesDate() + `, ` + s.SqlCreatedAt() + `, ` + s.SqlUpdatedAt() + `
FROM ` + s.SqlTableName() + whereAndSql

	L.Print(`QUERY :`, query)
	s.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 5 {
			price := X.ToI(row[0])
			salesDate := X.ToS(row[2])

			var isExist bool = false

			for _, r := range revenues {
				if r.SalesDate == salesDate  {
					r.PropertyBought += 1
					r.Revenue += price

					isExist = true
				}
			}

			if !isExist {
				revenues = append(revenues, &mBusiness.RealtorRevenue{
					PropertyId: X.ToU(row[1]),
					Revenue: price,
					PropertyBought: 1,
					SalesDate: salesDate,
					CreatedAt: X.ToI(row[3]),
					UpdatedAt: X.ToI(row[4]),
				})
			}
		}
	})

	return
}