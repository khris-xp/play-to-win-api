package usecase

import (
	"context"
	"play-to-win-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ctx context.Context, category *domain.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCategoryUseCase_Create(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	category := &domain.Category{ID: primitive.NewObjectID(), Name: "Test Category"}
	mockRepo.On("Create", mock.Anything, category).Return(nil)

	err := uc.Create(context.Background(), category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_GetByID_ValidID(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	categoryID := primitive.NewObjectID().Hex()
	objectID, _ := primitive.ObjectIDFromHex(categoryID)
	expectedCategory := &domain.Category{ID: objectID, Name: "Test Category"}
	mockRepo.On("FindByID", mock.Anything, categoryID).Return(expectedCategory, nil)

	category, err := uc.GetByID(context.Background(), categoryID)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_GetByID_InvalidID(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	invalidID := "invalid-id"
	category, err := uc.GetByID(context.Background(), invalidID)
	assert.Nil(t, category)
	assert.ErrorIs(t, err, domain.ErrInvalidCategoryID)
}

func TestCategoryUseCase_GetByID_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	categoryID := primitive.NewObjectID().Hex()
	mockRepo.On("FindByID", mock.Anything, categoryID).Return((*domain.Category)(nil), domain.ErrCategoryNotFound)

	category, err := uc.GetByID(context.Background(), categoryID)
	assert.Nil(t, category)
	assert.ErrorIs(t, err, domain.ErrCategoryNotFound)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_GetAll(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	expectedCategories := []domain.Category{
		{ID: primitive.NewObjectID(), Name: "Category 1"},
		{ID: primitive.NewObjectID(), Name: "Category 2"},
	}
	mockRepo.On("FindAll", mock.Anything).Return(expectedCategories, nil)

	categories, err := uc.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_Update(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	category := &domain.Category{ID: primitive.NewObjectID(), Name: "Updated Category"}
	mockRepo.On("Update", mock.Anything, category).Return(nil)

	err := uc.Update(context.Background(), category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_Delete_ValidID(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	categoryID := primitive.NewObjectID().Hex()
	mockRepo.On("Delete", mock.Anything, categoryID).Return(nil)

	err := uc.Delete(context.Background(), categoryID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUseCase_Delete_InvalidID(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	uc := NewCategoryUseCase(mockRepo)

	invalidID := "invalid-id"
	err := uc.Delete(context.Background(), invalidID)
	assert.ErrorIs(t, err, domain.ErrInvalidCategoryID)
}
