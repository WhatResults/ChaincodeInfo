package routers

import "github.com/gofiber/fiber/v2"

// 首页
func Index() Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"CurrentUser": GetCurrentUser(c),
		})
	}
}