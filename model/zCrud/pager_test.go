package zCrud

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/kokizzu/gotro/D/Tt"
)

//go:generate go test -count=1 -run TestPagination . -update

func TestPagination(t *testing.T) {
	testCases := []struct {
		name        string
		in          PagerIn
		fieldToType map[string]Tt.DataType
		countResult int

		expectOut      autogold.Value
		whereAndSql    autogold.Value
		orderBySql     autogold.Value
		limitOffsetSql autogold.Value
	}{
		{
			name:        `empty`,
			in:          PagerIn{},
			countResult: 1,
			fieldToType: map[string]Tt.DataType{},

			limitOffsetSql: autogold.Expect("LIMIT 10 OFFSET 0"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 1, Total: 1}),
		},
		{
			name: `orderAsc`,
			in: PagerIn{
				Order: []string{`+id`},
			},
			countResult: 1,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("LIMIT 10 OFFSET 0"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "id"`),
			expectOut: autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 1, Total: 1}),
		},
		{
			name: `orderDesc`,
			in: PagerIn{
				Order: []string{`-id`},
			},
			countResult: 1,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("LIMIT 10 OFFSET 0"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "id" DESC`),
			expectOut: autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 1, Total: 1}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := PagerOut{}

			limitOffsetSql := out.LimitOffsetSql(&tc.in)
			tc.limitOffsetSql.Equal(t, limitOffsetSql)

			whereAndSql, orderBySql := out.WhereOrderSql(tc.in.Filters, tc.in.Order, tc.fieldToType)
			tc.whereAndSql.Equal(t, whereAndSql)
			tc.orderBySql.Equal(t, orderBySql)

			out.CalculatePages(tc.countResult)
			tc.expectOut.Equal(t, out)

		})
	}
}
