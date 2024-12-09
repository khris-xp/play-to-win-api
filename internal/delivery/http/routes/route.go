package route

import (
	"play-to-win-api/internal/delivery/http/handler"
	"play-to-win-api/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, handlers *handler.Handlers) {
	middleware.SetupMiddleware(e)

	v1 := e.Group("/api/v1")

	categories := v1.Group("/categories")
	categories.POST("", handlers.Category.Create)
	categories.GET("", handlers.Category.GetAll)
	categories.GET("/:id", handlers.Category.GetByID)
	categories.PUT("/:id", handlers.Category.Update)
	categories.DELETE("/:id", handlers.Category.Delete)
}
