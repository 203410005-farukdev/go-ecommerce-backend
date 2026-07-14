package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Farukcoder/eCommerce-go/backend/database"
	"github.com/Farukcoder/eCommerce-go/backend/internal/config"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
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
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	rbacRepo := repository.NewRBACRepository(db)
	authService := service.NewAuthService(userRepo, rbacRepo, refreshTokenRepo, cfg.JwtSecret, cfg.JwtRefreshSecret)

	resp, err := authService.Login(context.Background(), dto.LoginInput{
		EmailOrPhone: "admin@example.com",
		Password:     "Password123!",
	})
	if err != nil {
		slog.Error("Login failed", "error", err)
		return
	}

	fmt.Printf("login succeeded; access_token=%s refresh_token=%s\n", resp.AccessToken, resp.RefreshToken)
}
