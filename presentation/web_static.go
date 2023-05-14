package presentation

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"

	"street/domain"
	"street/model/mAuth/rqAuth"
)

func WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get("/", func(c *fiber.Ctx) error {
		// TODO: fetch from session and other stuff need for first render
		var in domain.UserProfileIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction)
		var user *rqAuth.Users
		if err == nil {
			ctx := context.Background()
			in.FromFiberCtx(c, ctx)
			out := d.UserProfile(&in)
			user = out.User
			L.Print(out)
		}
		return views.RenderIndex(c, M.SX{
			`title`: `Street`,
			`user`:  user,
		})
	})

	fw.Get("/guest/verifyEmail", func(c *fiber.Ctx) error {
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
			L.Print(out)
		}
		return views.RenderGuestVerifyEmail(c, M.SX{
			`title`:    `Email Verification`,
			`verified`: verified,
			`email`:    email,
			`error`:    errStr,
		})
	})

	fw.Get(`/guest/resetPassword`, func(c *fiber.Ctx) error {
		return views.RenderGuestResetPassword(c, M.SX{
			`title`: `Reset Password`,
		})
	})
}
