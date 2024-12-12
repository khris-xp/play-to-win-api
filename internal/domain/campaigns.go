package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Category  string             `bson:"category" json:"category" validate:"required"`
	IsActive  bool               `bson:"is_active" json:"is_active" default:"true"`
	StartDate time.Time          `bson:"start_date" json:"start_date" default:"time.Now()"`
	EndDate   time.Time          `bson:"end_date" json:"end_date" default:"time.Now().AddDate(0, 0, 7)"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CampaignRepository interface {
	Create(ctx context.Context, campaign *Campaign) error
	FindByID(ctx context.Context, id string) (*Campaign, error)
	FindAll(ctx context.Context) ([]Campaign, error)
	Update(ctx context.Context, campaign *Campaign) error
	Delete(ctx context.Context, id string) error
}

type CampaignUseCase interface {
	Create(ctx context.Context, campaign *Campaign) error
	GetByID(ctx context.Context, id string) (*Campaign, error)
	GetAll(ctx context.Context) ([]Campaign, error)
	Update(ctx context.Context, campaign *Campaign) error
	Delete(ctx context.Context, id string) error
}
