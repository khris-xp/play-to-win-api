package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Email        string             `bson:"email" json:"email" validate:"required,email"`
	Password     string             `bson:"password" json:"password,omitempty" validate:"required,min=6"`
	Role         string             `bson:"role" json:"role"`
	RefreshToken string             `bson:"refresh_token,omitempty" json:"-"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserProfile struct {
	ID        primitive.ObjectID `json:"id"`
	Email     string             `json:"email"`
	Name      string             `json:"name"`
	Role      string             `json:"role"`
	CreatedAt time.Time          `json:"created_at"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	UpdateRefreshToken(ctx context.Context, userID primitive.ObjectID, token string) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*User, error)
}

type AuthUseCase interface {
	Register(ctx context.Context, user *User) (*TokenPair, error)
	Login(ctx context.Context, email, password string) (*TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error)
	GetProfile(ctx context.Context, id string) (*UserProfile, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Claims struct {
	ID string `json:"id"`

	Email string `json:"email"`
}
