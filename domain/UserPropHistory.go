package domain

import (
	"street/model/mProperty/rqProperty"
	"strings"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPropHistory.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPropHistory.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPropHistory.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPropHistory.go
//go:generate farify doublequote --file UserPropHistory.go

type (
	UserPropHistoryIn struct {
		RequestCommon `json:"request_common"`

		PropertyKey string `json:"propertyKey" form:"propertyKey" query:"propertyKey" long:"propertyKey" msg:"propertyKey"`
	}

	UserPropHistoryOut struct {
		ResponseCommon

		History []*rqProperty.PropertyHistory `json:"history" form:"history" query:"history" long:"history" msg:"history"`
	}
)

const (
	UserPropHistoryAction = `user/propHistory`
)

func splitStringWithChars(s, char string) []string {
	var result []string
	start := 0

	for {
		index := strings.Index(s[start:], char)
		if index == -1 {
			result = append(result, s[start:])
			break
		}
		result = append(result, s[start:start+index])
		start += index + len(char)
	}

	return result
}

func (d *Domain) UserPropHistory(in *UserPropHistoryIn) (out UserPropHistoryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	hist := rqProperty.NewPropertyHistory(d.PropOltp)

	if in.PropertyKey == "" {
		// throw response error code
		return
	}

	// Get property serial number from property key
	propKeyArrs := splitStringWithChars(in.PropertyKey, "#")

	if len(propKeyArrs) == 0 {
		return
	}
	propSerialNumber := propKeyArrs[0]

	out.History = hist.FindBySerialNumber(propSerialNumber)

	// find property being referenced, since the FK is not prop.Id
	prop := rqProperty.NewProperty(d.PropOltp)
	prop.UniqPropKey = in.PropertyKey
	if prop.FindByUniqPropKey() {
		out.refId = prop.Id
	}

	return

}
