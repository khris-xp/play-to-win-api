package handler

import (
	"net/http"
	"time"

	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
)

type CartItemHandler struct {
	BaseHandler
	cartItemUseCase domain.CartItemUseCase
}

func NewCartItemHandler(uc domain.CartItemUseCase) CartItemHandler {
	return CartItemHandler{
		BaseHandler:     BaseHandler{validator: validator.NewValidator()},
		cartItemUseCase: uc,
	}
}

func (h *CartItemHandler) Create(c echo.Context) error {
	var cartItem domain.CartItem
	if err := c.Bind(&cartItem); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	cartItem.CreatedAt = time.Now()
	cartItem.UpdatedAt = time.Now()

	if err := h.cartItemUseCase.Create(c.Request().Context(), &cartItem); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.CartItemCreatedSuccess, cartItem)
}

func (h *CartItemHandler) GetByCartID(c echo.Context) error {
	cartID := c.Param("cart_id")
	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartItemRetrievedSuccess, cartItems)
}

func (h *CartItemHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	cartItem, err := h.cartItemUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartItemRetrievedSuccess, cartItem)
}

func (h *CartItemHandler) GetAll(c echo.Context) error {
	cartItems, err := h.cartItemUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartItemRetrievedSuccess, cartItems)
}

func (h *CartItemHandler) Update(c echo.Context) error {
	var cartItem domain.CartItem
	if err := c.Bind(&cartItem); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.cartItemUseCase.Update(c.Request().Context(), &cartItem); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartItemUpdatedSuccess, cartItem)
}

func (h *CartItemHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.cartItemUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartItemDeletedSuccess, nil)
}
