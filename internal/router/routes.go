package router

import (
	"github.com/Farukcoder/eCommerce-go/backend/internal/config"
	"github.com/Farukcoder/eCommerce-go/backend/internal/handler"
	"github.com/Farukcoder/eCommerce-go/backend/internal/middleware"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler, logsHandler *handler.LogsHandler, rbacHandler *handler.RBACHandler, categoryHandler *handler.CategoryHandler, subcategoryHandler *handler.SubcategoryHandler, productHandler *handler.ProductHandler, stockHandler *handler.StockHandler, settingsHandler *handler.SettingHandler, homeHandler *handler.HomeHandler, dashboardHandler *handler.DashboardHandler, rbacService *service.RBACService, cfg *config.Config) {
	// Update auth handler with rbac service if not already set
	// This is handled in main.go during initialization
	api := app.Group("/api/v1")
	auth := api.Group("/auth", middleware.AuthRateLimiter())
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/logout", authHandler.Logout)

	api.Get("/home", homeHandler.GetHomePage)
	api.Get("/settings", settingsHandler.GetSettings)
	api.Get("/categories", categoryHandler.ListCategories)
	api.Get("/categories/:id", categoryHandler.GetCategory)
	api.Get("/subcategories", subcategoryHandler.ListSubcategories)
	api.Get("/subcategories/:id", subcategoryHandler.GetSubcategory)
	api.Get("/products", productHandler.ListProducts)
	api.Get("/products/:id", productHandler.GetProduct)
	api.Get("/products/:id/variants", productHandler.ListProductVariants)

	protected := api.Group("", jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(cfg.JwtSecret)},
		ContextKey:   "user",
		ErrorHandler: middleware.ErrorHandler,
		SuccessHandler: func(c *fiber.Ctx) error {
			token, ok := c.Locals("user").(*jwt.Token)
			if ok && token != nil && token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					c.Locals("jwt_claims", claims)
				}
			}
			return c.Next()
		},
	}))
	protected.Use(rbacService.Middleware())
	protected.Get("/me", authHandler.Me)
	protected.Get("/users", authHandler.GetAllUsers)
	protected.Get("/logs", logsHandler.GetAll)
	protected.Post("/categories", categoryHandler.CreateCategory)
	protected.Put("/categories/:id", categoryHandler.UpdateCategory)
	protected.Delete("/categories/:id", categoryHandler.DeleteCategory)
	protected.Post("/subcategories", subcategoryHandler.CreateSubcategory)
	protected.Put("/subcategories/:id", subcategoryHandler.UpdateSubcategory)
	protected.Delete("/subcategories/:id", subcategoryHandler.DeleteSubcategory)

	protected.Post("/products", productHandler.CreateProduct)
	protected.Put("/products/:id", productHandler.UpdateProduct)
	protected.Delete("/products/:id", productHandler.DeleteProduct)

	protected.Post("/products/:id/variants", productHandler.CreateProductVariant)
	protected.Put("/products/:id/variants/:variant_id", productHandler.UpdateProductVariant)
	protected.Delete("/products/:id/variants/:variant_id", productHandler.DeleteProductVariant)

	protected.Get("/stock", stockHandler.ListStock)
	protected.Get("/stock/:id", stockHandler.GetStock)
	protected.Put("/stock/:id", stockHandler.UpdateStock)
	protected.Get("/stock-movements", stockHandler.ListStockMovements)
	protected.Post("/stock-movements", stockHandler.CreateStockMovement)

	protected.Put("/settings", settingsHandler.UpdateSettings)
	protected.Get("/dashboard/stats", dashboardHandler.GetStats)

	protected.Get("/roles", rbacHandler.ListRoles)
	protected.Post("/roles", rbacHandler.CreateRole)
	protected.Put("/roles/:id", rbacHandler.UpdateRole)
	protected.Delete("/roles/:id", rbacHandler.DeleteRole)
	protected.Get("/roles/:id/permissions", rbacHandler.GetRolePermissions)
	protected.Get("/permissions", rbacHandler.ListPermissions)
	protected.Post("/permissions", rbacHandler.CreatePermission)
	protected.Put("/permissions/:id", rbacHandler.UpdatePermission)
	protected.Delete("/permissions/:id", rbacHandler.DeletePermission)
	protected.Post("/roles/:id/permissions", rbacHandler.AssignPermissionToRole)
	protected.Delete("/roles/:id/permissions/:permission_id", rbacHandler.RevokePermissionFromRole)
	protected.Patch("/users/:id/role", rbacHandler.AssignRoleToUser)
	protected.Get("/me/permissions", rbacHandler.GetUserPermissions)
}
