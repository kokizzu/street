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
			expectOut:      autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 1, Total: 1}),
		},
		{
			name: `orderInvalid`,
			in: PagerIn{
				Order: []string{`+`, `-`, `noDiraction`, `+notExists`},
			},
			countResult: 1,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
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

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
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

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql: autogold.Expect(`
ORDER BY "id" DESC`),
			expectOut: autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 1, Total: 1}),
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
			expectOut: autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 5, Total: 44}),
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
			name: `filterNeqEmpty`,
			in: PagerIn{
				Filters: map[string][]string{
					`name`: {`<>`},
				},
			},
			countResult: 100,
			fieldToType: map[string]Tt.DataType{
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("name" NOT IN (''))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 10, Total: 100,
				Filters: map[string][]string{
					"name": {"<>"},
				},
			}),
		},
		{
			name: `filterInvalidField`,
			in: PagerIn{
				Filters: map[string][]string{
					`somethingNotHere`: {`>1`},
				},
				PerPage: 40,
			},
			countResult: 88,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 40"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 1, PerPage: 40, Pages: 3, Total: 88}),
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
			name: `filterLtGe`,
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
			name: `filterGt`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`>1`},
					`name`: {`>=abc`},
				},
				PerPage: 40,
			},
			countResult: 88,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 40"),
			whereAndSql: autogold.Expect(`
WHERE ("id">1)
AND ("name">='abc')`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 40, Pages: 3, Total: 88,
				Filters: map[string][]string{
					"id":   {">1"},
					"name": {">=abc"},
				},
			}),
		},
		{
			name: `filterLt`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`<=1`},
					`name`: {`<abc`},
				},
				PerPage: 2,
			},
			countResult: 3,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 2"),
			whereAndSql: autogold.Expect(`
WHERE ("id"<=1)
AND ("name"<'abc')`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{Page: 1, PerPage: 2, Pages: 2, Total: 3, Filters: map[string][]string{
				"id":   {"<=1"},
				"name": {"<abc"},
			}}),
		},
		{
			name: `filterLtGt`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`>=1`},
					`name`: {`<abc`},
				},
			},
			countResult: 95,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("id">=1)
AND ("name"<'abc')`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 10, Total: 95,
				Filters: map[string][]string{
					"id":   {">=1"},
					"name": {"<abc"},
				},
			}),
		},
		{
			name: `filterLtGtMulti`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`>1`, `<=23`},
					`name`: {`abc`, `<def`, `>=ghij`},
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
AND ("name"<'def' OR "name">='ghij' OR "name" IN ('abc'))`),
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
						">=ghij",
						"<def",
					},
				},
			}),
		},
		{
			name: `filterLike`,
			in: PagerIn{
				Filters: map[string][]string{
					`name`: {`<>a*bc`, `*def*`},
				},
			},
			countResult: 3,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("name" NOT LIKE 'a%bc' OR "name" LIKE '%def%')`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 3,
				Filters: map[string][]string{
					"name": {
						"<>a*bc",
						"*def*",
					},
				},
			}),
		},
		{
			name: `filterMultiOp`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`4`, `9`, `<>-5`, `<>44`, `<=1`, `>23`},
					`name`: {`abc`, ``, `>def`, `<=ghij`, `xyz`, `<>a`, `<>`, `foo`},
				},
			},
			countResult: 10,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("id"<=1 OR "id">23 OR "id" IN (4,9) OR "id" NOT IN (-5,44))
AND (("name">'def' AND "name"<='ghij') OR "name" IN ('abc','','xyz','foo') OR "name" NOT IN ('a',''))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 10,
				Filters: map[string][]string{
					"id": {
						"4",
						"9",
						"<>-5",
						"<>44",
						">23",
						"<=1",
					},
					"name": {
						"abc",
						"",
						"xyz",
						"<>a",
						"<>",
						"foo",
						">def",
						"<=ghij",
					},
				},
			}),
		},
		{
			name: `filterSqlInjection`,
			in: PagerIn{
				Filters: map[string][]string{
					`id`:   {`' OR 1=1; -- `},
					`name`: {`'; DROP TABLE users; -- `, `') OR 1=1; -- `},
				},
			},
			countResult: 2,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql: autogold.Expect(`
WHERE ("name" IN ('&apos;; DROP TABLE users; --','&apos;) OR 1=1; --'))`),
			orderBySql: autogold.Expect(""),
			expectOut: autogold.Expect(PagerOut{
				Page: 1, PerPage: 10, Pages: 1, Total: 2,
				Filters: map[string][]string{
					"id": {},
					"name": {
						"'; DROP TABLE users; -- ",
						"') OR 1=1; -- ",
					},
				},
			}),
		},
		{
			name: `offset`,
			in: PagerIn{
				Page:    2,
				PerPage: 5,
			},
			countResult: 11,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 5 OFFSET 5"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 2, PerPage: 5, Pages: 3, Total: 11}),
		},
		{
			name: `offsetNegative`,
			in: PagerIn{
				Page:    -55,
				PerPage: 5,
			},
			countResult: 11,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 5"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 1, PerPage: 5, Pages: 3, Total: 11}),
		},
		{
			name: `offsetOverflow`,
			in: PagerIn{
				Page:    555,
				PerPage: 5,
			},
			countResult: 11,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 5 OFFSET 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 3, PerPage: 5, Pages: 3, Total: 11}),
		},
		{
			name: `limitNegative`,
			in: PagerIn{
				Page:    -1,
				PerPage: -2,
			},
			countResult: 11,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 10"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 1, PerPage: 10, Pages: 2, Total: 11}),
		},
		{
			name: `limitOverflow`,
			in: PagerIn{
				Page:    300,
				PerPage: 2000,
			},
			countResult: 2100,
			fieldToType: map[string]Tt.DataType{
				`id`:   Tt.Integer,
				`name`: Tt.String,
			},

			limitOffsetSql: autogold.Expect("\nLIMIT 1000 OFFSET 2000"),
			whereAndSql:    autogold.Expect(""),
			orderBySql:     autogold.Expect(""),
			expectOut:      autogold.Expect(PagerOut{Page: 3, PerPage: 1000, Pages: 3, Total: 2100}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := PagerOut{}

			out.CalculatePages(tc.in.Page, tc.in.PerPage, tc.countResult)

			limitOffsetSql := out.LimitOffsetSql()
			tc.limitOffsetSql.Equal(t, limitOffsetSql)

			whereAndSql := out.WhereAndSqlTt(tc.in.Filters, tc.fieldToType)
			tc.whereAndSql.Equal(t, whereAndSql)

			orderBySql := out.OrderBySqlTt(tc.in.Order, tc.fieldToType)
			tc.orderBySql.Equal(t, orderBySql)

			tc.expectOut.Equal(t, out)

		})
	}
}
