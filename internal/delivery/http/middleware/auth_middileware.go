package middleware

import (
	"fmt"
	"net/http"
	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/delivery/http/response"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	accessSecret string
}

func NewAuthMiddleware(accessSecret string) *AuthMiddleware {
	if accessSecret == "" {
		panic("access secret cannot be empty")
	}
	return &AuthMiddleware{
		accessSecret: accessSecret,
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (m *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return response.ErrorResponse(c, http.StatusUnauthorized, constants.MissingAuthHeader)
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidAuthHeader)
		}

		token, err := jwt.ParseWithClaims(tokenParts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.accessSecret), nil
		})

		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidToken+err.Error())
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			return response.ErrorResponse(c, http.StatusUnauthorized, constants.InvalidUserClaims)
		}
		c.Set("user", claims)
		return next(c)
	}
}

func RequireRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*Claims)
			if !ok {
				return response.ErrorResponse(c, http.StatusUnauthorized, constants.NotAuthenticated)
			}

			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}

			return response.ErrorResponse(c, http.StatusForbidden, constants.InsufficientPerms)
		}
	}
}
