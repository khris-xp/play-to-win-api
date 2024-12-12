package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	User        User               `bson:"user" json:"user"`
	TotalAmount float64            `bson:"total_amount" json:"total_amount" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type CartRepository interface {
	Create(cart *Cart) error
	FindByUserID(ctx context.Context, userID string) ([]Cart, error)
	FindByID(ctx context.Context, id string) (*Cart, error)
	FindAll(ctx context.Context) ([]Cart, error)
	Update(ctx context.Context, cart *Cart) error
	Delete(ctx context.Context, id string) error
}

type CartUseCase interface {
	Create(ctx context.Context, cart *Cart) error
	GetByUserID(ctx context.Context, userID string) ([]Cart, error)
	GetByID(ctx context.Context, id string) (*Cart, error)
	GetAll(ctx context.Context) ([]Cart, error)
	Update(ctx context.Context, cart *Cart) error
	Delete(ctx context.Context, id string) error
}
