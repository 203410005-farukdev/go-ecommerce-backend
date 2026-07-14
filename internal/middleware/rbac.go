package middleware

import (
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RBAC(svc *service.RBACService) fiber.Handler {
	return svc.Middleware()
}
