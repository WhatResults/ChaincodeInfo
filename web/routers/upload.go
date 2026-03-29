package routers

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// 上传文件
func Upload() Handler {
	return func(c *fiber.Ctx) error {
		start := "{"
		content := ""
		end := "}"

		fh, err := c.FormFile("file")
		if err != nil {
			content = "\"error\":1,\"result\":{\"msg\":\"指定了无效的文件\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}

		file, err := fh.Open()
		if err != nil {
			content = "\"error\":1,\"result\":{\"msg\":\"文件打开失败\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			content = "\"error\":1,\"result\":{\"msg\":\"无法读取文件内容\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}

		filetype := http.DetectContentType(fileBytes)
		//log.Println("filetype = " + filetype)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			content = "\"error\":1,\"result\":{\"msg\":\"文件类型错误\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}

		fileName := randToken(12)                           // 指定文件名
		fileEndings, err := mime.ExtensionsByType(filetype) // 获取文件扩展名
		//log.Println("fileEndings = " + fileEndings[0])
		// 指定文件存储路径
		newPath := filepath.Join("web", "static", "photo", fileName+fileEndings[0])
		//fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

		newFile, err := os.Create(newPath)
		if err != nil {
			log.Println("创建文件失败：" + err.Error())
			content = "\"error\":1,\"result\":{\"msg\":\"创建文件失败\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}
		defer newFile.Close()

		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			log.Println("写入文件失败：" + err.Error())
			content = "\"error\":1,\"result\":{\"msg\":\"保存文件内容失败\",\"path\":\"\"}"
			c.Write([]byte(start + content + end))
			return err
		}

		path := "/static/photo/" + fileName + fileEndings[0]
		content = "\"error\":0,\"result\":{\"fileType\":\"image/png\",\"path\":\"" + path + "\"}"
		c.Write([]byte(start + content + end))
		return err
	}
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
