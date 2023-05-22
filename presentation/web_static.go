package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"street/domain"
	"street/model/mAuth/rqAuth"
)

func WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/`, func(c *fiber.Ctx) error {
		var in domain.UserProfileIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction)
		var user *rqAuth.Users
		if err == nil {
			out := d.UserProfile(&in)
			user = out.User
		}
		google := d.GuestExternalAuth(&domain.GuestExternalAuthIn{
			RequestCommon: in.RequestCommon,
			Provider:      domain.OauthGoogle,
		})
		google.DecorateSession(c)
		return views.RenderIndex(c, M.SX{
			`title`:  `Street`,
			`user`:   user,
			`google`: google.Link,
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
}
