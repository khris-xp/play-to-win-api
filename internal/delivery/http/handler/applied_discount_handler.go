package handler

import (
	"net/http"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DiscountHandler struct {
	BaseHandler
	cartItemUseCase        domain.CartItemUseCase
	appliedDiscountUseCase domain.AppliedDiscountUseCase
}

func NewDiscountHandler(cartItemUC domain.CartItemUseCase, appliedDiscountUC domain.AppliedDiscountUseCase) *DiscountHandler {
	return &DiscountHandler{
		cartItemUseCase:        cartItemUC,
		appliedDiscountUseCase: appliedDiscountUC,
	}
}

func (h *DiscountHandler) CalculateFixedAmount(c echo.Context) error {
	cartID := c.Param("cart_id")
	amount := c.QueryParam("amount")

	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	finalPrice, err := h.appliedDiscountUseCase.CalculateFixedAmountDiscount(c.Request().Context(), cartItems, parseFloat(amount))

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, "Discount calculated successfully", map[string]interface{}{
		"final_price": finalPrice,
	})
}

func (h *DiscountHandler) CalculatePercentage(c echo.Context) error {
	cartID := c.Param("cart_id")
	percentage := c.QueryParam("percentage")

	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	finalPrice, err := h.appliedDiscountUseCase.CalculatePercentageDiscount(c.Request().Context(), cartItems, parseFloat(percentage))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, "Discount calculated successfully", map[string]interface{}{
		"final_price": finalPrice,
	})
}

func (h *DiscountHandler) CalculateCategoryDiscount(c echo.Context) error {
	cartID := c.Param("cart_id")
	category := c.QueryParam("category")
	percentage := c.QueryParam("percentage")

	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	finalPrice, err := h.appliedDiscountUseCase.CalculateCategoryDiscount(c.Request().Context(), cartItems, category, parseFloat(percentage))

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, "Discount calculated successfully", map[string]interface{}{
		"final_price": finalPrice,
	})
}

func (h *DiscountHandler) CalculatePointsDiscount(c echo.Context) error {
	cartID := c.Param("cart_id")
	points := c.QueryParam("points")

	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	finalPrice, err := h.appliedDiscountUseCase.CalculatePointsDiscount(c.Request().Context(), cartItems, parseInt(points))

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, "Discount calculated successfully", map[string]interface{}{
		"final_price": finalPrice,
	})
}

func (h *DiscountHandler) CalculateSpecialDiscount(c echo.Context) error {
	cartID := c.Param("cart_id")
	threshold := c.QueryParam("threshold")
	discount := c.QueryParam("discount")

	cartItems, err := h.cartItemUseCase.GetByCartID(c.Request().Context(), cartID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	finalPrice, err := h.appliedDiscountUseCase.CalculateSpecialDiscount(c.Request().Context(), cartItems, parseFloat(threshold), parseFloat(discount))

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, "Discount calculated successfully", map[string]interface{}{
		"final_price": finalPrice,
	})
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
