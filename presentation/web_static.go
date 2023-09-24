package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/X"

	"street/domain"
	"street/model/mAuth/rqAuth"
	"street/model/mProperty/rqProperty"
	"street/model/zCrud"
	"street/model/zImport"
)

func (w *WebServer) WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/`, func(c *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(c, d)
		google := d.GuestExternalAuth(&domain.GuestExternalAuthIn{
			RequestCommon: in.RequestCommon,
			Provider:      domain.OauthGoogle,
		})
		google.ResponseCommon.DecorateSession(c)
		const lat = 23.708740595481036
		const lng = 120.78636646165934
		const defaultDistanceKm = 20
		var props domain.UserSearchPropOut
		if user != nil && user.Id > 0 {
			copy := in.RequestCommon
			copy.Action = domain.UserSearchPropAction
			props = d.UserSearchProp(&domain.UserSearchPropIn{
				RequestCommon: copy,
				CenterLat:     lat,
				CenterLong:    lng,
				Offset:        0,
				Limit:         0,
				MaxDistanceKM: defaultDistanceKm,
			})
		}
		return views.RenderIndex(c, M.SX{
			`title`:  `Street`,
			`user`:   user,
			`google`: google.Link,

			`segments`:          segments,
			`randomProps`:       props.Properties,
			`initialLatLong`:    []any{lat, lng},
			`defaultDistanceKm`: defaultDistanceKm,
		})
	})

	fw.Get(`/`+domain.GuestExternalAuthAction, func(c *fiber.Ctx) error {
		in := domain.GuestExternalAuthIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestExternalAuthAction); err != nil {
			return err
		}
		out := d.GuestExternalAuth(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
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

	fw.Get(`/`+domain.GuestPropertyAction+`/:propId`, func(ctx *fiber.Ctx) error {
		in, _, _ := userInfoFromContext(ctx, d)
		in.RequestCommon.Action = domain.GuestPropertyAction
		out := d.GuestProperty(&domain.GuestPropertyIn{
			RequestCommon: in.RequestCommon,
			Id:            X.ToU(ctx.Params(`propId`)),
		})
		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`error`: out.Error,
			})
		}
		return views.RenderGuestProperty(ctx, M.SX{
			`title`:        `Property`,
			`propItem`:     out.Property,
			`propertyMeta`: out.Meta,
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
		in, _, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderBuyer(ctx, M.SX{
			`title`:    `Buyer`,
			`segments`: segments,
		})
	})
	fw.Get(`/realtor`, func(ctx *fiber.Ctx) error {
		in, _, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.RealtorOwnedProperties(&domain.RealtorOwnedPropertiesIn{
			RequestCommon: in.RequestCommon,
			ShowMeta:      true,
		})
		if len(out.Properties) == 0 {
			out.Properties = []rqProperty.Property{}
		}
		return views.RenderRealtor(ctx, M.SX{
			`title`:           `Realtor`,
			`segments`:        segments,
			`ownedProperties`: out.Properties,
			`pager`:           out.Pager,
			`propertyMeta`:    out.Meta,
		})
	})
	fw.Get(`/`+domain.RealtorPropertyAction, func(ctx *fiber.Ctx) error {
		// create new property
		in, _, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderRealtorProperty(ctx, M.SX{
			`title`:    `Realtor Property`,
			`segments`: segments,
			`property`: M.SX{},
		})
	})
	fw.Get(`/`+domain.RealtorPropertyAction+`/:propId`, func(ctx *fiber.Ctx) error {
		// edit property
		in, _, segments := userInfoFromContext(ctx, d)
		in.RequestCommon.Action = domain.RealtorPropertyAction
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.RealtorProperty(&domain.RealtorPropertyIn{
			RequestCommon: in.RequestCommon,
			Id:            X.ToU(ctx.Params(`propId`)),
		})
		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`error`: out.Error,
			})
		}
		return views.RenderRealtorProperty(ctx, M.SX{
			`title`:    `Realtor Property`,
			`segments`: segments,
			`property`: out.Property,
		})
	})
	fw.Get(`/admin`, func(ctx *fiber.Ctx) error {
		in, _, segments := userInfoFromContext(ctx, d)
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
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
			`registeredUserToday`:    out.RegisteredUserToday,
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
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
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
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminProperties(&in)
		return views.RenderAdminProperties(ctx, M.SX{
			`title`:      `Properties`,
			`segments`:   segments,
			`properties`: out.Properties,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})
	fw.Get(`/`+domain.AdminPropHistoriesAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminPropHistoriesIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminPropHistoriesAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminPropHistories(&in)
		return views.RenderAdminPropHistories(ctx, M.SX{
			`title`:         `Prop Histories`,
			`segments`:      segments,
			`propHistories`: out.PropHistories,
			`fields`:        out.Meta.Fields,
			`pager`:         out.Pager,
		})
	})
	fw.Get(`/user`, func(ctx *fiber.Ctx) error {
		countriesData, _ := zImport.GoogleSheetCountryDataToJson("1TmAjrclFHUwDA1487ifQjX4FzYt9y7eJ0gwyxtwZMJU", 522117981)
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		in.RequestCommon.Action = domain.UserSessionsActiveAction
		out := d.UserSessionsActive(&domain.UserSessionsActiveIn{
			RequestCommon: in.RequestCommon,
		})
		return views.RenderUser(ctx, M.SX{
			`title`:          `Profile`,
			`user`:           user,
			`segments`:       segments,
			`activeSessions`: out.SessionsActive,
			`countryData`:    countriesData,
		})
	})
	fw.Get(`/`+domain.GuestAutoLoginAction, func(ctx *fiber.Ctx) error {
		var in domain.GuestAutoLoginIn
		_ = webApiParseInput(ctx, &in.RequestCommon, &in, domain.GuestAutoLoginAction)
		in.Uid = ctx.Query(`uid`)
		in.Token = ctx.Query(`token`)
		in.Path = ctx.Query(`path`)
		out := d.GuestAutoLogin(&in)

		if out.Error != `` {
			return views.RenderError(ctx, M.SX{
				`error`: out.Error,
			})
		}
		return in.ToFiberCtx(ctx, out, &out.ResponseCommon, in)
	})
	fw.Get(`/`+domain.AdminAccessLogsAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminAccessLogsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminUsersAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		out := d.AdminAccessLogs(&in)
		return views.RenderAdminAccessLog(ctx, M.SX{
			`title`:    `Access Log`,
			`segments`: segments,
			`logs`:     out.Logs,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})
	fw.Get(`/`+domain.AdminFilesAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminFilesIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminUsersAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = `list`
		out := d.AdminFiles(&in)
		return views.RenderAdminFiles(ctx, M.SX{
			`title`:    `Access Log`,
			`segments`: segments,
			`files`:    out.Files,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})
	fw.All(`/`+domain.GuestFilesAction+`/:base62id-:modifier.:ext`, func(ctx *fiber.Ctx) error {
		method := ctx.Method()
		if method != fiber.MethodGet && method != fiber.MethodHead {
			return nil
		}
		var in domain.GuestFilesIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.GuestFilesAction)
		if err != nil {
			return err
		}

		in.RequestCommon.Action = domain.GuestFilesAction
		in.Base62id = utils.CopyString(ctx.Params(`base62id`))
		in.Modifier = utils.CopyString(ctx.Params(`modifier`))
		in.Ext = utils.CopyString(ctx.Params(`ext`))

		out := d.GuestFiles(&in)
		ctx.Set("content-type", out.ContentType)
		_, err = ctx.Write(out.Raw)
		return err
	})
	fw.Get(`/debug`, func(ctx *fiber.Ctx) error {
		return views.RenderDebug(ctx, M.SX{})
	})
}

func notLogin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	sess := d.MustLogin(in, &check)
	if sess == nil {
		_ = views.RenderError(ctx, M.SX{
			`error`: check.Error,
		})
		return true
	}
	return false
}

func notAdmin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	sess := d.MustLogin(in, &check)
	if sess == nil {
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
