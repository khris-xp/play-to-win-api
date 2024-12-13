package usecase

import (
	"context"
	"fmt"
	"math"
	"play-to-win-api/internal/domain"
)

type appliedDiscountUseCase struct{}

func NewAppliedDiscountUseCase() domain.AppliedDiscountUseCase {
	return &appliedDiscountUseCase{}
}

func validateCartItems(cartItems []domain.CartItem) error {
	if len(cartItems) == 0 {
		return domain.ErrEmptyCart
	}

	for _, item := range cartItems {
		if item.Quantity <= 0 {
			return fmt.Errorf("invalid quantity for item: %v", item)
		}
		if item.UnitPrice <= 0 {
			return fmt.Errorf("invalid unit price for item: %v", item)
		}
		if item.TotalPrice <= 0 {
			return fmt.Errorf("invalid total price for item: %v", item)
		}
	}
	return nil
}

func calculateTotalPrice(cartItems []domain.CartItem) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.TotalPrice
	}
	return totalPrice
}

func (uc *appliedDiscountUseCase) CalculateFixedAmountDiscount(ctx context.Context, cartItems []domain.CartItem, amount float64) (float64, error) {
	if err := validateCartItems(cartItems); err != nil {
		return 0, err
	}
	if amount < 0 {
		return 0, domain.ErrInvalidDiscountAmount
	}

	totalPrice := calculateTotalPrice(cartItems)
	return math.Max(0, totalPrice-amount), nil
}

func (uc *appliedDiscountUseCase) CalculatePercentageDiscount(ctx context.Context, cartItems []domain.CartItem, percentage float64) (float64, error) {
	if err := validateCartItems(cartItems); err != nil {
		return 0, err
	}
	if percentage < 0 || percentage > 100 {
		return 0, domain.ErrInvalidDiscountPercentage
	}

	totalPrice := calculateTotalPrice(cartItems)
	discount := totalPrice * (percentage / 100)
	return math.Max(0, totalPrice-discount), nil
}

func (uc *appliedDiscountUseCase) CalculateCategoryDiscount(ctx context.Context, cartItems []domain.CartItem, category string, percentage float64) (float64, error) {
	if err := validateCartItems(cartItems); err != nil {
		return 0, err
	}
	if category == "" {
		return 0, domain.ErrInvalidCategory
	}
	if percentage < 0 || percentage > 100 {
		return 0, domain.ErrInvalidDiscountPercentage
	}

	var categoryTotal float64
	var hasCategory bool

	for _, item := range cartItems {
		if item.Category == category {
			hasCategory = true
			discount := item.TotalPrice * (percentage / 100)
			categoryTotal += item.TotalPrice - discount
		} else {
			categoryTotal += item.TotalPrice
		}
	}

	if !hasCategory {
		return 0, domain.ErrCategoryNotFound
	}

	return math.Max(0, categoryTotal), nil
}

func (uc *appliedDiscountUseCase) CalculatePointsDiscount(ctx context.Context, cartItems []domain.CartItem, points int) (float64, error) {
	if err := validateCartItems(cartItems); err != nil {
		return 0, err
	}
	if points < 0 {
		return 0, domain.ErrInvalidPoints
	}

	totalPrice := calculateTotalPrice(cartItems)
	maxDiscount := totalPrice * 0.20 // 20% cap
	pointValue := float64(points)    // 1 point = 1 THB
	discount := math.Min(pointValue, maxDiscount)
	return math.Max(0, totalPrice-discount), nil
}

func (uc *appliedDiscountUseCase) CalculateSpecialDiscount(ctx context.Context, cartItems []domain.CartItem, threshold, discountAmount float64) (float64, error) {
	if err := validateCartItems(cartItems); err != nil {
		return 0, err
	}
	if threshold <= 0 {
		return 0, domain.ErrInvalidThreshold
	}
	if discountAmount <= 0 {
		return 0, domain.ErrInvalidDiscountAmount
	}

	totalPrice := calculateTotalPrice(cartItems)
	if totalPrice < threshold {
		return totalPrice, nil // Return original price if below threshold
	}

	discountTimes := math.Floor(totalPrice / threshold)
	totalDiscount := discountTimes * discountAmount
	return math.Max(0, totalPrice-totalDiscount), nil
}
