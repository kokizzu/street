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
		countryName := tsv.String()
		if countryName == `` || countryName == `country_name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		iso2 := tsv.String()
		if iso2 == `` || iso2 == `iso_2` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		iso3 := tsv.String()
		if iso3 == `` || iso3 == `iso_3` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		countryCode := tsv.String()
		if countryCode == `` || countryCode == `country_code` {
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
		regionCode := tsv.String()
		if regionCode == `` || regionCode == `region_code` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		unitMeasurement := tsv.String()
		if unitMeasurement == `` || unitMeasurement == `unit_measurement` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		coordLat := tsv.String()
		if coordLat == `` || coordLat == `coordinate.latitude` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		coordLng := tsv.String()
		if coordLng == `` || coordLng == `coordinate.longitude` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		currencyName := tsv.String()
		if currencyName == `` || currencyName == `currency.name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		currencyCode := tsv.String()
		if currencyCode == `` || currencyCode == `currency.code` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		CountriesData = append(CountriesData, CountryData{
			CountryName:     strings.TrimSpace(countryName),
			CountryISO2:     strings.TrimSpace(iso2),
			CountryISO3:     strings.TrimSpace(iso3),
			CountryCode:     strings.TrimSpace(countryCode),
			Region:          strings.TrimSpace(region),
			RegionCode:      strings.TrimSpace(regionCode),
			UnitMeasurement: strings.TrimSpace(unitMeasurement),
			Coordinate: coordinate{
				Lat: strings.TrimSpace(coordLat),
				Lng: strings.TrimSpace(coordLng),
			},
			Currency: currency{
				Name: currencyName,
				Code: countryCode,
			},
		})
		for tsv.HasCols() {
			_ = tsv.String()
		}
	}
	return
}
