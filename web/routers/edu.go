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
			CurrentUser: GetCurrentUser(c),
			Msg:         "",
			Flag:        false,
		}
		return c.Render("addEdu", data)
	}
}

// 创建电子证书信息
func AddEdu(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		currentUser := GetCurrentUser(c)
		if !currentUser.IsAdmin {
			return c.Status(fiber.StatusForbidden).SendString("无权限访问")
		}

		edu := service.ChaincodeInfo{
			Name:           c.FormValue("name"),
			SignDate:       c.FormValue("SignDate"),
			Nation:         c.FormValue("nation"),
			CopyRightID:    c.FormValue("CopyRightID"),
			Creator:        c.FormValue("Creator"),
			Holder:         c.FormValue("Holder"),
			EnrollDate:     c.FormValue("enrollDate"),
			NowState:       c.FormValue("NowState"),
			DoneDate:       c.FormValue("DoneDate"),
			Major:          c.FormValue("major"),
			QuaType:        c.FormValue("quaType"),
			Length:         c.FormValue("length"),
			Mode:           c.FormValue("mode"),
			Level:          c.FormValue("level"),
			CopyRightLevel: c.FormValue("CopyRightLevel"),
			CertNo:         c.FormValue("certNo"),
			Photo:          c.FormValue("photo"),
		}

		if edu.Name == "" || edu.CertNo == "" || edu.CopyRightID == "" {
			return c.Render("addEdu", fiber.Map{
				"CurrentUser": currentUser,
				"Msg":         "姓名、身份证号、证书编号不能为空",
				"Flag":        true,
			})
		}

		if _, err := s.SaveEdu(edu); err != nil {
			return c.Render("addEdu", fiber.Map{
				"CurrentUser": currentUser,
				"Msg":         err.Error(),
				"Flag":        true,
			})
		}

		c.Set("certNo", edu.CertNo)
		c.Set("name", edu.Name)
		return FindCertByNoAndName(s)(c)
	}
}