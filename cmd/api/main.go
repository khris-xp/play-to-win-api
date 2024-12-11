package main

import (
	"context"
	"log"
	"play-to-win-api/internal/config"
	"play-to-win-api/internal/delivery/http/handler"
	"play-to-win-api/internal/delivery/http/middleware"
	route "play-to-win-api/internal/delivery/http/routes"
	"play-to-win-api/internal/repository/mongodb"
	"play-to-win-api/internal/usecase"
	mongoClient "play-to-win-api/pkg/mongodb"
	"play-to-win-api/pkg/validator"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadConfig()

	if cfg.JWT.AccessSecret == "" || cfg.JWT.RefreshSecret == "" {
		log.Fatal("JWT secrets must be configured")
	}

	db, err := mongoClient.NewClient(
		context.Background(),
		cfg.MongoDB.URI,
		cfg.MongoDB.Database,
	)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	categoryRepo := mongodb.NewCategoryRepository(db)
	userRepo := mongodb.NewUserRepository(db)
	productRepo := mongodb.NewProductRepository(db)
	campaignRepo := mongodb.NewCampaignRepository(db)

	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	authUseCase := usecase.NewAuthUseCase(
		userRepo,
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		24*time.Hour,
		7*24*time.Hour,
	)
	productUseCase := usecase.NewProductUseCase(productRepo)
	campaignUseCase := usecase.NewCampaignUseCase(campaignRepo)

	e := echo.New()

	v := validator.NewValidator()

	authMiddleware := middleware.NewAuthMiddleware(cfg.JWT.AccessSecret)

	handlers := &handler.Handlers{
		Category: handler.NewCategoryHandler(categoryUseCase),
		Auth:     handler.NewAuthHandler(authUseCase, v),
		AuthMW:   authMiddleware,
		Product:  handler.NewProductHandler(productUseCase),
		Campaign: handler.NewCampaignHandler(campaignUseCase),
	}

	route.SetupRoutes(e, handlers)

	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
