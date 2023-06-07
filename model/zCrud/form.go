package zCrud

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file form.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type form.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type form.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type form.go
//go:generate farify doublequote --file form.go

type DataType string

type InputType string

type Validation string

const (
	String DataType = `string`
	Int    DataType = `int`
	Float  DataType = `float`
	IntArr DataType = `intArr`

	Text     InputType = `text`
	TextArea InputType = `textarea`
	Password InputType = `password`
	Combobox InputType = `combobox`
	Checkbox InputType = `checkbox`

	Required Validation = `required`
	MinLen   Validation = `minLen`
	MaxLen   Validation = `maxLen`
	Regex    Validation = `regex`
)

type Fields struct {
	Name        string    `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Label       string    `json:"label" form:"label" query:"label" long:"label" msg:"label"`
	Description string    `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	DataType    DataType  `json:"type" form:"type" query:"type" long:"type" msg:"type"`
	InputType   InputType `json:"inputType" form:"inputType" query:"inputType" long:"inputType" msg:"inputType"`

	Validations map[Validation]any `json:"validations" form:"validations" query:"validations" long:"validations" msg:"validations"`

	// endpoint to find the combobox reference
	RefEndpoint string `json:"refEndpoint" form:"refEndpoint" query:"refEndpoint" long:"refEndpoint" msg:"refEndpoint"`
}
