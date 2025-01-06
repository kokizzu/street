package rqBusiness

import (
	"street/model/mBusiness"
	"time"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/kpango/fastime"
	"github.com/vburenin/nsync"
)

var (
	CACHED_ORDERS_ANNUALLY mBusiness.Cache = mBusiness.Cache{
		CacheName:     `ordersAnnually`,
		CacheUnixTime: 0,
		CacheData:     []any{},
	}
	CACHED_REVENUES_MONTHLY mBusiness.Cache = mBusiness.Cache{
		CacheName:     `ordersAnnually`,
		CacheUnixTime: 0,
		CacheData:     []any{},
	}
)

func (s *Sales) FindRealtorRevenuesMonthlyByRealtorId(yearMonth string) (revenues []*mBusiness.Revenue) {
	const comment = `-- Sales) FindRealtorRevenuesMonthlyByRealtorId`

	var startDate string // YYYY-MM-DD
	var endDate string   // YYYY-MM-DD

	year, month, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		startDate = startOfYearMonth(now.Year(), now.Month())
		endDate = endOfYearMonth(now.Year(), now.Month())
	} else {
		startDate = startOfYearMonth(year, month)
		endDate = endOfYearMonth(year, month)
	}

	whereAndSql := ` WHERE ` + s.SqlRealtorId() + ` = ` + I.UToS(s.RealtorId) + `
		AND ` + s.SqlSalesDate() + ` >= ` + S.Z(startDate) + ` AND ` + s.SqlSalesDate() + ` <= ` + S.Z(endDate)

	query := comment + `
SELECT ` + s.SqlPrice() + `, ` + s.SqlSalesDate() + `
FROM ` + s.SqlTableName() + whereAndSql

	L.Print(`QUERY :`, query)
	s.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 2 {
			price := X.ToI(row[0])
			salesDate := X.ToS(row[1])

			var isExist bool = false

			for _, r := range revenues {
				if r.SalesDate == salesDate {
					r.PropertyBought += 1
					r.Revenue += price

					isExist = true
				}
			}

			if !isExist {
				revenues = append(revenues, &mBusiness.Revenue{
					Revenue:        price,
					PropertyBought: 1,
					SalesDate:      salesDate,
				})
			}
		}
	})

	return
}

var revenuesMonthlyLock = nsync.NewNamedMutex()

func (s *Sales) FindRevenuesMonthly(yearMonth string) (revenues []*mBusiness.Revenue) {
	const comment = `-- Sales) FindRevenuesMonthly`

	if !CACHED_REVENUES_MONTHLY.IsExpired() {
		L.Print(`From Cache: ` + comment)
		revenues = CACHED_REVENUES_MONTHLY.CacheData.([]*mBusiness.Revenue)
		return
	}

	if !revenuesMonthlyLock.TryLock(CACHED_REVENUES_MONTHLY.CacheName) {
		revenues = CACHED_REVENUES_MONTHLY.CacheData.([]*mBusiness.Revenue)
		return
	}
	defer revenuesMonthlyLock.Unlock(CACHED_REVENUES_MONTHLY.CacheName)

	var startDate string // YYYY-MM-DD
	var endDate string   // YYYY-MM-DD

	year, month, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		startDate = startOfYearMonth(now.Year(), now.Month())
		endDate = endOfYearMonth(now.Year(), now.Month())
	} else {
		startDate = startOfYearMonth(year, month)
		endDate = endOfYearMonth(year, month)
	}

	whereAndSql := ` WHERE ` + s.SqlSalesDate() + ` >= ` + S.Z(startDate) + ` AND ` + s.SqlSalesDate() + ` <= ` + S.Z(endDate)

	query := comment + `
SELECT ` + s.SqlPrice() + `, ` + s.SqlSalesDate() + `
FROM ` + s.SqlTableName() + whereAndSql

	L.Print(`QUERY :`, query)
	s.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 2 {
			price := X.ToI(row[0])
			salesDate := X.ToS(row[1])

			var isExist bool = false

			for _, r := range revenues {
				if r.SalesDate == salesDate {
					r.PropertyBought += 1
					r.Revenue += price

					isExist = true
				}
			}

			if !isExist {
				revenues = append(revenues, &mBusiness.Revenue{
					Revenue:        price,
					PropertyBought: 1,
					SalesDate:      salesDate,
				})
			}
		}
	})

	CACHED_REVENUES_MONTHLY.CacheData = revenues
	CACHED_REVENUES_MONTHLY.CacheUnixTime = fastime.UnixNow()

	return
}

var ordersAnnuallyLock = nsync.NewNamedMutex()

func (s *Sales) FindOrdersAnnually() (orders []*mBusiness.Order) {
	const comment = `-- Sales) FindRevenuesMonthly`

	if !CACHED_ORDERS_ANNUALLY.IsExpired() {
		L.Print(`From Cache: ` + comment)
		orders = CACHED_ORDERS_ANNUALLY.CacheData.([]*mBusiness.Order)
		return
	}

	if !ordersAnnuallyLock.TryLock(CACHED_ORDERS_ANNUALLY.CacheName) {
		orders = CACHED_ORDERS_ANNUALLY.CacheData.([]*mBusiness.Order)
		return
	}
	defer ordersAnnuallyLock.Unlock(CACHED_ORDERS_ANNUALLY.CacheName)

	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0)
	strOneYearAgo := oneYearAgo.Format("2006-01-02")

	whereAndSql := ` WHERE ` + s.SqlSalesDate() + ` >= ` + S.Z(strOneYearAgo)

	query := comment + `
SELECT ` + s.SqlPrice() + `, ` + s.SqlSalesDate() + `
FROM ` + s.SqlTableName() + whereAndSql

	//L.Print(query)
	s.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 2 {
			price := X.ToI(row[0])
			salesDate := X.ToS(row[1])

			var isExist bool = false

			for _, r := range orders {
				if r.SalesDate == salesDate {
					r.TotalTransaction += 1
					r.Revenue += price

					isExist = true
				}
			}

			if !isExist {
				orders = append(orders, &mBusiness.Order{
					Revenue:          price,
					TotalTransaction: 1,
					SalesDate:        salesDate,
				})
			}
		}
	})

	CACHED_ORDERS_ANNUALLY.CacheData = orders
	CACHED_ORDERS_ANNUALLY.CacheUnixTime = fastime.UnixNow()
	return
}

func endOfYearMonth(year int, month time.Month) string { // returns: YYYY-MM-DD
	now := time.Now()
	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())
	return firstOfNextMonth.AddDate(0, 0, -1).Format("2006-01-02")
}

func startOfYearMonth(year int, month time.Month) string { // returns: YYYY-MM-DD
	now := time.Now()
	firstOfNextMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	return firstOfNextMonth.Format("2006-01-02")
}

func isValidMonthYear(in string) (year int, month time.Month, isValid bool) {
	yearMonthArr := S.Split(in, `-`)

	if len(yearMonthArr) != 2 {
		return
	}

	if _, err := time.Parse("2006-01", in); err != nil {
		return
	}

	year = int(S.ToI(yearMonthArr[0]))
	monthInt := S.ToI(yearMonthArr[1])

	if year < 1900 || year > 2100 {
		return
	}

	if monthInt < 1 || monthInt > 12 {
		return
	}

	month = time.Month(monthInt)

	isValid = true
	return
}

func (s *Sales) FindBuyerMonthly(yearMonth string) (buyers []mBusiness.Buyer) {
	const comment = `-- Sales) FindBuyerMonthly`

	var startDate string // YYYY-MM-DD
	var endDate string   // YYYY-MM-DD

	year, month, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		startDate = startOfYearMonth(now.Year(), now.Month())
		endDate = endOfYearMonth(now.Year(), now.Month())
	} else {
		startDate = startOfYearMonth(year, month)
		endDate = endOfYearMonth(year, month)
	}

	whereAndSql := ` WHERE ` + s.SqlSalesDate() + ` >= ` + S.Z(startDate) + ` AND ` + s.SqlSalesDate() + ` <= ` + S.Z(endDate)

	query := comment + `
SELECT ` + s.SqlPropertyId() + `, ` + s.SqlBuyerEmail() + `, ` + s.SqlEmailNotFound() + `, ` + s.SqlPrice() + `, ` + s.SqlSalesDate() + `
FROM ` + s.SqlTableName() + whereAndSql

	L.Print(`QUERY :`, query)
	s.Adapter.QuerySql(query, func(row []any) {
		email := X.ToS(row[1])
		if email == `` {
			email = X.ToS(row[2])
		}

		buyers = append(buyers, mBusiness.Buyer{
			PropertyId:    X.ToU(row[0]),
			BuyerEmail:    email,
			PropertyPrice: X.ToI(row[3]),
			SalesDate:     X.ToS(row[4]),
		})
	})

	return
}
