package zImport

import (
	"os"
	"path/filepath"
	"street/model/mProperty"
	"street/model/mProperty/saProperty"
	"time"

	chBuffer "github.com/kokizzu/ch-timed-buffer"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/valyala/tsvreader"
)

func ImportGeolocationDatabase(conn *Ch.Adapter, resourcePath string) {
	defer subTaskPrint(`ImportGeolocationDatabase: import geolocation data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()

	var geolocation []saProperty.Geolocation

	tsv := tsvreader.New(file)
	for tsv.Next() {
		id := tsv.String()
		city := tsv.String()
		stateId := tsv.String()
		stateCode := tsv.String()
		stateName := tsv.String()
		countryId := tsv.String()
		countryCode := tsv.String()
		countryName := tsv.String()
		latitude := tsv.String()
		longitude := tsv.String()
		wikiDataId := tsv.String()

		geolocation = append(geolocation, saProperty.Geolocation{
			Id:          S.ToU(id),
			City:        city,
			StateId:     S.ToU(stateId),
			StateCode:   stateCode,
			StateName:   stateName,
			CountryId:   S.ToU(countryId),
			CountryCode: countryCode,
			CountryName: countryName,
			Latitude:    S.ToF(latitude),
			Longitude:   S.ToF(longitude),
			WikiDataId:  wikiDataId,
		})
	}

	if len(geolocation) > 0 {
		geolocation = geolocation[1:]
	} else {
		panic(`Geolocation data from ` + resourcePath + ` is empty`)
	}

	stat := &ImporterStat{Total: len(geolocation)}
	defer stat.Print(`last`)

	timedBuffer := chBuffer.NewTimedBuffer(conn.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableGeolocation])

	for _, g := range geolocation {
		stat.Print()

		stat.Ok(timedBuffer.Insert([]any{
			g.Id,
			g.City,
			g.StateId,
			g.StateCode,
			g.StateName,
			g.CountryId,
			g.CountryCode,
			g.CountryName,
			g.Latitude,
			g.Longitude,
			g.WikiDataId,
		}))
	}
}
