package handler

import (
	"net/http"

	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryHandler struct {
	BaseHandler
	categoryUseCase domain.CategoryUseCase
}

func NewCategoryHandler(uc domain.CategoryUseCase) CategoryHandler {
	return CategoryHandler{
		BaseHandler:     BaseHandler{validator: validator.NewValidator()},
		categoryUseCase: uc,
	}
}
func (h *CategoryHandler) Create(c echo.Context) error {
	var category domain.Category
	if err := c.Bind(&category); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.categoryUseCase.Create(c.Request().Context(), &category); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.CategoryCreatedSuccess, category)
}

func (h *CategoryHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	category, err := h.categoryUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CategoryRetrievedSuccess, category)
}

func (h *CategoryHandler) GetAll(c echo.Context) error {
	categories, err := h.categoryUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CategoriesRetrievedSuccess, categories)
}

func (h *CategoryHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var category domain.Category
	if err := c.Bind(&category); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.CategoryInvalidDataError)
	}

	category.ID = objectID
	if err := h.categoryUseCase.Update(c.Request().Context(), &category); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CategoryUpdatedSuccess, category)
}

func (h *CategoryHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.categoryUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CategoryDeletedSuccess, nil)
}
