package routers

import (
	"ChaincodeInfo/web/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New()

// 获取当前登录用户
func GetCurrentUser(c *fiber.Ctx) models.User {
	sess, err := Store.Get(c)
	if err != nil {
		return models.User{}
	}

	user := models.User{}

	if v := sess.Get("loginName"); v != nil {
		if s, ok := v.(string); ok {
			user.LoginName = s
		}
	}

	if v := sess.Get("isAdmin"); v != nil {
		if b, ok := v.(bool); ok {
			user.IsAdmin = b
		}
	}

	return user
}

// 需要登录
func RequireLogin(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Redirect("/")
	}

	if sess.Get("loginName") == nil {
		return c.Redirect("/")
	}

	return c.Next()
}

// 需要管理员
func RequireAdmin(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString("无权限访问")
	}

	v := sess.Get("isAdmin")
	isAdmin, ok := v.(bool)
	if !ok || !isAdmin {
		return c.Status(fiber.StatusForbidden).SendString("无权限访问")
	}

	return c.Next()
}