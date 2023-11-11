package domain

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/X"

	"street/model/mAuth"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
	"street/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminFeedbacks.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminFeedbacks.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminFeedbacks.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminFeedbacks.go
//go:generate farify doublequote --file AdminFeedbacks.go

type (
	AdminFeedbacksIn struct {
		RequestCommon

		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`

		// for modifying files
		Feedback rqAuth.Feedbacks `json:"feedback" form:"feedback" query:"feedback" long:"feedback" msg:"feedback"`

		// will be filled by default with form id=0
		WithMeta bool `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`

		Pager zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}

	AdminFeedbacksOut struct {
		ResponseCommon

		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`

		Meta *zCrud.Meta `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`

		Feedback *rqAuth.Feedbacks `json:"feedback" form:"feedback" query:"feedback" long:"feedback" msg:"feedback"`

		Feedbacks [][]any `json:"feedbacks" form:"feedbacks" query:"feedbacks" long:"feedbacks" msg:"feedbacks"`

		Users map[string]string `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	AdminFeedbacksAction = `admin/feedbacks`

	ErrAdminFeedbackIdNotFound = `feedback id not found`
	ErrAdminFeedbackSaveFailed = `feedback save failed`
)

var (
	AdminFeedbacksMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mAuth.Id,
				Label:     `ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mAuth.CreatedAt,
				Label:     `Created At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mAuth.CreatedBy,
				Label:     `User`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeMapping,
				Mapping:   `users`,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mAuth.UserMessage,
				Label:     `User Message`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mAuth.UpdatedBy,
				Label:     `Admin`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeMapping,
				Mapping:   `users`,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mAuth.UpdatedAt,
				Label:     `Updated At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mAuth.DeletedAt,
				Label:     `Deleted At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeHidden,
			},
			{
				Name:      mAuth.AdminReply,
				Label:     `Admin Reply`,
				ReadOnly:  false,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeTextArea,
			},
		},
	}
)

func (d *Domain) AdminFeedbacks(in *AdminFeedbacksIn) (out AdminFeedbacksOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.Feedback.Id

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminFeedbacksMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.Feedback.Id <= 0 {
			out.Meta = &AdminFeedbacksMeta
			return
		}

		fb := rqAuth.NewFeedbacks(d.AuthOltp)
		fb.Id = in.Feedback.Id
		if !fb.FindById() {
			out.SetError(400, ErrAdminFeedbackIdNotFound)
			return
		}
		out.Feedback = fb
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		fb := wcAuth.NewFeedbacksMutator(d.AuthOltp)
		fb.Id = in.Feedback.Id
		if fb.Id > 0 {
			if !fb.FindById() {
				out.SetError(400, ErrAdminFeedbackIdNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if fb.DeletedAt == 0 {
					fb.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if fb.DeletedAt > 0 {
					fb.SetDeletedAt(0)
				}
			}
		} else {
			fb.SetCreatedAt(in.UnixNow())
		}

		haveMutation := fb.SetAll(in.Feedback, M.SB{}, M.SB{})

		sendEmail := false

		if haveMutation {
			fb.SetUpdatedAt(in.UnixNow())
			fb.SetUpdatedBy(sess.UserId)
			if fb.Id == 0 {
				fb.SetCreatedAt(in.UnixNow())
				fb.SetCreatedBy(sess.UserId)
			} else {
				sendEmail = true
			}
		}
		if !fb.DoUpsert() {
			out.SetError(500, ErrAdminFeedbackSaveFailed)
			break
		}

		fb.Adapter = nil
		out.Feedback = &fb.Feedbacks

		if sendEmail {
			user := rqAuth.NewUsers(d.AuthOltp)
			user.Id = fb.CreatedBy
			if !user.FindById() {
				return
			}

			d.runSubtask(func() {
				err := d.Mailer.SendNotifReplyFeedback(user.Email,
					X.ToS(fb.Id),
					fb.AdminReply,
				)
				L.IsError(err, `SendNotifReplyFeedback`)
			})
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		fbs := rqAuth.NewFeedbacks(d.AuthOltp)
		out.Feedbacks = fbs.FindByPagination(&AdminFeedbacksMeta, &in.Pager, &out.Pager)

		fb := rqAuth.NewFeedbacks(nil)

		uniqUsers := M.SB{}
		for _, row := range out.Feedbacks {
			uniqUsers[X.ToS(row[fb.IdxCreatedBy()])] = true
			uniqUsers[X.ToS(row[fb.IdxCreatedAt()])] = true
		}

		users := rqAuth.NewUsers(d.AuthOltp)
		out.Users = users.EmailMapByIds(uniqUsers.KeysConcat(`,`))

	}

	return
}
