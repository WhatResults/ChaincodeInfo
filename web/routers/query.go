package routers

import (
	"ChaincodeInfo/service"
	"ChaincodeInfo/web/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// 查询页面
func QueryPage() Handler {
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

		return c.Render("query", data)
	}
}

func QueryPage2() Handler {
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

		return c.Render("query2", data)
	}
}

// 根据电子证书信息序号和名称查询
func FindCertByNoAndName(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		certNo := c.FormValue("certNo")
		if certNo == "" {
			certNo = c.Get("certNo")
		}

		name := c.FormValue("name")
		if name == "" {
			name = c.Get("name")
		}

		result, err := s.FindEduByCertNoAndName(certNo, name)
		var edu = service.ChaincodeInfo{}
		json.Unmarshal(result, &edu)

		fmt.Println("身份证号与姓名信息成功：")
		fmt.Println(edu)

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
			History     bool
		}{
			Edu:         edu,
			CurrentUser: curUser,
			Msg:         "",
			Flag:        false,
			History:     false,
		}

		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
		}

		return c.Render("queryResult", data)
	}
}

func FindCertByID(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		CopyRightID := c.FormValue("CopyRightID")
		result, err := s.FindEduInfoByCopyRightID(CopyRightID)
		var edu = service.ChaincodeInfo{}
		json.Unmarshal(result, &edu)

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
			History     bool
		}{
			Edu:         edu,
			CurrentUser: curUser,
			Msg:         "",
			Flag:        false,
			History:     true,
		}

		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
		}

		return c.Render("queryResult", data)
	}
}
