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
	DataTypeString DataType = `string`
	DataTypeInt    DataType = `int`
	DataTypeFloat  DataType = `float`
	DataTypeIntArr DataType = `intArr`

	InputTypeText     InputType = `text`
	InputTypeTextArea InputType = `textarea`
	InputTypeEmail    InputType = `email`
	InputTypePassword InputType = `password`
	InputTypeCombobox InputType = `combobox`
	InputTypeCheckbox InputType = `checkbox`
	InputTypeHidden   InputType = `hidden`
	InputTypeDateTime InputType = `datetime`

	ValidationRequired Validation = `required`
	ValidationMinLen   Validation = `minLen`
	ValidationMaxLen   Validation = `maxLen`
	ValidationRegex    Validation = `regex`
)

type Field struct {
	Name        string    `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Label       string    `json:"label" form:"label" query:"label" long:"label" msg:"label"`
	Description string    `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	DataType    DataType  `json:"type" form:"type" query:"type" long:"type" msg:"type"`
	InputType   InputType `json:"inputType" form:"inputType" query:"inputType" long:"inputType" msg:"inputType"`
	ReadOnly    bool      `json:"readOnly" form:"readOnly" query:"readOnly" long:"readOnly" msg:"readOnly"`

	Validations map[Validation]any `json:"validations" form:"validations" query:"validations" long:"validations" msg:"validations"`

	// endpoint to find the combobox reference, if combobox/select source for autocomplete is too large
	RefEndpoint string `json:"refEndpoint" form:"refEndpoint" query:"refEndpoint" long:"refEndpoint" msg:"refEndpoint"`
}
