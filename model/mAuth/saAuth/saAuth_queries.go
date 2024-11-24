package saAuth

import (
	"database/sql"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"street/model/mAuth"
	"street/model/mAuth/rqAuth"
	"street/model/zCrud"
)

func scanMap(rows *sql.Rows, err error) map[string]int {
	if err != nil {
		return map[string]int{
			err.Error(): 0,
		}
	}
	defer rows.Close()
	res := map[string]int{}
	for rows.Next() {
		var date string
		var count int
		if err := rows.Scan(&date, &count); err != nil {
			return map[string]int{
				err.Error(): 0,
			}
		}
		res[date] = count
	}
	return res
}

func (a ActionLogs) StatRequestsPerDate() (res map[string]int) {
	rows, err := a.Adapter.Query(`-- StatRequestsPerDate
SELECT DATE(createdAt), COUNT(1)
FROM actionLogs
GROUP BY 1
ORDER BY 1 DESC
LIMIT 90
`)
	return scanMap(rows, err)
}

func (a ActionLogs) StatUniqueUserPerDate() (res map[string]int) {
	rows, err := a.Adapter.Query(`-- StatUniqueUserPerDate
SELECT DATE(createdAt), COUNT(DISTINCT actorId)
FROM actionLogs
GROUP BY 1
ORDER BY 1 DESC
LIMIT 90
`)
	return scanMap(rows, err)
}

func (a ActionLogs) StatUniqueIpPerDate() map[string]int {
	rows, err := a.Adapter.Query(`-- StatUniqueIpPerDate
SELECT DATE(createdAt), COUNT(DISTINCT ipAddr4)
FROM actionLogs
GROUP BY 1
ORDER BY 1 DESC
LIMIT 90
`)
	return scanMap(rows, err)
}

func (a ActionLogs) StatPerActionsPerDate() map[string]map[string]int {
	rows, err := a.Adapter.Query(`-- StatPerActionsPerDate
SELECT action, DATE(createdAt), COUNT(1)
FROM actionLogs
WHERE DATE(createdAt) >= DATE_SUB(NOW(), INTERVAL 90 DAY)
GROUP BY 1, 2
ORDER BY 1, 2 DESC
`)
	if err != nil {
		return map[string]map[string]int{
			err.Error(): {},
		}
	}
	defer rows.Close()
	res := map[string]map[string]int{}
	for rows.Next() {
		var action, date string
		var count int
		if err := rows.Scan(&action, &date, &count); err != nil {
			return map[string]map[string]int{
				err.Error(): {},
			}
		}
		if _, ok := res[action]; !ok {
			res[action] = map[string]int{}
		}
		res[action][date] = count
	}
	return res
}

func (a ActionLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []ActionLogs) {
	const comment = `-- Property) FindByPagination`

	validFields := ActionLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + a.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := a.Adapter.QueryRow(queryCount)
	err := row.Err()
	if L.IsError(err, `FindByPagination.Adapter.QueryRow error: `+queryCount) {
		return
	}
	err = row.Scan(&out.Total)
	L.IsError(err, comment+`: error while scanning total`)
	out.CalculatePages(in.Page, in.PerPage, out.Total)

	orderBySql := out.OrderBySqlCh(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + a.SqlAllFields() + `
FROM ` + a.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := a.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = a.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (a ActionLogs) FindUserRegistered() (res []mAuth.UserRegisterStat) {
	query := `-- FindUserRegistered
	SELECT DISTINCT toDate(createdAt) dt,
		COUNT(*) OVER (PARTITION BY toDate(createdAt)) AS count
	FROM actionLogs
	WHERE action = 'guest/register'
		AND createdAt > (today() - INTERVAL 1 YEAR)
	ORDER BY dt
	`
	L.Print(query)
	rows, err := a.Adapter.Query(query)
	if err != nil {
		L.IsError(err, `failed to get user registered: `+query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			dt string
			count int64
		)

		rows.Scan(&dt, &count)

		res = append(res, mAuth.UserRegisterStat{
			Date: dt,
			Count: count,
		})
	}

	return
}

func (a ActionLogs) FindRealtorActivity() (res []mAuth.RealtorStat) {
	query := `-- FindRealtorActivity
	SELECT DISTINCT toDate(createdAt) dt,
		COUNT(*) OVER (PARTITION BY toDate(createdAt)) AS count
	FROM actionLogs
	WHERE action = 'realtor/upsertProperty'
		AND createdAt > (today() - INTERVAL 1 YEAR)
	ORDER BY dt
	`
	L.Print(query)
	rows, err := a.Adapter.Query(query)
	if err != nil {
		L.IsError(err, `failed to get user registered: `+query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			dt string
			count int64
		)

		rows.Scan(&dt, &count)

		res = append(res, mAuth.RealtorStat{
			Date: dt,
			TotalActivity: count,
		})
	}

	return
}

type MostLoggedInUser struct {
	TimePeriod string `json:"time_period"`
	Email string `json:"email"`
	FullName string `json:"full_name"`
	Total int64 `json:"total"`
}

func (a ActionLogs) FindMostLoggedInUsers(ttConn *Tt.Adapter) (res []MostLoggedInUser) {
	comment := `-- ActionLogs) FindMostLoggedInUsers`
	query := comment + `
SELECT
  'Daily' AS time_period,
  actorId,
  count(1) AS total
FROM actionLogs
WHERE (createdAt > (now() - toIntervalDay(1)))
	AND action = 'guest/login'
	AND actorId > 0
GROUP BY 1, 2
ORDER BY total DESC
LIMIT 1

UNION ALL

SELECT
  'Weekly' AS time_period,
  actorId,
  count(1) AS total
FROM actionLogs
WHERE (createdAt > (now() - toIntervalWeek(1)))
	AND action = 'guest/login'
	AND actorId > 0
GROUP BY 1, 2
ORDER BY total DESC
LIMIT 1

UNION ALL

SELECT
  'Monthly' AS time_period,
  actorId,
  count(1) AS total
FROM actionLogs
WHERE (createdAt > (now() - toIntervalMonth(1)))
	AND action = 'guest/login'
	AND actorId > 0
GROUP BY 1, 2
ORDER BY total DESC
LIMIT 1

UNION ALL

SELECT
  'Quarterly' AS time_period,
  actorId,
  count(1) AS total
FROM actionLogs
WHERE (createdAt > (now() - toIntervalQuarter(1)))
	AND action = 'guest/login'
	AND actorId > 0
GROUP BY 1, 2
ORDER BY total DESC
LIMIT 1

UNION ALL

SELECT
  'Annually' AS time_period,
  actorId,
  count(1) AS total
FROM actionLogs
WHERE (createdAt > (now() - toIntervalYear(1)))
	AND action = 'guest/login'
	AND actorId > 0
GROUP BY 1, 2
ORDER BY total DESC
LIMIT 1`

	L.Print(query)
	rows, err := a.Adapter.Query(query)
	if err != nil {
		L.IsError(err, `failed to get most logged in users: `+query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			timePeriod string
			userId uint64
			total int64
		)
		rows.Scan(&timePeriod, &userId, &total)

		user := rqAuth.NewUsers(ttConn)
		user.Id = userId
		if !user.FindById() {
			L.Print(`user with id `, userId, ` not found`)
		}

		res = append(res, MostLoggedInUser{
			TimePeriod: timePeriod,
			Email: user.Email,
			FullName: user.FullName,
			Total: total,
		})
	}

	return
}