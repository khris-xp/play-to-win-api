package domain

import (
	"context"
)

type AppliedDiscount struct {
	Category   string  `json:"category"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
	Points     int     `json:"points"`
}

type AppliedDiscountUseCase interface {
	CalculateFixedAmountDiscount(ctx context.Context, cartItems []CartItem, amount float64) (float64, error)
	CalculatePercentageDiscount(ctx context.Context, cartItems []CartItem, percentage float64) (float64, error)
	CalculateCategoryDiscount(ctx context.Context, cartItems []CartItem, category string, percentage float64) (float64, error)
	CalculatePointsDiscount(ctx context.Context, cartItems []CartItem, points int) (float64, error)
	CalculateSpecialDiscount(ctx context.Context, cartItems []CartItem, threshold, discount float64) (float64, error)
}
