package conf

import (
	"fmt"
	"net/http"

	"github.com/valyala/tsvreader"
)

type CountryData struct {
	CountryName string `json:"country"`
}

var CountriesData []CountryData

func GoogleSheetCountryDataToJson(docId string, gId int) error {
	//DocID:	1TmAjrclFHUwDA1487ifQjX4FzYt9y7eJ0gwyxtwZMJU
	//GID:		522117981
	url := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/export?format=tsv&gid=%d", docId, gId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	tsv := tsvreader.New(resp.Body)

	for tsv.Next() {
		countryName := tsv.String()
		if countryName == `` || countryName == `country_name` {
			for tsv.HasCols() {
				_ = tsv.String()
			}
			continue
		}
		CountriesData = append(CountriesData, CountryData{
			CountryName: countryName,
		})
		for tsv.HasCols() {
			_ = tsv.String()
		}
	}
	return nil
}
