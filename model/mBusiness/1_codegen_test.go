package mBusiness

import (
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
)

//go:generate go test -bench=BenchmarkGenerateOrm

func BenchmarkGenerateOrm(b *testing.B) {
	Tt.GenerateOrm(TarantoolTables)
	b.SkipNow()
}