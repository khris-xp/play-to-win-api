package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description" validate:"required"`
	Content     string             `bson:"content" json:"content" validate:"required"`
	Price       float64            `bson:"price" json:"price" validate:"required"`
	Image       string             `bson:"image" json:"image" validate:"required"`
	Sold        int                `bson:"sold" json:"sold"`
	Stock       int                `bson:"stock" json:"stock"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

type ProductUseCase interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}
