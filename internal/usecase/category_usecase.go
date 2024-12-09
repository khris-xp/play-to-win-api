package usecase

import (
	"context"
	"errors"
	"fmt"
	"play-to-win-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUseCase(cr domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: cr,
	}
}

func (uc *categoryUseCase) Create(ctx context.Context, category *domain.Category) error {
	return uc.categoryRepo.Create(ctx, category)
}

func (uc *categoryUseCase) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidCategoryID
	}

	category, err := uc.categoryRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (uc *categoryUseCase) GetAll(ctx context.Context) ([]domain.Category, error) {
	return uc.categoryRepo.FindAll(ctx)
}

func (uc *categoryUseCase) Update(ctx context.Context, category *domain.Category) error {
	return uc.categoryRepo.Update(ctx, category)
}

func (uc *categoryUseCase) Delete(ctx context.Context, id string) error {
	if !primitive.IsValidObjectID(id) {
		return domain.ErrInvalidCategoryID
	}

	return uc.categoryRepo.Delete(ctx, id)
}
