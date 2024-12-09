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
)
