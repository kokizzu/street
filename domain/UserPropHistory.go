package domain

import (
	"street/model/mProperty/rqProperty"
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

func (d *Domain) UserPropHistory(in *UserPropHistoryIn) (out UserPropHistoryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	hist := rqProperty.NewPropertyHistory(d.PropOltp)

	out.History = hist.FindByPropertyKey(in.PropertyKey)

	// find property being referenced, since the FK is not prop.Id
	prop := rqProperty.NewProperty(d.PropOltp)
	prop.UniqPropKey = in.PropertyKey
	if prop.FindByUniqPropKey() {
		out.refId = prop.Id
	}

	return

}
