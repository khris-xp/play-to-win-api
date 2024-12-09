package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewResponse(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, Response{
		Status:  code >= 200 && code < 300,
		Message: message,
		Code:    code,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, code int, message string) error {
	return NewResponse(c, code, message, nil)
}
