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

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(pr domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{
		productRepo: pr,
	}
}

func (uc *productUseCase) Create(ctx context.Context, product *domain.Product) error {
	return uc.productRepo.Create(ctx, product)
}

func (uc *productUseCase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidProductID
	}

	product, err := uc.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.ProductNotFoundError))
	}
	return product, nil
}

func (uc *productUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return uc.productRepo.FindAll(ctx)
}

func (uc *productUseCase) Update(ctx context.Context, product *domain.Product) error {
	return uc.productRepo.Update(ctx, product)
}

func (uc *productUseCase) Delete(ctx context.Context, id string) error {
	if !primitive.IsValidObjectID(id) {
		return domain.ErrInvalidProductID
	}

	return uc.productRepo.Delete(ctx, id)
}