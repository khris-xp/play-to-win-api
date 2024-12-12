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

type cartUseCase struct {
	cartRepo domain.CartRepository
}

func NewCartUseCase(cr domain.CartRepository) domain.CartUseCase {
	return &cartUseCase{
		cartRepo: cr,
	}
}

func (uc *cartUseCase) Create(ctx context.Context, cart *domain.Cart) error {
	return uc.cartRepo.Create(cart)
}

func (uc *cartUseCase) GetByUserID(ctx context.Context, userID string) ([]domain.Cart, error) {
	if !primitive.IsValidObjectID(userID) {
		return nil, domain.ErrInvalidUserID
	}

	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrCartNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.CartNotFound))
	}
	return cart, nil
}

func (uc *cartUseCase) GetByID(ctx context.Context, id string) (*domain.Cart, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidCartID
	}

	cart, err := uc.cartRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCartNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.CartNotFound))
	}
	return cart, nil
}

func (uc *cartUseCase) GetAll(ctx context.Context) ([]domain.Cart, error) {
	return uc.cartRepo.FindAll(ctx)
}

func (uc *cartUseCase) Update(ctx context.Context, cart *domain.Cart) error {
	return uc.cartRepo.Update(ctx, cart)
}

func (uc *cartUseCase) Delete(ctx context.Context, id string) error {
	if !primitive.IsValidObjectID(id) {
		return domain.ErrInvalidCartID
	}

	return uc.cartRepo.Delete(ctx, id)
}
