package routers

import (
	"ChaincodeInfo/web/models"

	"github.com/gofiber/fiber/v2"
)

var (
	// 内置用户
	_builtin_users = []models.User{
		{LoginName: "admin", Password: "123456", IsAdmin: true},
		{LoginName: "creater", Password: "123456", IsAdmin: true},
		{LoginName: "visitor", Password: "123456", IsAdmin: false},
		{LoginName: "zhangsan", Password: "123456", IsAdmin: false},
	}

	// 当前用户
	curUser models.User
)

// 登录页面
func LoginView() Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	}
}

// 登录逻辑
func Login() Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("loginName")
		password := c.FormValue("password")

		var flag bool
		for _, user := range _builtin_users {
			if user.LoginName == username && user.Password == password {
				curUser = user
				flag = true
				break
			}
		}

		data := &struct {
			CurrentUser models.User
			Flag        bool
		}{
			CurrentUser: curUser,
			Flag:        false,
		}

		if flag {
			// 登录成功
			return c.Render("index", data)
		} else {
			// 登录失败
			data.Flag = true
			data.CurrentUser.LoginName = username
			return c.Render("login", data)
		}
	}
}

// 登出
func Logout() Handler {
	return func(c *fiber.Ctx) error {
		curUser = models.User{}
		return c.Render("login", fiber.Map{})
	}
}
