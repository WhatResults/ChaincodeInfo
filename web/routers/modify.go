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
		certNo := c.FormValue("certNo")
		if certNo == "" {
			certNo = c.Query("certNo")
		}

		name := c.FormValue("name")
		if name == "" {
			name = c.Query("name")
		}

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
		}{
			Edu:         service.ChaincodeInfo{},
			CurrentUser: GetCurrentUser(c),
			Flag:        false,
			Msg:         "",
		}

		if certNo == "" || name == "" {
			data.Msg = "身份证号和姓名不能为空"
			data.Flag = true
			return c.Render("modify", data)
		}

		result, err := s.FindEduByCertNoAndName(certNo, name)
		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
			return c.Render("modify", data)
		}

		if len(result) > 0 {
			if err := json.Unmarshal(result, &data.Edu); err != nil {
				data.Msg = "查询结果解析失败: " + err.Error()
				data.Flag = true
				return c.Render("modify", data)
			}
		}

		return c.Render("modify", data)
	}
}

func Modify(s *service.ServiceSetup) Handler {
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
			return c.Render("modify", fiber.Map{
				"Edu":         edu,
				"CurrentUser": currentUser,
				"Msg":         "姓名、身份证号、证书编号不能为空",
				"Flag":        true,
			})
		}

		if _, err := s.ModifyEdu(edu); err != nil {
			return c.Render("modify", fiber.Map{
				"Edu":         edu,
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
