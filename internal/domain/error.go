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

	ErrCampaignNotFound      = errors.New("campaign not found")
	ErrCampaignAlreadyExists = errors.New("campaign already exists")
	ErrInvalidCampaignID     = errors.New("invalid campaign ID")
	ErrInvalidCampaignData   = errors.New("invalid campaign data")

	ErrCartNotFound  = errors.New("cart not found")
	ErrInvalidCartID = errors.New("invalid cart ID")

	ErrCartItemNotFound  = errors.New("cart item not found")
	ErrInvalidCartItemID = errors.New("invalid cart item ID")

	ErrInvalidDiscountRuleID = errors.New("invalid discount rule ID")
	ErrDiscountRuleNotFound  = errors.New("discount rule not found")
)
