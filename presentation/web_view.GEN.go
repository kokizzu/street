package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)

var viewList = map[string]string{
	`Index`: `../svelte/index.html`, // ../svelte/index.svelte
}


func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Index`].Str(m))
}
