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
			CurrentUser: GetCurrentUser(c),
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
			CurrentUser: GetCurrentUser(c),
			Msg:         "",
			Flag:        false,
		}
		return c.Render("query2", data)
	}
}

// 根据身份证号和名称查询
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

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
			History     bool
		}{
			Edu:         service.ChaincodeInfo{},
			CurrentUser: GetCurrentUser(c),
			Msg:         "",
			Flag:        false,
			History:     false,
		}

		if certNo == "" || name == "" {
			data.Msg = "身份证号和姓名不能为空"
			data.Flag = true
			return c.Render("queryResult", data)
		}

		result, err := s.FindEduByCertNoAndName(certNo, name)
		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
			return c.Render("queryResult", data)
		}

		if len(result) > 0 {
			if err := json.Unmarshal(result, &data.Edu); err != nil {
				data.Msg = "查询结果解析失败: " + err.Error()
				data.Flag = true
				return c.Render("queryResult", data)
			}
		}

		fmt.Println("身份证号与姓名信息查询成功：")
		fmt.Println(data.Edu)

		return c.Render("queryResult", data)
	}
}

func FindCertByID(s *service.ServiceSetup) Handler {
	return func(c *fiber.Ctx) error {
		copyRightID := c.FormValue("CopyRightID")
		if copyRightID == "" {
			copyRightID = c.Get("CopyRightID")
		}

		data := &struct {
			Edu         service.ChaincodeInfo
			CurrentUser models.User
			Msg         string
			Flag        bool
			History     bool
		}{
			Edu:         service.ChaincodeInfo{},
			CurrentUser: GetCurrentUser(c),
			Msg:         "",
			Flag:        false,
			History:     true,
		}

		if copyRightID == "" {
			data.Msg = "证书编号不能为空"
			data.Flag = true
			return c.Render("queryResult", data)
		}

		result, err := s.FindEduInfoByCopyRightID(copyRightID)
		if err != nil {
			data.Msg = err.Error()
			data.Flag = true
			return c.Render("queryResult", data)
		}

		if len(result) > 0 {
			if err := json.Unmarshal(result, &data.Edu); err != nil {
				data.Msg = "查询结果解析失败: " + err.Error()
				data.Flag = true
				return c.Render("queryResult", data)
			}
		}

		return c.Render("queryResult", data)
	}
}