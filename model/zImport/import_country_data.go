package zImport

import (
	"fmt"
	"net/http"

	"github.com/valyala/tsvreader"
)

type CountryData struct {
	CountryName string `json:"country"`
}

func GoogleSheetCountryDataToJson(docId string, gId int) ([]CountryData, error) {
	// DocID:	1TmAjrclFHUwDA1487ifQjX4FzYt9y7eJ0gwyxtwZMJU
	// GID:		522117981
	var c CountryData
	var CountriesData []CountryData
	url := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/export?format=tsv&gid=%d", docId, gId)
	resp, _ := http.Get(url)

	tsv := tsvreader.New(resp.Body)

	for tsv.Next() {
		countryName := tsv.String()
		if countryName == `` || countryName == `country_name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		c.CountryName = countryName
		CountriesData = append(CountriesData, c)
		for tsv.HasCols() {
			_ = tsv.String()
		}
	}
	return CountriesData, nil
}
