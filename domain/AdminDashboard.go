package domain

import (
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/saAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminDashboard.go
//go:generate farify doublequote --file AdminDashboard.go

type (
	AdminDashboardIn struct {
		RequestCommon
	}

	AdminDashboardOut struct {
		ResponseCommon

		RegisteredUserTotal int64 `json:"registeredUserTotal" form:"registeredUserTotal" query:"registeredUserTotal" long:"registeredUserTotal" msg:"registeredUserTotal"`
		RegisteredUserToday int64 `json:"registeredUserToday" form:"registeredUserToday" query:"registeredUserToday" long:"registeredUserToday" msg:"registeredUserToday"`

		RequestsPerDate   map[string]int `json:"requestsPerDate" form:"requestsPerDate" query:"requestsPerDate" long:"requestsPerDate" msg:"requestsPerDate"`
		UniqueUserPerDate map[string]int `json:"uniqueUserPerDate" form:"uniqueUserPerDate" query:"uniqueUserPerDate" long:"uniqueUserPerDate" msg:"uniqueUserPerDate"`
		UniqueIpPerDate   map[string]int `json:"uniqueIpPerDate" form:"uniqueIpPerDate" query:"uniqueIpPerDate" long:"uniqueIpPerDate" msg:"uniqueIpPerDate"`

		CountPerActionsPerDate map[string]map[string]int `json:"countPerActionsPerDate" form:"countPerActionsPerDate" query:"countPerActionsPerDate" long:"countPerActionsPerDate" msg:"countPerActionsPerDate"`
	}
)

const (
	AdminDashboardAction = `admin/dashboard`
)

func (d *Domain) AdminDashboard(in *AdminDashboardIn) (out AdminDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	rq := rqAuth.NewUsers(d.AuthOltp)
	out.RegisteredUserTotal = rq.Total()
	out.RegisteredUserToday = rq.CountUserRegisterToday()

	sa := saAuth.NewActionLogs(d.AuthOlap)
	out.RequestsPerDate = sa.StatRequestsPerDate()
	out.UniqueUserPerDate = sa.StatUniqueUserPerDate()
	out.UniqueIpPerDate = sa.StatUniqueIpPerDate()

	out.CountPerActionsPerDate = sa.StatPerActionsPerDate()
	return
}
