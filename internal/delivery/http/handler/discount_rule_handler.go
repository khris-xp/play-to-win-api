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

type DiscountRuleHandler struct {
	BaseHandler
	discountRuleUseCase domain.DiscountRuleUseCase
}

func NewDiscountRuleHandler(uc domain.DiscountRuleUseCase) DiscountRuleHandler {
	return DiscountRuleHandler{
		BaseHandler:         BaseHandler{validator: validator.NewValidator()},
		discountRuleUseCase: uc,
	}
}

func (h *DiscountRuleHandler) Create(c echo.Context) error {
	var discountRule domain.DiscountRule
	if err := c.Bind(&discountRule); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.discountRuleUseCase.Create(c.Request().Context(), &discountRule); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.DiscountRuleCreatedSuccess, discountRule)
}

func (h *DiscountRuleHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	discountRule, err := h.discountRuleUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.DiscountRuleRetrievedSuccess, discountRule)
}

func (h *DiscountRuleHandler) GetAll(c echo.Context) error {
	discountRules, err := h.discountRuleUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.DiscountRulesRetrievedSuccess, discountRules)
}

func (h *DiscountRuleHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var discountRule domain.DiscountRule
	if err := c.Bind(&discountRule); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	discountRule.ID, _ = primitive.ObjectIDFromHex(id)
	if err := h.discountRuleUseCase.Update(c.Request().Context(), &discountRule); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.DiscountRuleUpdatedSuccess, discountRule)
}

func (h *DiscountRuleHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.discountRuleUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.DiscountRuleDeletedSuccess, nil)
}
