package main

import (
	"context"
	"log"
	"play-to-win-api/internal/config"
	"play-to-win-api/internal/delivery/http/handler"
	route "play-to-win-api/internal/delivery/http/routes"
	"play-to-win-api/internal/repository/mongodb"
	"play-to-win-api/internal/usecase"
	mongoClient "play-to-win-api/pkg/mongodb"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()

	db, err := mongoClient.NewClient(
		context.Background(),
		cfg.MongoDB.URI,
		cfg.MongoDB.Database,
	)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	categoryRepo := mongodb.NewCategoryRepository(db)

	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	e := echo.New()

	handlers := handler.NewHandlers(e, categoryUseCase)

	route.SetupRoutes(e, handlers)

	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
