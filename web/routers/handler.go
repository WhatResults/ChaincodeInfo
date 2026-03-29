package routers

import "github.com/gofiber/fiber/v2"

// 请求处理器结构定义
type Handler = func(*fiber.Ctx) error
