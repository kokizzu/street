package domain

import (
	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserSendFeedback.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserSendFeedback.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserSendFeedback.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserSendFeedback.go
//go:generate farify doublequote --file UserSendFeedback.go

type (
	UserSendFeedbackIn struct {
		RequestCommon

		UserMessage string `json:"userMessage" form:"userMessage" query:"userMessage" long:"userMessage" msg:"userMessage"`
	}

	UserSendFeedbackOut struct {
		ResponseCommon
	}
)

const (
	UserSendFeedbackAction = `user/sendFeedback`

	ErrUserSendFeedbackFailed = `failed to send user feedback`
)

func (d *Domain) UserSendFeedback(in *UserSendFeedbackIn) (out UserSendFeedbackOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}
	out.refId = sess.UserId

	feedback := wcAuth.NewFeedbacksMutator(d.AuthOltp)
	feedback.CreatedBy = sess.UserId
	feedback.CreatedAt = in.UnixNow()
	feedback.UserMessage = in.UserMessage

	if !feedback.DoInsert() {
		out.SetError(500, ErrUserSendFeedbackFailed)
		return
	}

	return
}
