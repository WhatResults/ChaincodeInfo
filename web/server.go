package web

import (
	"ChaincodeInfo/service"
	"ChaincodeInfo/web/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// web服务
type Server struct {
	// Setup服务
	service *service.ServiceSetup
	// 服务应用实例
	app *fiber.App
}

// 创建web服务端
func NewServer(ss *service.ServiceSetup) *Server {
	engine := html.New("web/views", ".html")
	return &Server{
		service: ss,
		app:     fiber.New(fiber.Config{Views: engine}),
	}
}

// 创建注册路由
func (s *Server) registerRoutes() {
	s.app.Static("/static/", "web/static")
	s.app.Get("/", routers.LoginView())
	s.app.Post("/login", routers.Login())
	s.app.Get("/logout", routers.Logout())
	s.app.Get("/index", routers.Index())
	s.app.Get("/help", routers.Help())
	s.app.Get("/addEdu", routers.AddEduView())
	s.app.Post("/addEdu", routers.AddEdu(s.service))
	s.app.Get("/queryPage", routers.QueryPage())
	s.app.Post("/query", routers.FindCertByNoAndName(s.service))
	s.app.Get("/queryPage2", routers.QueryPage2())
	s.app.Post("/query2", routers.FindCertByID(s.service))
	s.app.Get("/modify", routers.ModifyView(s.service))
	s.app.Post("/modify", routers.Modify(s.service))
	s.app.Post("/upload", routers.Upload())
}

// 启动服务端
func (s *Server) Run() error {
	// 注册路由
	s.registerRoutes()
	// 启动服务
	return s.app.Listen(":3000")
}
