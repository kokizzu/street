package zCrud

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file pagination.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type pagination.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type pagination.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type pagination.go
//go:generate farify doublequote --file pagination.go

type PaginationIn struct {
	Page  int `json:"page" form:"page" query:"page" long:"page" msg:"page"`
	Limit int `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`

	// filter AND by column, if value is array then filter OR on that field
	Filters map[string]any `json:"filters" form:"filters" query:"filters" long:"filters" msg:"filters"`

	// Order: [+col1, -col2] (+ is ascending, - is descending)
	Order []string `json:"order" form:"order" query:"order" long:"order" msg:"order"`
}

type PaginationOut struct {
	Page  int `json:"page" form:"page" query:"page" long:"page" msg:"page"`
	Limit int `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`
	Total int `json:"total" form:"total" query:"total" long:"total" msg:"total"`

	Order []string `json:"order" form:"order" query:"order" long:"order" msg:"order"`
}
