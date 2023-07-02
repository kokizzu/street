package zUpgrade

import (
	"time"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/model/mAuth"
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

	prop := rqProperty.NewProperty(propOltp)

	const wrongImportAt = 1688287266000

	in := zCrud.PagerIn{
		Page:    1,
		PerPage: 1000,
		Filters: map[string][]string{
			mAuth.CreatedAt: {`>` + I.ToS(wrongImportAt)}, // around 2 months ago
		},
		Order: []string{mAuth.Id},
	}

	meta := zCrud.Meta{
		Fields: []zCrud.Field{
			{Name: mAuth.Id},
			{Name: mAuth.CreatedAt},
			{Name: mAuth.UpdatedAt},
		},
	}

	out := &zCrud.PagerOut{}

	rows := prop.FindByPagination(&meta, &in, out)
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
		}
		if mutator.HaveMutation() {
			mutator.DoUpdateById()
		}
	}
}

func fixPropertyHistoryTable(propOltp *Tt.Adapter) {
	ph := rqProperty.NewPropertyHistory(propOltp)

	const wrongImportAt = 1688289858000

	in := zCrud.PagerIn{
		Page:    1,
		PerPage: 1000,
		Filters: map[string][]string{
			mAuth.CreatedAt: {`>` + I.ToS(wrongImportAt)}, // around 2 months ago
		},
		Order: []string{mAuth.Id},
	}

	meta := zCrud.Meta{
		Fields: []zCrud.Field{
			{Name: mAuth.Id},
			{Name: mAuth.CreatedAt},
			{Name: mAuth.UpdatedAt},
		},
	}

	out := &zCrud.PagerOut{}

	rows := ph.FindByPagination(&meta, &in, out)

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
			mutator.DoUpdateById()
		}
	}
}
