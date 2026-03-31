package routers

import (
	"ChaincodeInfo/web/models"

	"github.com/gofiber/fiber/v2"
)

var users = []models.User{
	{
		LoginName: "admin",
		Password:  "123456",
		IsAdmin:   true,
	},
	{
		LoginName: "user",
		Password:  "123456",
		IsAdmin:   false,
	},
}

func LoginView() Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"Flag": false,
		})
	}
}

func Login() Handler {
	return func(c *fiber.Ctx) error {
		loginName := c.FormValue("loginName")
		password := c.FormValue("password")

		for _, user := range users {
			if user.LoginName == loginName && user.Password == password {
				sess, err := Store.Get(c)
				if err != nil {
					return c.Render("login", fiber.Map{
						"Flag": true,
					})
				}

				sess.Set("loginName", user.LoginName)
				sess.Set("isAdmin", user.IsAdmin)

				if err := sess.Save(); err != nil {
					return c.Render("login", fiber.Map{
						"Flag": true,
					})
				}

				return c.Redirect("/index")
			}
		}

		return c.Render("login", fiber.Map{
			"Flag": true,
		})
	}
}

func Logout() Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err == nil {
			_ = sess.Destroy()
		}
		return c.Redirect("/")
	}
}