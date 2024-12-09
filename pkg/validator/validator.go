package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	v.RegisterValidation("alpha_space", validateAlphaSpace)
	v.RegisterValidation("mongodb_objectid", validateMongoDBObjectID)

	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validateAlphaSpace(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match(`^[a-zA-Z\s]*$`, []byte(fl.Field().String()))
	return matched
}

func validateMongoDBObjectID(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match(`^[0-9a-fA-F]{24}$`, []byte(fl.Field().String()))
	return matched
}

func (cv *CustomValidator) RegisterValidation(tag string, fn validator.Func) error {
	return cv.validator.RegisterValidation(tag, fn)
}
