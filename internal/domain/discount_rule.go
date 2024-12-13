package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiscountRule struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	CampaignID                  primitive.ObjectID `bson:"campaign_id,omitempty" json:"campaign_id" validate:"required"`
	DiscuntType                 string             `bson:"discount_type,omitempty" json:"discount_type" validate:"required"`
	Amount                      float64            `bson:"amount,omitempty" json:"amount" validate:"required"`
	Percentage                  float64            `bson:"percentage,omitempty" json:"percentage" validate:"required"`
	ItemCategory                string             `bson:"item_category,omitempty" json:"item_category" validate:"required"`
	PointsRatio                 float64            `bson:"points_ratio,omitempty" json:"points_ratio" validate:"required"`
	MaxDiscountPercentage       float64            `bson:"max_discount_percentage,omitempty" json:"max_discount_percentage" validate:"required"`
	ThresholdAmount             float64            `bson:"threshold_amount,omitempty" json:"threshold_amount" validate:"required"`
	DiscountPercentageThreshold float64            `bson:"discount_percentage_threshold,omitempty" json:"discount_percentage_threshold" validate:"required"`
	CreatedAt                   time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt                   time.Time          `bson:"updated_at,omitempty" json:"updated_at"`

	CampaignName string `bson:"campaign_name,omitempty" json:"campaign_name"`
}

type DiscountRuleRepository interface {
	Create(ctx context.Context, discountRule *DiscountRule) error
	FindByID(ctx context.Context, id string) (*DiscountRule, error)
	FindAll(ctx context.Context) ([]DiscountRule, error)
	Update(ctx context.Context, discountRule *DiscountRule) error
	Delete(ctx context.Context, id string) error
}

type DiscountRuleUseCase interface {
	Create(ctx context.Context, discountRule *DiscountRule) error
	GetByID(ctx context.Context, id string) (*DiscountRule, error)
	GetAll(ctx context.Context) ([]DiscountRule, error)
	Update(ctx context.Context, discountRule *DiscountRule) error
	Delete(ctx context.Context, id string) error
}
