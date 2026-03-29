package routers

import (
	"ChaincodeInfo/service"
	"ChaincodeInfo/web/models"

	"github.com/gofiber/fiber/v2"
)

// 创建电子证书信息页面
func AddEduView() Handler {
	return func(c *fiber.Ctx) error {
		data := &struct {
			CurrentUser models.User
			Msg         string
			Flag        bool
		}{
			CurrentUser: curUser,
			Msg:         "",
			Flag:        false,
		}
		return c.Render("addEdu", data)
	}
}

// 创建电子证书信息
func AddEdu(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		edu := service.ChaincodeInfo{
			Name:           c.FormValue("name"),
			SignDate:         c.FormValue("SignDate"),
			Nation:         c.FormValue("nation"),
			CopyRightID:       c.FormValue("CopyRightID"),
			Creator:          c.FormValue("Creator"),
			Holder:       c.FormValue("Holder"),
			EnrollDate:     c.FormValue("enrollDate"),
			NowState: c.FormValue("NowState"),
			DoneDate:     c.FormValue("DoneDate"),
			Major:          c.FormValue("major"),
			QuaType:        c.FormValue("quaType"),
			Length:         c.FormValue("length"),
			Mode:           c.FormValue("mode"),
			Level:          c.FormValue("level"),
			CopyRightLevel:     c.FormValue("CopyRightLevel"),
			CertNo:         c.FormValue("certNo"),
			Photo:          c.FormValue("photo"),
		}

		// 存储电子证书信息
		s.SaveEdu(edu)

		c.Set("certNo", edu.CertNo)
		c.Set("name", edu.Name)

		return FindCertByNoAndName(s)(c)
	}
}
