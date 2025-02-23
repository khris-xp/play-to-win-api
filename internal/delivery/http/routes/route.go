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

	campaigns := v1.Group("/campaigns")
	campaigns.GET("", handlers.Campaign.GetAll)
	campaigns.GET("/:id", handlers.Campaign.GetByID)

	protectedCampaigns := campaigns.Group("")
	protectedCampaigns.Use(handlers.AuthMW.Authenticate)

	adminCampaigns := protectedCampaigns.Group("")
	adminCampaigns.Use(middleware.RequireRole("admin"))
	adminCampaigns.POST("", handlers.Campaign.Create)
	adminCampaigns.PUT("/:id", handlers.Campaign.Update)
	adminCampaigns.DELETE("/:id", handlers.Campaign.Delete)

	carts := v1.Group("/carts")

	protectedCart := carts.Group("")
	protectedCart.Use(handlers.AuthMW.Authenticate)

	protectedCart.GET("", handlers.Cart.GetByUserID)
	protectedCart.GET("/:id", handlers.Cart.GetByID)
	protectedCart.POST("", handlers.Cart.Create)
	protectedCart.PUT("/:id", handlers.Cart.Update)
	protectedCart.DELETE("/:id", handlers.Cart.Delete)

	cartItems := v1.Group("/cart-items")

	protectedCartItems := cartItems.Group("")
	protectedCartItems.Use(handlers.AuthMW.Authenticate)

	protectedCartItems.GET("/:cart_id", handlers.CartItem.GetByCartID)
	protectedCartItems.POST("", handlers.CartItem.Create)
	protectedCartItems.PUT("/:id", handlers.CartItem.Update)
	protectedCartItems.DELETE("/:id", handlers.CartItem.Delete)

	adminCartItems := protectedCartItems.Group("")
	adminCartItems.Use(middleware.RequireRole("admin"))
	adminCartItems.GET("", handlers.CartItem.GetAll)

	discountRule := v1.Group("/discount-rules")
	discountRule.GET("", handlers.DiscountRule.GetAll)
	discountRule.GET("/:id", handlers.DiscountRule.GetByID)

	protectedDiscountRule := discountRule.Group("")
	protectedDiscountRule.Use(handlers.AuthMW.Authenticate)

	adminDiscountRule := protectedDiscountRule.Group("")
	adminDiscountRule.Use(middleware.RequireRole("admin"))
	adminDiscountRule.POST("", handlers.DiscountRule.Create)
	adminDiscountRule.PUT("/:id", handlers.DiscountRule.Update)
	adminDiscountRule.DELETE("/:id", handlers.DiscountRule.Delete)

	discounts := v1.Group("/discounts")

	discounts.GET("/fixed-amount/:cart_id", handlers.Discount.CalculateFixedAmount)
	discounts.GET("/percentage/:cart_id", handlers.Discount.CalculatePercentage)
	discounts.GET("/category/:cart_id", handlers.Discount.CalculateCategoryDiscount)
	discounts.GET("/points/:cart_id", handlers.Discount.CalculatePointsDiscount)
	discounts.GET("/special/:cart_id", handlers.Discount.CalculateSpecialDiscount)
}
