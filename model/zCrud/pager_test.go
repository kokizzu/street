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

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 1,
				Filters: map[string][]string{},
			}),
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

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "id"`),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 1,
				Filters: map[string][]string{},
			}),
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

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "id" DESC`),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 1,
				Filters: map[string][]string{},
			}),
		},
		{
			name: `orderMulti`,
			in: PagerIn{
				Order: []string{`-name`, `+id`},
			},
			countResult: 44,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "name" DESC, "id"`),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 5, Total: 44,
				Filters: map[string][]string{},
			}),
		},
		{
			name: `filterEq`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`1`},
					`name`: {`23`},
				},
			},
			countResult: 100,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("id" IN (1))
AND ("name" IN ('23'))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 10, Total: 100,
				Filters: map[string][]string{
					"id":   {"1"},
					"name": {"23"},
				},
			}),
		},
		{
			name: `filterEqMulti`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`1`, `23`},
					`name`: {`abc`, `def`, `ghij`},
				},
			},
			countResult: 50,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("id" IN (1,23))
AND ("name" IN ('abc','def','ghij'))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 5, Total: 50,
				Filters: map[string][]string{
					"id": {
						"1",
						"23",
					},
					"name": {
						"abc",
						"def",
						"ghij",
					},
				},
			}),
		},
		{
			name: `filterNeq`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`<>1`, `<>23`},
					`name`: {`<>abc`, `<>`},
				},
				PerPage: 20,
			},
			countResult: 41,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 20"),
			whereAndSql: autogold.Expect(`
WHERE ("id" NOT IN (1,23))
AND ("name" NOT IN ('abc',''))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 20, Pages: 3, Total: 41,
				Filters: map[string][]string{
					"id": {
						"<>1",
						"<>23",
					},
					"name": {
						"<>abc",
						"<>",
					},
				},
			}),
		},
		{
			name: `filterLtGt`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`>1`},
					`name`: {`<=abc`},
				},
			},
			countResult: 31,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("id">1)
AND ("name"<='abc')`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 4, Total: 31,
				Filters: map[string][]string{
					"id":   {">1"},
					"name": {"<=abc"},
				},
			}),
		},
		{
			name: `filterLtGtMulti`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`>1`, `<=23`},
					`name`: {`abc`, `<=def`, `>ghij`},
				},
			},
			countResult: 39,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE (("id">1 AND "id"<=23))
AND ("name"<='def' OR "name">'ghij' OR "name" IN ('abc'))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 4, Total: 39,
				Filters: map[string][]string{
					"id": {
						">1",
						"<=23",
					},
					"name": {
						"abc",
						">ghij",
						"<=def",
					},
				},
			}),
		},
		{
			name: `filterMultiOp`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`4`, `9`, `<>-5`, `<>44`, `>1`, `<=23`},
					`name`: {`abc`, `<=def`, `>ghij`, `xyz`, `<>a`, `<>`, `foo`},
				},
			},
			countResult: 39,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE (("id">1 AND "id"<=23) OR "id" IN (4,9) OR "id" NOT IN (-5,44))
AND ("name"<='def' OR "name">'ghij' OR "name" IN ('abc','xyz','foo') OR "name" NOT IN ('a',''))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 4, Total: 39,
				Filters: map[string][]string{
					"id": {
						"4",
						"9",
						"<>-5",
						"<>44",
						">1",
						"<=23",
					},
					"name": {
						"abc",
						"xyz",
						"<>a",
						"<>",
						"foo",
						">ghij",
						"<=def",
					},
				},
			}),
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
