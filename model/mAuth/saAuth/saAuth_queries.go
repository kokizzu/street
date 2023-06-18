package saAuth

import (
	"database/sql"
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
