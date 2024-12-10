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
	categories.GET("", handlers.Category.GetAll)
	categories.GET("/:id", handlers.Category.GetByID)

	protectedCategories := categories.Group("")
	protectedCategories.Use(handlers.AuthMW.Authenticate)

	adminCategories := protectedCategories.Group("")
	adminCategories.Use(middleware.RequireRole("admin"))
	adminCategories.POST("", handlers.Category.Create)
	adminCategories.PUT("/:id", handlers.Category.Update)
	adminCategories.DELETE("/:id", handlers.Category.Delete)

	auth := v1.Group("/auth")
	auth.POST("/register", handlers.Auth.Register)
	auth.POST("/login", handlers.Auth.Login)
	auth.POST("/refresh", handlers.Auth.RefreshToken)

	user := v1.Group("/user")
	user.Use(handlers.AuthMW.Authenticate)
	user.GET("/profile", handlers.Auth.GetProfile)

	products := v1.Group("/products")
	products.GET("", handlers.Product.GetAll)
	products.GET("/:id", handlers.Product.GetByID)

	protectedProducts := products.Group("")
	protectedProducts.Use(handlers.AuthMW.Authenticate)

	adminProducts := protectedProducts.Group("")
	adminProducts.Use(middleware.RequireRole("admin"))
	adminProducts.POST("", handlers.Product.Create)
	adminProducts.PUT("/:id", handlers.Product.Update)
	adminProducts.DELETE("/:id", handlers.Product.Delete)
}
