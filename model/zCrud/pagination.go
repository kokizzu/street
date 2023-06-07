package zCrud

import (
	"fmt"

	"github.com/kokizzu/gotro/A"
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
	Filters map[string]any `json:"filters" form:"filters" query:"filters" long:"filters" msg:"filters"`

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

	Filters map[string]any `json:"filters" form:"filters" query:"filters" long:"filters" msg:"filters"`

	Order []string `json:"order" form:"order" query:"order" long:"order" msg:"order"`
}

func (p *PaginationOut) LimitOffsetSql(in *PaginationIn) string {
	offset := in.Offset()
	p.PerPage = in.Limit()
	p.Page = offset/p.PerPage + 1
	return fmt.Sprintf(`LIMIT %d OFFSET %d`, p.PerPage, offset)
}

func (p *PaginationOut) WhereOrderSql(filters map[string]any, orders []string, fieldToType map[string]string) (string, string) {
	// TODO: generate filter order
	// TODO: check if array then generate OR/IN filter

	whereAnd := []string{}
	orderBy := []string{}

	whereAndSql := `WHERE 1=1`
	if len(whereAnd) > 0 {
		whereAndSql = `WHERE (` + A.StrJoin(whereAnd, `) 
	AND (`) + `)`
	}

	orderBySql := `ORDER BY 1`
	if len(orderBy) > 0 {
		orderBySql = `ORDER BY ` + A.StrJoin(orderBy, `, `)
	}
	return whereAndSql, orderBySql
}

func (p *PaginationOut) CalculatePages(total int) {
	p.Total = total
	p.Pages = total / p.PerPage
}
