package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("requested resource not found")
	ErrConflict            = errors.New("resource already exists")
	ErrBadRequest          = errors.New("invalid request")
	ErrUnauthorized        = errors.New("unauthorized access")

	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryAlreadyExists = errors.New("category already exists")
	ErrInvalidCategoryID     = errors.New("invalid category ID")

	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrInvalidUserData   = errors.New("invalid user data")

	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrInvalidProductID     = errors.New("invalid product ID")
	ErrInvalidProductData   = errors.New("invalid product data")
)
