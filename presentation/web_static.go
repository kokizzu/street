package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"street/domain"
	"street/model/mAuth/rqAuth"
	"street/model/zCrud"
)

func WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/`, func(c *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(c, d)
		google := d.GuestExternalAuth(&domain.GuestExternalAuthIn{
			RequestCommon: in.RequestCommon,
			Provider:      domain.OauthGoogle,
		})
		google.ResponseCommon.DecorateSession(c)
		return views.RenderIndex(c, M.SX{
			`title`:  `Street`,
			`user`:   user,
			`google`: google.Link,

			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.GuestOauthCallbackAction, func(c *fiber.Ctx) error {
		var in domain.GuestOauthCallbackIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthCallbackAction)
		var errStr, email string
		createdAt := int64(0)
		provider := `unknown`
		if err != nil {
			errStr = err.Error()
		} else {
			out := d.GuestOauthCallback(&in)
			errStr = out.Error
			email = out.Email
			provider = out.Provider
			createdAt = out.CurrentUser.CreatedAt
			out.DecorateSession(c)
		}
		return views.RenderGuestOauthCallback(c, M.SX{
			`title`:     `OAuth from ` + provider,
			`email`:     email,
			`error`:     errStr,
			`createdAt`: createdAt,
		})
	})

	fw.Get(`/`+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		var in domain.GuestVerifyEmailIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction)
		var errStr, email string
		var verified bool
		if err != nil {
			errStr = err.Error()
		} else {
			out := d.GuestVerifyEmail(&in)
			verified = out.Ok
			errStr = out.Error
			email = out.Email
			out.DecorateSession(c)
		}
		return views.RenderGuestVerifyEmail(c, M.SX{
			`title`:    `Email Verification`,
			`verified`: verified,
			`email`:    email,
			`error`:    errStr,
		})
	})

	fw.Get(`/`+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		return views.RenderGuestResetPassword(c, M.SX{
			`title`: `Reset Password`,
		})
	})

	fw.Get(`/privacy`, func(ctx *fiber.Ctx) error {
		return views.RenderPrivacy(ctx, M.SX{
			`title`: `HapSTR Privacy Policy`,
		})
	})

	fw.Get(`/tos`, func(ctx *fiber.Ctx) error {
		return views.RenderTos(ctx, M.SX{
			`title`: `HapSTR Terms of Service`,
		})
	})

	fw.Get(`/buyer`, func(ctx *fiber.Ctx) error {
		_, _, segments := userInfoFromContext(ctx, d)
		return views.RenderBuyer(ctx, M.SX{
			`title`:    `Buyer`,
			`segments`: segments,
		})
	})
	fw.Get(`/realtor`, func(ctx *fiber.Ctx) error {
		_, _, segments := userInfoFromContext(ctx, d)
		return views.RenderRealtor(ctx, M.SX{
			`title`:    `Realtor`,
			`segments`: segments,
		})
	})
	fw.Get(`/admin`, func(ctx *fiber.Ctx) error {
		in, _, segments := userInfoFromContext(ctx, d)
		if notAdmin(ctx, d, in.RequestCommon) {
			return nil
		}
		out := d.AdminDashboard(&domain.AdminDashboardIn{
			RequestCommon: in.RequestCommon,
		})
		return views.RenderAdmin(ctx, M.SX{
			`title`:                  `Admin`,
			`segments`:               segments,
			`uniqueIpPerDate`:        out.UniqueIpPerDate,
			`requestsPerDate`:        out.RequestsPerDate,
			`uniqueUserPerDate`:      out.UniqueUserPerDate,
			`registeredUserTotal`:    out.RegisteredUserTotal,
			`countPerActionsPerDate`: out.CountPerActionsPerDate,
		})
	})
	fw.Get(`/`+domain.AdminUsersAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminUsersIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminUsersAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return nil
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Action = zCrud.ActionList
		out := d.AdminUsers(&in)
		return views.RenderAdminUsers(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})
	fw.Get(`/`+domain.AdminPropertiesAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminPropertiesIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminPropertiesAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return nil
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Action = zCrud.ActionList
		out := d.AdminProperties(&in)
		return views.RenderAdminProperties(ctx, M.SX{
			`title`:      `Properties`,
			`segments`:   segments,
			`properties`: out.Properties,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})
	fw.Get(`/user`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return nil
		}
		return views.RenderUser(ctx, M.SX{
			`title`:    `Profile`,
			`user`:     user,
			`segments`: segments,
		})
	})
}

func notLogin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	if d.MustLogin(in, &check) == nil {
		_ = views.RenderError(ctx, M.SX{
			`error`: check.Error,
		})
		return true
	}
	return false
}

func notAdmin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	if d.MustAdmin(in, &check) == nil {
		_ = views.RenderError(ctx, M.SX{
			`error`: check.Error,
		})
		return true
	}
	return false
}

func userInfoFromContext(c *fiber.Ctx, d *domain.Domain) (domain.UserProfileIn, *rqAuth.Users, M.SB) {
	var in domain.UserProfileIn
	err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction)
	var user *rqAuth.Users
	segments := M.SB{}
	if err == nil {
		out := d.UserProfile(&in)
		user = out.User
		segments = out.Segments
	}
	return in, user, segments
}

func userInfoFromRequest(rc domain.RequestCommon, d *domain.Domain) (*rqAuth.Users, M.SB) {
	var user *rqAuth.Users
	segments := M.SB{}
	out := d.UserProfile(&domain.UserProfileIn{
		RequestCommon: rc,
	})
	user = out.User
	segments = out.Segments
	return user, segments
}
