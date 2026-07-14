package main

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/database"
	"github.com/Farukcoder/eCommerce-go/backend/internal/config"
	"github.com/Farukcoder/eCommerce-go/backend/internal/handler"
	"github.com/Farukcoder/eCommerce-go/backend/internal/middleware"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
	"github.com/Farukcoder/eCommerce-go/backend/internal/router"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitLogger()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get database handle", "error", err)
		return
	}
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	rbacRepo := repository.NewRBACRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	requestLogRepo := repository.NewRequestLogRepository(db)
	rbacService := service.NewRBACService(rbacRepo)
	if err := rbacService.Reload(context.Background()); err != nil {
		slog.Error("Failed to load RBAC cache", "error", err)
		return
	}
	rbacService.StartAutoReload(context.Background(), 5*time.Minute)
	authService := service.NewAuthService(userRepo, rbacRepo, refreshTokenRepo, cfg.JwtSecret, cfg.JwtRefreshSecret)
	authHandler := handler.NewAuthHandler(authService, rbacService)
	logsHandler := handler.NewLogsHandler(requestLogRepo)
	rbacHandler := handler.NewRBACHandler(rbacService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	subcategoryRepo := repository.NewSubcategoryRepository(db)
	subcategoryService := service.NewSubcategoryService(subcategoryRepo)
	subcategoryHandler := handler.NewSubcategoryHandler(subcategoryService)

	productRepo := repository.NewProductRepository(db)
	productVariantRepo := repository.NewProductVariantRepository(db)
	productImageRepo := repository.NewProductImageRepository(db)
	productService := service.NewProductService(productRepo, productVariantRepo, productImageRepo)
	productHandler := handler.NewProductHandler(productService)

	stockRepo := repository.NewStockRepository(db)
	stockService := service.NewStockService(stockRepo)
	stockHandler := handler.NewStockHandler(stockService)

	settingRepo := repository.NewSettingRepository(db)
	settingService := service.NewSettingService(settingRepo)
	settingHandler := handler.NewSettingHandler(settingService)
	homeService := service.NewHomeService(settingRepo, categoryRepo, productRepo)
	homeHandler := handler.NewHomeHandler(homeService)

	dashboardService := service.NewDashboardService(db)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	app := fiber.New()
	app.Use(middleware.CORS(cfg.AllowedOrigins))
	app.Use(middleware.SecurityHeaders(cfg.AppEnv))
	app.Use(middleware.RequestLogger(requestLogRepo))

	uploadsDir, err := filepath.Abs(filepath.Join("storage", "uploads"))
	if err != nil {
		slog.Error("Failed to resolve uploads path", "error", err)
		return
	}
	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		slog.Error("Failed to create uploads directory", "error", err)
		return
	}

	app.Static("/uploads", uploadsDir)
	router.Setup(app, authHandler, logsHandler, rbacHandler, categoryHandler, subcategoryHandler, productHandler, stockHandler, settingHandler, homeHandler, dashboardHandler, rbacService, cfg)

	slog.Info("Server starting", "port", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		slog.Error("Server error", "error", err)
	}
}
