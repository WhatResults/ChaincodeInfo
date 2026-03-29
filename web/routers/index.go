package routers

import "github.com/gofiber/fiber/v2"

// 登出
func Index() Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	}
}
