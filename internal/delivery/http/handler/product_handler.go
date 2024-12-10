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

type ProductHandler struct {
	BaseHandler
	productUseCase domain.ProductUseCase
}

func NewProductHandler(uc domain.ProductUseCase) ProductHandler {
	return ProductHandler{
		BaseHandler:    BaseHandler{validator: validator.NewValidator()},
		productUseCase: uc,
	}
}

func (h *ProductHandler) Create(c echo.Context) error {
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.productUseCase.Create(c.Request().Context(), &product); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.ProductCreatedSuccess, product)
}

func (h *ProductHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	product, err := h.productUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.ProductRetrievedSuccess, product)
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.productUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.ProductsRetrievedSuccess, products)
}

func (h *ProductHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.CategoryInvalidDataError)
	}

	product.ID = objectID

	if err := h.productUseCase.Update(c.Request().Context(), &product); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.ProductUpdatedSuccess, product)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.productUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.ProductDeletedSuccess, nil)
}
