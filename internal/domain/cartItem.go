package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CartId     primitive.ObjectID `bson:"cart_id,omitempty" json:"cart_id" validate:"required"`
	ProductId  primitive.ObjectID `bson:"product_id,omitempty" json:"product_id" validate:"required"`
	Quantity   int                `bson:"quantity,omitempty" json:"quantity" validate:"required"`
	Category   string             `bson:"category,omitempty" json:"category" validate:"required"`
	UnitPrice  float64            `bson:"unit_price,omitempty" json:"unit_price" validate:"required"`
	TotalPrice float64            `bson:"total_price,omitempty" json:"total_price" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at"`

	ProductName        string  `bson:"product_name,omitempty" json:"product_name"`
	ProductDescription string  `bson:"product_description,omitempty" json:"product_description"`
	ProductImage       string  `bson:"product_image,omitempty" json:"product_image"`
	ProductPrice       float64 `bson:"product_price,omitempty" json:"product_price"`
}

type CartItemRepository interface {
	Create(ctx context.Context, cartItem *CartItem) error
	FindByCartID(ctx context.Context, cartID string) ([]CartItem, error)
	FindByID(ctx context.Context, id string) (*CartItem, error)
	FindAll(ctx context.Context) ([]CartItem, error)
	Update(ctx context.Context, cartItem *CartItem) error
	Delete(ctx context.Context, id string) error
}

type CartItemUseCase interface {
	Create(ctx context.Context, cartItem *CartItem) error
	GetByCartID(ctx context.Context, cartID string) ([]CartItem, error)
	GetByID(ctx context.Context, id string) (*CartItem, error)
	GetAll(ctx context.Context) ([]CartItem, error)
	Update(ctx context.Context, cartItem *CartItem) error
	Delete(ctx context.Context, id string) error
}
