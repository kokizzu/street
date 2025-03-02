package zImport

import (
	"time"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/model"
	"street/model/mProperty"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/wcProperty"
	"street/model/zCrud"
)

func FixCreatedUpdatedAt(propOltp *Tt.Adapter) {
	// data was imported with unix millis, it should be unix
	start := time.Now()

	fixPropertyTable(propOltp)
	fixPropertyHistoryTable(propOltp)

	L.TimeTrack(start, `FixCreatedUpdatedAt`)
}

func fixPropertyTable(propOltp *Tt.Adapter) {
	defer subTaskPrint(`fixPropertyTable: fix created_at/updated_at`)()

	prop := rqProperty.NewProperty(propOltp)

	const wrongImportAt = 1678_000_000_000

	in := zCrud.PagerIn{
		Page:    1,
		PerPage: 1000,
		Filters: map[string][]string{
			mProperty.CreatedAt: {`>` + I.ToS(wrongImportAt)}, // around 2 months ago
		},
		Order: []string{mProperty.Id},
	}

	meta := zCrud.Meta{
		Fields: []zCrud.Field{
			{Name: mProperty.Id},
			{Name: mProperty.CreatedAt},
			{Name: mProperty.UpdatedAt},
		},
	}

	out := &zCrud.PagerOut{}

	rows := prop.FindByPagination(&meta, &in, out)

	stat := &model.ImporterStat{Total: len(rows)}
	defer stat.Print(`last`)

	for len(rows) > 0 {
		stat.Total += len(rows)
		for _, row := range rows {
			createdAt := X.ToI(row[1])
			updatedAt := X.ToI(row[2])
			mutator := wcProperty.NewPropertyMutator(propOltp)
			mutator.Id = X.ToU(row[0])
			if createdAt > wrongImportAt {
				mutator.SetCreatedAt(createdAt / 1000)
			}
			if updatedAt > wrongImportAt {
				mutator.SetUpdatedAt(updatedAt / 1000)
			} else if updatedAt == 0 {
				mutator.SetUpdatedAt(createdAt / 1000)
			}
			if mutator.HaveMutation() {
				stat.Ok(mutator.DoUpdateById())
			}
			stat.Print()
		}
		rows = prop.FindByPagination(&meta, &in, out)
	}
}

func fixPropertyHistoryTable(propOltp *Tt.Adapter) {
	defer subTaskPrint(`fixPropertyHistoryTable: fix created_at/updated_at`)()

	ph := rqProperty.NewPropertyHistory(propOltp)

	const wrongImportAt = 1678_000_000_000

	in := zCrud.PagerIn{
		Page:    1,
		PerPage: 1000,
		Filters: map[string][]string{
			mProperty.CreatedAt: {`>` + I.ToS(wrongImportAt)}, // around 2 months ago
		},
		Order: []string{mProperty.Id},
	}

	meta := zCrud.Meta{
		Fields: []zCrud.Field{
			{Name: mProperty.Id},
			{Name: mProperty.CreatedAt},
			{Name: mProperty.UpdatedAt},
		},
	}

	out := &zCrud.PagerOut{}

	rows := ph.FindByPagination(&meta, &in, out)

	stat := &model.ImporterStat{Total: len(rows)}
	defer stat.Print(`last`)

	for len(rows) > 0 {
		stat.Total += len(rows)
		for _, row := range rows {
			createdAt := X.ToI(row[1])
			updatedAt := X.ToI(row[2])
			mutator := wcProperty.NewPropertyHistoryMutator(propOltp)
			mutator.Id = X.ToU(row[0])
			if createdAt > wrongImportAt {
				mutator.SetCreatedAt(createdAt / 1000)
			}
			if updatedAt > wrongImportAt {
				mutator.SetUpdatedAt(updatedAt / 1000)
			}
			if mutator.HaveMutation() {
				stat.Ok(mutator.DoUpdateById())
			}
			stat.Print()
		}
		rows = ph.FindByPagination(&meta, &in, out)
	}
}
