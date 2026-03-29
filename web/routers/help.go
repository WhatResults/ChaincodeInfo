package routers

import (
	"ChaincodeInfo/web/models"

	"github.com/gofiber/fiber/v2"
)

// 帮助
func Help() Handler {
	return func(c *fiber.Ctx) error {
		data := &struct {
			CurrentUser models.User
		}{
			CurrentUser: curUser,
		}
		return c.Render("help", data)
	}
}
