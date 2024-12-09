package handler

import (
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Category CategoryHandler
}

func NewHandlers(e *echo.Echo, categoryUseCase domain.CategoryUseCase) *Handlers {
	return &Handlers{
		Category: NewCategoryHandler(categoryUseCase),
	}
}

type BaseHandler struct {
	validator *validator.CustomValidator
}
