package handler

import (
	"net/http"
	"time"

	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignHandler struct {
	BaseHandler
	campaignUseCase domain.CampaignUseCase
}

func NewCampaignHandler(uc domain.CampaignUseCase) CampaignHandler {
	return CampaignHandler{
		BaseHandler:     BaseHandler{validator: validator.NewValidator()},
		campaignUseCase: uc,
	}
}

func (h *CampaignHandler) Create(c echo.Context) error {
	var campaign domain.Campaign
	if err := c.Bind(&campaign); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	campaign.IsActive = true
	campaign.StartDate = time.Now()
	campaign.EndDate = time.Now().AddDate(0, 0, 7)

	if err := h.campaignUseCase.Create(c.Request().Context(), &campaign); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusCreated, constants.CampaignCreatedSuccess, campaign)
}

func (h *CampaignHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	campaign, err := h.campaignUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CampaignRetrievedSuccess, campaign)
}

func (h *CampaignHandler) GetAll(c echo.Context) error {
	campaigns, err := h.campaignUseCase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CampaignsRetrievedSuccess, campaigns)
}

func (h *CampaignHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var campaign domain.Campaign
	if err := c.Bind(&campaign); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.CampaignInvalidDataError)
	}

	campaign.ID = objectID
	if err := h.campaignUseCase.Update(c.Request().Context(), &campaign); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CampaignUpdatedSuccess, campaign)
}

func (h *CampaignHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.campaignUseCase.Delete(c.Request().Context(), id); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.NewResponse(c, http.StatusOK, constants.CampaignDeletedSuccess, nil)
}
