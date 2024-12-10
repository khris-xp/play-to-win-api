package handler

import (
	"net/http"
	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/middleware"
	"play-to-win-api/internal/delivery/http/response"
	"play-to-win-api/internal/domain"
	"play-to-win-api/pkg/validator"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	validator   *validator.CustomValidator
	authUseCase domain.AuthUseCase
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func NewAuthHandler(uc domain.AuthUseCase, validator *validator.CustomValidator) AuthHandler {
	return AuthHandler{
		validator:   validator,
		authUseCase: uc,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidCredentials)
	}

	if err := h.validator.Validate(&user); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	tokens, err := h.authUseCase.Register(c.Request().Context(), &user)
	if err != nil {
		switch err {
		case domain.ErrUserAlreadyExists:
			return response.ErrorResponse(c, http.StatusConflict, constants.UserDuplicateError)
		default:
			return response.ErrorResponse(c, http.StatusInternalServerError, constants.RegistrationError)
		}
	}

	return response.NewResponse(c, http.StatusCreated, constants.RegistrationSuccess, tokens)
}

func (h *AuthHandler) Login(c echo.Context) error {

	var credentials struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.Bind(&credentials); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	if err := h.validator.Validate(&credentials); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	tokens, err := h.authUseCase.Login(c.Request().Context(), credentials.Email, credentials.Password)
	if err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidCredentials)
	}

	return response.NewResponse(c, http.StatusOK, constants.LoginSuccess, tokens)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var req struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, constants.InvalidRequestError)
	}

	tokens, err := h.authUseCase.RefreshToken(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidToken)
	}

	return response.NewResponse(c, http.StatusOK, constants.RefreshSuccess, tokens)
}

func (h *AuthHandler) GetProfile(c echo.Context) error {
	claims, ok := c.Get("user").(*middleware.Claims)
	if !ok {
		return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidUserClaims)
	}

	user, err := h.authUseCase.GetUserByEmail(c.Request().Context(), claims.Email)

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, constants.UserNotFoundError)
	}

	profile, err := h.authUseCase.GetProfile(c.Request().Context(), user.ID.Hex())

	if err != nil {
		if err == domain.ErrUserNotFound {
			return response.ErrorResponse(c, http.StatusNotFound, constants.UserNotFoundError)
		}
		return response.ErrorResponse(c, http.StatusInternalServerError, constants.ProfileError)
	}

	return response.NewResponse(c, http.StatusOK, constants.UserRetrievedSuccess, profile)
}
