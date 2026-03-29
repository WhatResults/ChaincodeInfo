package routers

import (
	"ChaincodeInfo/service"
	"ChaincodeInfo/web/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// 修改页面
func ModifyView(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		// 身份证号与姓名信息
		certNo := c.FormValue("certNo")
		name := c.FormValue("name")
		result, err := s.FindEduByCertNoAndName(certNo, name)

		var edu = service.ChaincodeInfo{}
		json.Unmarshal(result, &edu)

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
		}{
			Edu:         edu,
			CurrentUser: curUser,
			Flag:        true,
			Msg:         "",
		}

		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
		}

		return c.Render("modify", data)
	}
}

func Modify(s *service.ServiceSetup) Handler {
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

		s.ModifyEdu(edu)

		c.Set("certNo", edu.CertNo)
		c.Set("name", edu.Name)

		return FindCertByNoAndName(s)(c)
	}
}
