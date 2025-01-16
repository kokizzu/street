package presentation

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/kokizzu/gotro/C"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"

	"street/conf"
	"street/domain"
	"street/model/mAuth"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/saAuth"
	"street/model/mBusiness"
	"street/model/mBusiness/rqBusiness"
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/saProperty"
	"street/model/mStorage/rqStorage"
	"street/model/zCrud"
)

func (w *WebServer) WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/`, func(c *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(c, d)
		google := d.GuestExternalAuth(&domain.GuestExternalAuthIn{
			RequestCommon: in.RequestCommon,
			Provider:      domain.OauthGoogle,
		})
		google.ResponseCommon.DecorateSession(c)

		var userRegistered []mAuth.UserRegisterStat
		var realtorStats []mAuth.RealtorStat
		var buyerStats []mAuth.BuyerStat
		var revenues []*mBusiness.Revenue
		var orders []*mBusiness.Order
		var mostLoggedInUsers []saAuth.MostLoggedInUser
		var mostScannedAreas []saProperty.ScannedAreasToRender
		var mostScannedProperties []saProperty.ScannedPropertiesToRender

		if segments[domain.UserSegment] {
			actionLog := saAuth.NewActionLogs(d.AuthOlap)
			userRegistered = actionLog.FindUserRegistered()
			realtorStats = actionLog.FindRealtorActivity()
			mostLoggedInUsers = actionLog.FindMostLoggedInUsers(d.AuthOltp)
			buyerStats = actionLog.FindBuyerActivity()

			scannedArea := saProperty.NewScannedAreas(d.PropOlap)
			mostScannedAreas = scannedArea.FindMostScannedAreas()

			scannedProp := saProperty.NewScannedProperties(d.PropOlap)
			mostScannedProperties = scannedProp.FindMostScannedProperties(d.PropOltp)

			or := rqBusiness.NewSales(d.BusinessOltp)
			orders = or.FindOrdersAnnually()

			if segments[domain.AdminSegment] {
				r := rqBusiness.NewSales(d.BusinessOltp)
				revenues = r.FindRevenuesMonthly("XD") // default to current month
			} else {
				r := rqBusiness.NewSales(d.BusinessOltp)
				r.RealtorId = user.Id
				revenues = r.FindRealtorRevenuesMonthlyByRealtorId("XD") // default to current month
			}
		}
		return views.RenderIndex(c, M.SX{
			`title`:                   `HapSTR`,
			`user`:                    user,
			`google`:                  google.Link,
			`segments`:                segments,
			`user_registered`:         userRegistered,
			`realtor_stats`:           realtorStats,
			`buyer_stats`:             buyerStats,
			`revenues`:                revenues,
			`orders`:                  orders,
			`users_most_logged_in`:    mostLoggedInUsers,
			`most_scanned_areas`:      mostScannedAreas,
			`most_scanned_properties`: mostScannedProperties,
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

	fw.Get(`/`+domain.GuestOauthCallbackRedirectAction, func(c *fiber.Ctx) error {
		var in domain.GuestOauthCallbackRedirectIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthCallbackRedirectAction)
		var errStr, email string
		createdAt := int64(0)
		provider := `unknown`
		if err != nil {
			errStr = err.Error()
		} else {
			out := d.GuestOauthCallbackRedirect(&in)
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

	fw.Get(`/`+domain.GuestOauthTokenExchangeAction, func(c *fiber.Ctx) error {
		var in domain.GuestOauthTokenExchangeIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthTokenExchangeAction)
		var errStr, email string
		createdAt := int64(0)
		provider := `unknown`
		if err != nil {
			errStr = err.Error()
		} else {
			out := d.GuestOauthTokenExchange(&in)
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
		countryCode, propId := splitCountryPropId(ctx.Params(`propId`))
		out := d.GuestProperty(&domain.GuestPropertyIn{
			RequestCommon: in.RequestCommon,
			Id:            propId,
			CountryCode:   countryCode,
		})
		if out.Property != nil && out.Property.DeletedAt > 0 {
			return views.RenderError(ctx, M.SX{
				`error`: `property deleted`,
			})
		}

		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`title`: out.Error,
				`error`: out.Error,
			})
		}
		imgUrl := ""
		if len(out.Property.Images) != 0 {
			imgUrl = fmt.Sprintf("%s%s", w.Domain.WebCfg.WebProtoDomain, out.Property.Images[0])
		}
		const ISO8601 = "2006-01-02T15:04:05Z07:00"
		descr := out.Property.SizeM2 + ` m2`
		if out.Property.Bedroom > 0 {
			descr += `, ` + X.ToS(out.Property.Bedroom) + ` bedroom`
		}
		if out.Property.Bathroom > 0 {
			descr += `, ` + X.ToS(out.Property.Bathroom) + ` bathroom`
		}
		if out.Property.NumberOfFloors != `0` && out.Property.NumberOfFloors != `` {
			descr += `, ` + X.ToS(out.Property.NumberOfFloors) + ` floor`

		}
		title := `Property #` + X.ToS(out.Property.Id)
		if out.Property.Address != `` {
			title += ` on ` + out.Property.Address
		} else if out.Property.FormattedAddress != `` {
			title += ` on ` + out.Property.FormattedAddress
		}
		ogUrl := fmt.Sprintf("%s/%s/%s%d", w.Domain.WebCfg.WebProtoDomain, domain.GuestPropertyAction, countryCode, out.Property.Id)
		return views.RenderGuestPropertyPublic(ctx, M.SX{
			`title`:         S.XSS(title),
			`propItem`:      out.Property,
			`propExtraUS`:   out.PropertyExtraUS,
			`propertyMeta`:  out.Meta,
			`ogURL`:         ogUrl,
			`ogImgURL`:      imgUrl,
			`ogDescription`: S.XSS(descr),
			`ogCreatedAt`:   time.Unix(out.Property.CreatedAt, 0).Format(ISO8601),
			`ogUpdatedAt`:   time.Unix(out.Property.UpdatedAt, 0).Format(ISO8601),
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

	fw.Get(`/guest/download3dFile`, func(ctx *fiber.Ctx) error {
		queries := ctx.Queries()
		countryPropId := fmt.Sprintf("%s:%d", queries[`country`], S.ToU(queries[`propertyId`]))
		img3d := rqStorage.NewDesignFiles(d.StorOltp)
		img3d.CountryPropId = countryPropId
		if !img3d.FindByCountryPropId() {
			return views.RenderError(ctx, M.SX{
				`title`: `3D file not found`,
				`error`: `3D file not found, make sure country and property id is valid`,
			})
		}

		if ctx.SendFile(d.UploadDir+img3d.FilePath) != nil {
			return views.RenderError(ctx, M.SX{
				`title`: `3D file not found`,
				`error`: `3D file is either missing from the server or has been deleted. Please check again or contact support for help.`,
			})
		}

		ctx.Set("Content-Disposition", `attachment; filename="`+img3d.FilePath+`"`)
		return nil
	})

	fw.Get(`/`+domain.UserPropertyAction+`/:propId`, func(ctx *fiber.Ctx) error {
		in, _, _ := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		in.RequestCommon.Action = domain.UserPropertyAction
		countryCode, propId := splitCountryPropId(ctx.Params(`propId`))
		out := d.UserProperty(&domain.UserPropertyIn{
			RequestCommon: in.RequestCommon,
			Id:            propId,
			CountryCode:   countryCode,
		})
		if out.Property.DeletedAt > 0 {
			return views.RenderError(ctx, M.SX{
				`title`: `Property is deleted`,
				`error`: `Property is deleted`,
			})
		}
		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`title`: out.Error,
				`error`: out.Error,
			})
		}
		descr := out.Property.SizeM2 + ` m2`
		if out.Property.Bedroom > 0 {
			descr += `, ` + X.ToS(out.Property.Bedroom) + ` bedroom`
		}
		if out.Property.Bathroom > 0 {
			descr += `, ` + X.ToS(out.Property.Bathroom) + ` bathroom`
		}
		if out.Property.NumberOfFloors != `0` && out.Property.NumberOfFloors != `` {
			descr += `, ` + X.ToS(out.Property.NumberOfFloors) + ` floor`

		}
		title := `Property #` + X.ToS(out.Property.Id)
		if out.Property.Address != `` {
			title += ` on ` + out.Property.Address
		} else if out.Property.FormattedAddress != `` {
			title += ` on ` + out.Property.FormattedAddress
		}
		return views.RenderUserPropertyIndex(ctx, M.SX{ // TODO: habibi change to RenderUserProperty (create own view, do not share view)
			`title`:         S.XSS(title),
			`propItem`:      out.Property,
			`propertyMeta`:  out.Meta,
			`propHistories`: out.PropHistories,
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

	fw.Get(`/`+domain.UserBuyersAction, func(ctx *fiber.Ctx) error {
		var in domain.UserBuyersIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.RealtorRevenueAction)
		if err != nil {
			return err
		}
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserBuyers(&in)
		return views.RenderUserBuyers(ctx, M.SX{
			`title`:    `Buyer`,
			`user`:     user,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.UserListingsAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

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
		return views.RenderUserListings(ctx, M.SX{
			`title`:             `Listings`,
			`user`:              user,
			`segments`:          segments,
			`randomProps`:       props.Properties,
			`initialLatLong`:    []any{lat, lng},
			`defaultDistanceKm`: defaultDistanceKm,
		})
	})

	fw.Get(`/`+domain.UserListingsAction+`/:propId`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.UserListing(&domain.UserListingIn{
			RequestCommon: in.RequestCommon,
			Id:            X.ToU(ctx.Params(`propId`)),
		})
		if out.Property.DeletedAt > 0 {
			return views.RenderError(ctx, M.SX{
				`error`: `property deleted`,
			})
		}
		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`error`: out.Error,
			})
		}
		title := `Property #` + X.ToS(out.Property.Id)
		if out.Property.Address != `` {
			title += ` on ` + out.Property.Address
		} else if out.Property.FormattedAddress != `` {
			title += ` on ` + out.Property.FormattedAddress
		}
		return views.RenderUserListingsListing(ctx, M.SX{
			`title`:    title,
			`user`:     user,
			`segments`: segments,
			`property`: out.Property,
		})
	})

	fw.Get(`/realtor`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.RealtorOwnedProperties(&domain.RealtorOwnedPropertiesIn{
			RequestCommon: in.RequestCommon,
			ShowMeta:      true,
		})
		if len(out.Properties) == 0 {
			out.Properties = []rqProperty.PropertyWithNote{}
		}
		return views.RenderRealtor(ctx, M.SX{
			`title`:           `Realtor`,
			`user`:            user,
			`segments`:        segments,
			`ownedProperties`: out.Properties,
			`pager`:           out.Pager,
			`propertyMeta`:    out.Meta,
		})
	})
	fw.Get(`/`+domain.RealtorRevenueAction, func(ctx *fiber.Ctx) error {
		var in domain.RealtorRevenueIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.RealtorRevenueAction)
		if err != nil {
			return err
		}
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdList
		out := d.RealtorRevenue(&in)
		return views.RenderRealtorRevenue(ctx, M.SX{
			`title`:    `Realtor Revenue`,
			`user`:     user,
			`segments`: segments,
			`revenues`: out.Revenues,
		})
	})
	fw.Get(`/`+domain.RealtorPropertyAction, func(ctx *fiber.Ctx) error {
		// create new property
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderRealtorProperty(ctx, M.SX{
			`title`:     `Realtor Property`,
			`segments`:  segments,
			`user`:      user,
			`property`:  M.SX{},
			`countries`: conf.CountriesData,
		})
	})
	fw.Get(`/`+domain.RealtorPropertyAction+`Old`, func(ctx *fiber.Ctx) error {
		// create new property
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderRealtorPropertyOld(ctx, M.SX{
			`title`:     `Realtor Property`,
			`segments`:  segments,
			`user`:      user,
			`property`:  M.SX{},
			`countries`: conf.CountriesData,
		})
	})
	fw.Get(`/`+domain.RealtorPropertyAction+`Old/:propId`, func(ctx *fiber.Ctx) error {
		// edit property
		in, user, segments := userInfoFromContext(ctx, d)
		in.RequestCommon.Action = domain.RealtorPropertyAction
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.RealtorProperty(&domain.RealtorPropertyIn{
			RequestCommon: in.RequestCommon,
			Id:            X.ToU(ctx.Params(`propId`)),
		})
		if out.Property.DeletedAt > 0 {
			return views.RenderError(ctx, M.SX{
				`error`: `property deleted`,
			})
		}
		if out.Error != `` {
			L.Print(out.Error)
			return views.RenderError(ctx, M.SX{
				`error`: out.Error,
			})
		}
		title := `Property #` + X.ToS(out.Property.Id)
		if out.Property.Id == 0 {
			title = `Create a new property`
		} else if out.Property.Address != `` {
			title += ` on ` + out.Property.Address
		} else if out.Property.FormattedAddress != `` {
			title += ` on ` + out.Property.FormattedAddress
		}
		return views.RenderRealtorPropertyOld(ctx, M.SX{
			`title`:     title,
			`segments`:  segments,
			`user`:      user,
			`property`:  out.Property,
			`countries`: conf.CountriesData,
		})
	})
	fw.Get(`/admin`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		out := d.AdminDashboard(&domain.AdminDashboardIn{
			RequestCommon: in.RequestCommon,
		})
		return views.RenderAdmin(ctx, M.SX{
			`user`:                   user,
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
	fw.Get(`/`+domain.AdminRevenueAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminRevenueIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminRevenueAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdList
		out := d.AdminRevenue(&in)
		return views.RenderAdminRevenue(ctx, M.SX{
			`title`:    `Revenue`,
			`user`:     user,
			`segments`: segments,
			`revenues`: out.Revenues,
		})
	})
	fw.Get(`/`+domain.AdminFeedbacksAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminFeedbacksIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminFeedbacksAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminFeedbacks(&in)
		return views.RenderAdminFeedbacks(ctx, M.SX{
			`title`:     `Feedbacks`,
			`segments`:  segments,
			`user`:      user,
			`feedbacks`: out.Feedbacks,
			`fields`:    out.Meta.Fields,
			`pager`:     out.Pager,
			`users`:     out.Users,
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
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminUsers(&in)
		return views.RenderAdminUsers(ctx, M.SX{
			`title`:    `Users`,
			`user`:     user,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})
	fw.Get(`/`+domain.AdminPropertiesUSAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminPropertiesUSIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminPropertiesUSAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminPropertiesUS(&in)
		return views.RenderAdminPropertiesUS(ctx, M.SX{
			`title`:      `Properties US`,
			`user`:       user,
			`segments`:   segments,
			`properties`: out.Properties,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})
	fw.Get(`/`+domain.AdminPropertiesTWAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminPropertiesTWIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminPropertiesTWAction)
		if err != nil {
			return err
		}
		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminPropertiesTW(&in)
		return views.RenderAdminPropertiesTW(ctx, M.SX{
			`title`:      `Properties TW`,
			`user`:       user,
			`segments`:   segments,
			`properties`: out.Properties,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
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
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminProperties(&in)
		return views.RenderAdminProperties(ctx, M.SX{
			`title`:      `Properties`,
			`user`:       user,
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
			`countries`:      conf.CountriesData,
		})
	})
	// TODO: Old version of UserNearbyFacilities with typo (Remove when no longer used)
	fw.Post("/user/nearbyFacilitites", func(c *fiber.Ctx) error {
		in := domain.UserNearbyFacilitiesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserNearbyFacilitiesAction); err != nil {
			return nil
		}
		out := d.UserNearbyFacilities(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
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
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		out := d.AdminAccessLogs(&in)
		return views.RenderAdminAccessLog(ctx, M.SX{
			`user`:     user,
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
			`title`:    `Files`,
			`segments`: segments,
			`files`:    out.Files,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})
	fw.Get(`/`+domain.Admin3DFilesAction, func(ctx *fiber.Ctx) error {
		var in domain.Admin3DFilesIn
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
		out := d.Admin3DFiles(&in)
		return views.RenderAdmin3DFiles(ctx, M.SX{
			`title`:    `3D Design Files`,
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

func splitCountryPropId(param string) (string, uint64) {
	// starts with 2 letter country code
	if len(param) > 2 && C.IsAlpha(param[0]) && C.IsAlpha(param[1]) {
		return S.ToUpper(param[:2]), S.ToU(param[2:])
	}
	return ``, S.ToU(param)
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
	sess := d.MustAdmin(in, &check)
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
	var segments = M.SB{}
	out := d.UserProfile(&domain.UserProfileIn{
		RequestCommon: rc,
	})
	user = out.User
	segments = out.Segments
	return user, segments
}
