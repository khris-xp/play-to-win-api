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

type discountRuleUseCase struct {
	discountRuleRepo domain.DiscountRuleRepository
}

func NewDiscountRuleUseCase(dr domain.DiscountRuleRepository) domain.DiscountRuleUseCase {
	return &discountRuleUseCase{
		discountRuleRepo: dr,
	}
}

func (uc *discountRuleUseCase) Create(ctx context.Context, discountRule *domain.DiscountRule) error {
	return uc.discountRuleRepo.Create(ctx, discountRule)
}

func (uc *discountRuleUseCase) GetByID(ctx context.Context, id string) (*domain.DiscountRule, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidDiscountRuleID
	}

	discountRule, err := uc.discountRuleRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrDiscountRuleNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.DiscountRuleNotFound))
	}
	return discountRule, nil
}

func (uc *discountRuleUseCase) GetAll(ctx context.Context) ([]domain.DiscountRule, error) {
	return uc.discountRuleRepo.FindAll(ctx)
}

func (uc *discountRuleUseCase) Update(ctx context.Context, discountRule *domain.DiscountRule) error {
	return uc.discountRuleRepo.Update(ctx, discountRule)
}

func (uc *discountRuleUseCase) Delete(ctx context.Context, id string) error {
	if !primitive.IsValidObjectID(id) {
		return domain.ErrInvalidDiscountRuleID
	}

	err := uc.discountRuleRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrDiscountRuleNotFound) {
			return err
		}
		return fmt.Errorf(strings.ToLower(constants.DiscountRuleNotFound))
	}
	return nil
}
