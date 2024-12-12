package handler

import (
	"net/http"
	"time"

	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/middleware"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	BaseHandler
	cartUseCase  domain.CartUseCase
	authUserCase domain.AuthUseCase
}

func NewCartHandler(uc domain.CartUseCase, ac domain.AuthUseCase) CartHandler {
	return CartHandler{
		BaseHandler:  BaseHandler{validator: validator.NewValidator()},
		cartUseCase:  uc,
		authUserCase: ac,
	}
}

func (h *CartHandler) Create(c echo.Context) error {
	claims, ok := c.Get("user").(*middleware.Claims)
	if !ok {
		return response.ErrorResponse(c, http.StatusInternalServerError, constants.InternalServerError)
	}

	user, err := h.authUserCase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var cart domain.Cart
	if err := c.Bind(&cart); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	cart.User = *user
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()

	if err := h.cartUseCase.Create(c.Request().Context(), &cart); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.CartCreatedSuccess, cart)
}

func (h *CartHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	cart, err := h.cartUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartRetrievedSuccess, cart)
}

func (h *CartHandler) GetByUserID(c echo.Context) error {
	claims, ok := c.Get("user").(*middleware.Claims)

	user, err := h.authUserCase.GetUserByEmail(c.Request().Context(), claims.Email)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !ok {
		return response.ErrorResponse(c, http.StatusInternalServerError, constants.InternalServerError)
	}

	cart, err := h.cartUseCase.GetByUserID(c.Request().Context(), user.ID.Hex())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartRetrievedSuccess, cart)
}

func (h *CartHandler) GetAll(c echo.Context) error {
	carts, err := h.cartUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartRetrievedSuccess, carts)
}

func (h *CartHandler) Update(c echo.Context) error {
	var cart domain.Cart
	if err := c.Bind(&cart); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.cartUseCase.Update(c.Request().Context(), &cart); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartUpdatedSuccess, cart)
}

func (h *CartHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.cartUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CartDeletedSuccess, nil)
}
