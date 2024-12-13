package handler

import (
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"play-to-win-api/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Category CategoryHandler
	Auth     AuthHandler
	AuthMW   *middleware.AuthMiddleware
	Product  ProductHandler
	Campaign CampaignHandler
	Cart     CartHandler
	CartItem CartItemHandler
	DiscountRule DiscountRuleHandler
}

func NewHandlers(e *echo.Echo, categoryUseCase domain.CategoryUseCase, authUseCase domain.AuthUseCase, productUseCase domain.ProductUseCase, campaignUseCase domain.CampaignUseCase, cartUseCase domain.CartUseCase, cartItemUseCase domain.CartItemUseCase) *Handlers {
	validator := validator.NewValidator()
	return &Handlers{
		Category: NewCategoryHandler(categoryUseCase),
		Auth:     NewAuthHandler(authUseCase, validator),
		Product:  NewProductHandler(productUseCase),
		Campaign: NewCampaignHandler(campaignUseCase),
		Cart:     NewCartHandler(cartUseCase, authUseCase),
		CartItem: NewCartItemHandler(cartItemUseCase),
	}
}

type BaseHandler struct {
	validator *validator.CustomValidator
}
