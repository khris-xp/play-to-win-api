package repository

import (
	"context"
	"play-to-win-api/internal/domain"
)

type Repositories struct {
	Category domain.CategoryRepository
}

type Transaction interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewRepositories(categoryRepo domain.CategoryRepository) *Repositories {
	return &Repositories{
		Category: categoryRepo,
	}
}
