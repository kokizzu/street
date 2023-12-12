package conf

import (
	"fmt"
	"io"
	"math"

	"github.com/kokizzu/gotro/S"
)

func TaiwanDateToInt(date string) int {
	if len(date) < 7 {
		return 0 + 1*100 + 1
	}

	yyymmdd := S.ToInt(date)
	yyyymmdd := 1911_00_00 + yyymmdd

	return yyyymmdd
}

func TaiwanDateToStr(date string) string {
	var yy, mm, dd int
	n, err := fmt.Sscanf(date, "%d-%d-%d", &yy, &mm, &dd)
	if err == nil {
		return fmt.Sprintf("%d-%02d-%02d", yy+1911, mm, dd)
	}
	if err == io.ErrUnexpectedEOF && n == 2 {
		return fmt.Sprintf("%d-%02d-01", yy+1911, mm)
	}
	return ""
}

func DistanceKm(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6371 // km
	const p = math.Pi / 180

	a := 0.5 - math.Cos((lat2-lat1)*p)/2 + math.Cos(lat1*p)*math.Cos(lat2*p)*(1-math.Cos((lon2-lon1)*p))/2

	return 2 * r * math.Asin(math.Sqrt(a))
}
