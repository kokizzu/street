package saProperty

import (
	"street/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
)

type ScannedAreasToRender struct {
	TimePeriod string `json:"time_period"`
	Views int64 `json:"views"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City string `json:"city"`
	State string `json:"state"`
}

func (s *ScannedAreas) FindMostScannedAreas() (res []ScannedAreasToRender) {
	const comment = `-- ScannedAreas) FindMostScannedAreas`
	query := comment + `
SELECT
  'Daily' AS time_period,
	count(1) AS views,
  any(city) AS city,
	any(state) AS state,
	any(latitude) AS lat,
	any(longitude) AS lng
FROM `+s.SqlTableName()+`
WHERE createdAt > (now() - toIntervalDay(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Weekly' AS time_period,
	count(1) AS views,
  any(city) AS city,
	any(state) AS state,
	any(latitude) AS lat,
	any(longitude) AS lng
FROM `+s.SqlTableName()+`
WHERE createdAt > (now() - toIntervalWeek(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Monthly' AS time_period,
	count(1) AS views,
  any(city) AS city,
	any(state) AS state,
	any(latitude) AS lat,
	any(longitude) AS lng
FROM `+s.SqlTableName()+`
WHERE createdAt > (now() - toIntervalMonth(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Quarterly' AS time_period,
	count(1) AS views,
  any(city) AS city,
	any(state) AS state,
	any(latitude) AS lat,
	any(longitude) AS lng
FROM `+s.SqlTableName()+`
WHERE createdAt > (now() - toIntervalQuarter(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Annually' AS time_period,
	count(1) AS views,
  any(city) AS city,
	any(state) AS state,
	any(latitude) AS lat,
	any(longitude) AS lng
FROM `+s.SqlTableName()+`
WHERE createdAt > (now() - toIntervalYear(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1
`
	L.Print(query)
	rows, err := s.Adapter.Query(query)
	if err != nil {
		L.IsError(err, `failed to get most scanned areas: `+query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			timePeriod string
			views int64
			city string
			state string
			lat float64
			lng float64
		)

		rows.Scan(&timePeriod, &views, &city, &state, &lat, &lng)

		res = append(res, ScannedAreasToRender{
			TimePeriod: timePeriod,
			Views: views,
			City: city,
			State: state,
			Latitude: lat,
			Longitude: lng,
		})
	}

	return
}

type ScannedPropertiesToRender struct {
	TimePeriod string `json:"time_period"`
	Views int64 `json:"views"`
	Price string `json:"price"`
	SizeSqft float64 `json:"total_sqft"`
	City string `json:"city"`
	State string `json:"state"`
	Address string `json:"address"`
}

func (sp *ScannedProperties) FindMostScannedProperties(ttConn *Tt.Adapter) (res []ScannedPropertiesToRender) {
	const comment = `-- ScannedProperties) FindMostScannedProperties`
	query := comment + `
SELECT
  'Daily' AS time_period,
	count(1) AS views,
  any(countryCode) AS countryCode,
	any(propertyId) AS propertyId
FROM `+sp.SqlTableName()+`
WHERE createdAt > (now() - toIntervalDay(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Weekly' AS time_period,
	count(1) AS views,
  any(countryCode) AS countryCode,
	any(propertyId) AS propertyId
FROM `+sp.SqlTableName()+`
WHERE createdAt > (now() - toIntervalWeek(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Monthly' AS time_period,
	count(1) AS views,
  any(countryCode) AS countryCode,
	any(propertyId) AS propertyId
FROM `+sp.SqlTableName()+`
WHERE createdAt > (now() - toIntervalMonth(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Quarterly' AS time_period,
	count(1) AS views,
  any(countryCode) AS countryCode,
	any(propertyId) AS propertyId
FROM `+sp.SqlTableName()+`
WHERE createdAt > (now() - toIntervalQuarter(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1

UNION ALL

SELECT
  'Annually' AS time_period,
	count(1) AS views,
  any(countryCode) AS countryCode,
	any(propertyId) AS propertyId
FROM `+sp.SqlTableName()+`
WHERE createdAt > (now() - toIntervalYear(1))
GROUP BY 1
ORDER BY views DESC
LIMIT 1
`
	L.Print(query)
	rows, err := sp.Adapter.Query(query)
	if err != nil {
		L.IsError(err, `failed to get most scanned areas: `+query)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			timePeriod string
			views int64
			countryCode string
			propertyId uint64
		)
		rows.Scan(&timePeriod, &views, &countryCode, &propertyId)

		var (
			price string
			sizeSqFt float64
			city string
			state string
			address string
		)

		switch countryCode {
		case `US`:
			propUS := rqProperty.NewPropertyUS(ttConn)
			propUS.Id = propertyId
			if !propUS.FindById() {
				L.Print(`property US with id `, propertyId, ` not found`)
			}

			price = propUS.LastPrice
			sizeSqFt = propUS.TotalSqft
			city = propUS.City
			state = propUS.State
			address = propUS.Address
			if address == `` {
				address = propUS.FormattedAddress
			}
		default:
			prop := rqProperty.NewProperty(ttConn)
			prop.Id = propertyId
			if !prop.FindById() {
				L.Print(`property with id `, propertyId, ` not found`)
			}

			price = prop.LastPrice
			sizeSqFt = prop.TotalSqft
			city = prop.City
			state = prop.State
			address = prop.Address
			if address == `` {
				address = prop.FormattedAddress
			}
		}

		res = append(res, ScannedPropertiesToRender{
			TimePeriod: timePeriod,
			Views: views,
			Price: price,
			SizeSqft: sizeSqFt,
			City: city,
			State: state,
			Address: address,
		})
	}

	return
}