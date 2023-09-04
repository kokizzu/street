package conf_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"street/conf"
)

func TestTaiwanDateToInt(t *testing.T) {
	tests := []struct {
		taiwan   string
		yyyymmdd int
	}{
		{
			taiwan:   "1070101",
			yyyymmdd: 20180101,
		},
		{
			taiwan:   "1111101",
			yyyymmdd: 20221101,
		},
		{
			taiwan:   "1120101",
			yyyymmdd: 20230101,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tt.taiwan, tt.yyyymmdd), func(t *testing.T) {
			if got := conf.TaiwanDateToInt(tt.taiwan); got != tt.yyyymmdd {
				t.Errorf("TaiwanDateToInt() = %v, yyyymmdd %v", got, tt.yyyymmdd)
			}
		})
	}
}

func TestDistanceKm(t *testing.T) {
	tests := []struct {
		lat1, lon1, lat2, lon2 float64
		distanceKm             float64
	}{
		{
			lat1:       25.0330,
			lon1:       121.5654,
			lat2:       25.0329,
			lon2:       121.5655,
			distanceKm: 0.015005080171989555,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f,%f -> %f,%f = %f", tt.lat1, tt.lon1, tt.lat2, tt.lon2, tt.distanceKm), func(t *testing.T) {
			got := conf.DistanceKm(tt.lat1, tt.lon1, tt.lat2, tt.lon2)
			require.EqualValuesf(t, got, tt.distanceKm, "DistanceKm() = %v, distanceKm %v", got, tt.distanceKm)
		})
	}
}
