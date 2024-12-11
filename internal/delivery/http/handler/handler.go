package handler

import (
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"play-to-win-api/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Category  CategoryHandler
	Auth      AuthHandler
	AuthMW    *middleware.AuthMiddleware
	Product   ProductHandler
	Campaign CampaignHandler
}

func NewHandlers(e *echo.Echo, categoryUseCase domain.CategoryUseCase, authUseCase domain.AuthUseCase, productUseCase domain.ProductUseCase) *Handlers {
	validator := validator.NewValidator()
	return &Handlers{
		Category: NewCategoryHandler(categoryUseCase),
		Auth:     NewAuthHandler(authUseCase, validator),
		Product:  NewProductHandler(productUseCase),
	}
}

type BaseHandler struct {
	validator *validator.CustomValidator
}
