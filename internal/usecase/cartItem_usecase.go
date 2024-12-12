package usecase

import (
	"context"
	"errors"
	"fmt"
	"play-to-win-api/internal/constants"
	"play-to-win-api/internal/domain"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartItemUseCase struct {
	cartItemRepo domain.CartItemRepository
}

func NewCartItemUseCase(cr domain.CartItemRepository) domain.CartItemUseCase {
	return &cartItemUseCase{
		cartItemRepo: cr,
	}
}

func (uc *cartItemUseCase) Create(ctx context.Context, cartItem *domain.CartItem) error {
	return uc.cartItemRepo.Create(ctx, cartItem)
}

func (uc *cartItemUseCase) GetByCartID(ctx context.Context, cartID string) ([]domain.CartItem, error) {
	if !primitive.IsValidObjectID(cartID) {
		return nil, domain.ErrInvalidCartID
	}

	cartItem, err := uc.cartItemRepo.FindByCartID(ctx, cartID)
	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.CartItemNotFound))
	}
	return cartItem, nil
}

func (uc *cartItemUseCase) GetByID(ctx context.Context, id string) (*domain.CartItem, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidCartItemID
	}

	cartItem, err := uc.cartItemRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCartItemNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.CartItemNotFound))
	}
	return cartItem, nil
}

func (uc *cartItemUseCase) GetAll(ctx context.Context) ([]domain.CartItem, error) {
	return uc.cartItemRepo.FindAll(ctx)
}

func (uc *cartItemUseCase) Update(ctx context.Context, cartItem *domain.CartItem) error {
	return uc.cartItemRepo.Update(ctx, cartItem)
}

func (uc *cartItemUseCase) Delete(ctx context.Context, id string) error {
	return uc.cartItemRepo.Delete(ctx, id)
}
