package zCrud

import (
	"fmt"
	"strconv"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file pagination.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type pagination.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type pagination.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type pagination.go
//go:generate farify doublequote --file pagination.go

type PaginationIn struct {
	Page    int `json:"page" form:"page" query:"page" long:"page" msg:"page"`
	PerPage int `json:"perPage" form:"perPage" query:"perPage" long:"perPage" msg:"perPage"`

	// filter AND by column, if value is array then filter OR on that field
	Filters map[string][]string `json:"filters" form:"filters" query:"filters" long:"filters" msg:"filters"`

	// Order: [+col1, -col2] (+ is ascending, - is descending)
	Order []string `json:"order" form:"order" query:"order" long:"order" msg:"order"`
}

const minPerPage = 10
const maxPerPage = 1000

func (p *PaginationIn) Limit() int {
	if p.PerPage <= minPerPage {
		return minPerPage
	}
	if p.PerPage >= maxPerPage {
		return maxPerPage
	}
	return p.PerPage
}

func (p *PaginationIn) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.PerPage
}

type PaginationOut struct {
	Page    int `json:"page" form:"page" query:"page" long:"page" msg:"page"`
	PerPage int `json:"perPage" form:"perPage" query:"perPage" long:"perPage" msg:"perPage"`

	Pages int `json:"pages" form:"pages" query:"pages" long:"pages" msg:"pages"`
	Total int `json:"total" form:"total" query:"total" long:"total" msg:"total"`

	Filters map[string][]string `json:"filters" form:"filters" query:"filters" long:"filters" msg:"filters"`

	Order []string `json:"order" form:"order" query:"order" long:"order" msg:"order"`
}

func (p *PaginationOut) LimitOffsetSql(in *PaginationIn) string {
	offset := in.Offset()
	p.PerPage = in.Limit()
	p.Page = offset/p.PerPage + 1
	return fmt.Sprintf(`LIMIT %d OFFSET %d`, p.PerPage, offset)
}

func (p *PaginationOut) WhereOrderSql(filters map[string][]string, orders []string, fieldToType map[string]Tt.DataType) (whereAndSql, orderBySql string) {

	whereAnd := []string{}
	for field, value := range filters {
		if typ, ok := fieldToType[field]; ok {
			quotedValue := equalityQuoteValue(value, typ, S.QQ(field))
			if len(quotedValue) > 1 {
				whereOr := A.StrJoin(quotedValue, ` OR `)
				whereAnd = append(whereAnd, whereOr)
			} else if len(quotedValue) == 1 {
				whereAnd = append(whereAnd, quotedValue[0])
			}
		}
	}
	if len(whereAnd) > 0 {
		whereAndSql = "\n" + `WHERE (` + A.StrJoin(whereAnd, `) 
	AND (`) + `)`
	}

	orderBy := []string{}
	for _, dirField := range orders {
		if len(dirField) <= 2 {
			continue
		}
		dir := dirField[0]
		dirStr := ``
		if dir == '+' {
		} else if dir != '-' {
			dirStr = ` DESC`
		} else {
			continue
		}
		orderBy = append(orderBy, dirField[1:]+dirStr)
	}
	if len(orderBy) > 0 {
		orderBySql = "\n" + `ORDER BY ` + A.StrJoin(orderBy, `, `)
	}
	return whereAndSql, orderBySql
}

func equalityQuoteValue(values []string, expectTyp Tt.DataType, field string) (res []string) {
	// allow >, <, >=, <=, <>, *LIKE* and NOT LIKE, if multiple = or <> then will use IN or NOT IN
	switch expectTyp {
	case Tt.Unsigned, Tt.Integer, Tt.Double:
		equalValues := []string{}
		inequalValues := []string{}
		for _, str := range values {
			operator, rhs := splitOperatorValue(str)
			if _, err := strconv.ParseFloat(rhs, 64); err != nil {
				continue
			}
			if operator == `=` {
				equalValues = append(equalValues, rhs)
			} else if operator == `<>` {
				inequalValues = append(inequalValues, rhs)
			} else {
				res = append(res, field+operator+rhs)
			}
			// TODO: autodetect intersection to use AND instead of OR
		}
		if len(equalValues) > 0 {
			res = append(res, field+` IN (`+A.StrJoin(equalValues, `,`)+`)`)
		}
		if len(inequalValues) > 0 {
			res = append(res, field+` NOT IN (`+A.StrJoin(inequalValues, `,`)+`)`)
		}
	case Tt.String:
		equalValues := []string{}
		inequalValues := []string{}
		for _, str := range values {
			operator, rhs := splitOperatorValue(str)
			hasWildcard := S.Contains(rhs, `*`)
			if operator == `=` {
				if hasWildcard {
					rhs = S.Replace(rhs, `*`, `%`)
					operator = ` LIKE `
					res = append(res, field+operator+S.Z(rhs))
					continue
				}
				equalValues = append(equalValues, S.Z(rhs))
			} else if operator == `<>` {
				if hasWildcard {
					rhs = S.Replace(rhs, `*`, `%`)
					operator = ` NOT LIKE `
					res = append(res, field+operator+S.Z(rhs))
					continue
				}
				inequalValues = append(inequalValues, S.Z(rhs))
			} else {
				res = append(res, field+operator+S.Z(rhs))
			}
		}
		if len(equalValues) > 0 {
			res = append(res, field+` IN (`+A.StrJoin(equalValues, `,`)+`)`)
		}
		if len(inequalValues) > 0 {
			res = append(res, field+` NOT IN (`+A.StrJoin(inequalValues, `,`)+`)`)
		}
	case Tt.Array: // assume geo
		// TODO: do geoquery, but with sql: https://t.me/tarantool/15882
	case Tt.Boolean: // ignore for now
	}
	// TODO: return error

	return []string{}
}

func splitOperatorValue(str string) (op string, rhs string) {
	l := len(str)
	if l < 1 {
		op = `=`
		return
	}
	equal := l > 1 && str[1] == '='
	startCh := 0
	if str[0] == '>' {
		startCh = 1
		if equal {
			startCh = 2
		}
		op = str[:startCh]
	} else if str[0] == '<' {
		startCh = 1
		if equal {
			startCh = 2
		} else if l > 1 && str[1] == '>' {
			startCh = 2
		}
		op = str[:startCh]
	} else {
		op = `=`
	}
	rhs = str[startCh:]
	return
}

func (p *PaginationOut) CalculatePages(total int) {
	p.Total = total
	p.Pages = total / p.PerPage
}
