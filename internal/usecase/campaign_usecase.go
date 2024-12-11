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

type compaginUseCase struct {
	campaignRepo domain.CampaignRepository
}

func NewCampaignUseCase(cr domain.CampaignRepository) domain.CampaignUseCase {
	return &compaginUseCase{
		campaignRepo: cr,
	}
}

func (uc *compaginUseCase) Create(ctx context.Context, campaign *domain.Campaign) error {
	return uc.campaignRepo.Create(ctx, campaign)
}

func (uc *compaginUseCase) GetByID(ctx context.Context, id string) (*domain.Campaign, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, domain.ErrInvalidCampaignID
	}

	campaign, err := uc.campaignRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrCampaignNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf(strings.ToLower(constants.CampaignNotFoundError))
	}
	return campaign, nil
}

func (uc *compaginUseCase) GetAll(ctx context.Context) ([]domain.Campaign, error) {
	return uc.campaignRepo.FindAll(ctx)
}

func (uc *compaginUseCase) Update(ctx context.Context, campaign *domain.Campaign) error {
	return uc.campaignRepo.Update(ctx, campaign)
}

func (uc *compaginUseCase) Delete(ctx context.Context, id string) error {
	if !primitive.IsValidObjectID(id) {
		return domain.ErrInvalidCampaignID
	}

	return uc.campaignRepo.Delete(ctx, id)
}