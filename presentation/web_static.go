package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"street/domain"
)

func WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get("/", func(c *fiber.Ctx) error {
		// TODO: fetch from session and other stuff need for first render
		return views.RenderIndex(c, M.SX{
			`title`: `Street`,
		})
	})

	fw.Get("/guest/verifyEmail", func(c *fiber.Ctx) error {
		return views.RenderGuestVerifyEmail(c, M.SX{
			`title`: `Email Verification`,
		})
	})

	fw.Get(`/guest/resetPassword`, func(c *fiber.Ctx) error {
		return views.RenderGuestResetPassword(c, M.SX{
			`title`: `Reset Password`,
		})
	})
}
