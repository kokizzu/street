package conf

import (
	"os"
	"strings"

	"github.com/valyala/tsvreader"
)

type (
	coordinate struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}
	currency struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
	CountryData struct {
		CountryName     string     `json:"country"`
		CountryISO2     string     `json:"iso_2"`
		CountryISO3     string     `json:"iso_3"`
		CountryCode     string     `json:"country_code"`
		Region          string     `json:"region"`
		RegionCode      string     `json:"region_code"`
		UnitMeasurement string     `json:"unit_measurement"`
		Coordinate      coordinate `json:"coordinate"`
		Currency        currency   `json:"currency"`
	}
)

var CountriesData []CountryData

func GetCountryData(file *os.File) {
	tsv := tsvreader.New(file)
	for tsv.Next() {
		country_name := tsv.String()
		if country_name == `` || country_name == `country_name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		iso_2 := tsv.String()
		if iso_2 == `` || iso_2 == `iso_2` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		iso_3 := tsv.String()
		if iso_3 == `` || iso_3 == `iso_3` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		country_code := tsv.String()
		if country_code == `` || country_code == `country_code` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		region := tsv.String()
		if region == `` || region == `region` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		region_code := tsv.String()
		if region_code == `` || region_code == `region_code` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		unit_measurement := tsv.String()
		if unit_measurement == `` || unit_measurement == `unit_measurement` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		coord_lat := tsv.String()
		if coord_lat == `` || coord_lat == `coordinate.latitude` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		coord_lng := tsv.String()
		if coord_lng == `` || coord_lng == `coordinate.longitude` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		currency_name := tsv.String()
		if currency_name == `` || currency_name == `currency.name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		currency_code := tsv.String()
		if currency_code == `` || currency_code == `currency.code` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		CountriesData = append(CountriesData, CountryData{
			CountryName:     strings.TrimSpace(country_name),
			CountryISO2:     strings.TrimSpace(iso_2),
			CountryISO3:     strings.TrimSpace(iso_3),
			CountryCode:     strings.TrimSpace(country_code),
			Region:          strings.TrimSpace(region),
			RegionCode:      strings.TrimSpace(region_code),
			UnitMeasurement: strings.TrimSpace(unit_measurement),
			Coordinate: coordinate{
				Lat: strings.TrimSpace(coord_lat),
				Lng: strings.TrimSpace(coord_lng),
			},
			Currency: currency{
				Name: currency_name,
				Code: country_code,
			},
		})

		for tsv.HasCols() {
			_ = tsv.String()
		}
	}
	return
}
