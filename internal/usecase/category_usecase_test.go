package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"play-to-win-api/internal/domain"
	"play-to-win-api/internal/usecase"

	"github.com/fatih/color"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestStats struct {
	passed uint32
	failed uint32
	total  uint32
}

func (ts *TestStats) IncrementPassed() {
	atomic.AddUint32(&ts.passed, 1)
	atomic.AddUint32(&ts.total, 1)
}

func (ts *TestStats) IncrementFailed() {
	atomic.AddUint32(&ts.failed, 1)
	atomic.AddUint32(&ts.total, 1)
}

func (ts *TestStats) Print() {
	color.Cyan("\n=== Test Summary ===")
	color.Green("✓ Passed: %d", ts.passed)
	color.Red("❌ Failed: %d", ts.failed)
	color.Cyan("Total Tests: %d", ts.total)
	color.Cyan("Pass Rate: %.2f%%", float64(ts.passed)/float64(ts.total)*100)
}

type CategoryTestSuite struct {
	suite.Suite
	mockRepo *MockCategoryRepository
	useCase  domain.CategoryUseCase
	ctx      context.Context
	stats    *TestStats
}

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
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

func (s *CategoryTestSuite) SetupSuite() {
	s.stats = &TestStats{}
	fmt.Printf("\n%s\n", color.CyanString("=== Starting Category Use Case Test Suite ==="))
}

func (s *CategoryTestSuite) SetupTest() {
	s.mockRepo = new(MockCategoryRepository)
	s.useCase = usecase.NewCategoryUseCase(s.mockRepo)
	s.ctx = context.Background()
}

func (s *CategoryTestSuite) TearDownSuite() {
	s.stats.Print()
}

func (s *CategoryTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func logTestCase(name string) {
	fmt.Printf("\n%s\n", color.YellowString("⚡ Test Case: %s", name))
}

func (s *CategoryTestSuite) logTestResult(name string, err error) {
	if err != nil {
		color.Red("❌ %s: Failed - %v", name, err)
		s.stats.IncrementFailed()
	} else {
		color.Green("✓ %s: Passed", name)
		s.stats.IncrementPassed()
	}
}

func (s *CategoryTestSuite) TestCreate() {
	tests := []struct {
		name        string
		category    *domain.Category
		mockError   error
		expectError bool
	}{
		{
			name: "Success - Create Valid Category",
			category: &domain.Category{
				Name:        "Test Category",
				Description: "Test Description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockError:   nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			logTestCase(tt.name)
			s.mockRepo.On("Create", s.ctx, tt.category).Return(tt.mockError).Once()
			err := s.useCase.Create(s.ctx, tt.category)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.logTestResult(tt.name, err)
		})
	}
}

func (s *CategoryTestSuite) TestGetByID() {
	validID := primitive.NewObjectID().Hex()
	expectedCategory := &domain.Category{
		ID:          primitive.NewObjectID(),
		Name:        "Test Category",
		Description: "Test Description",
	}

	tests := []struct {
		name          string
		id            string
		mockCategory  *domain.Category
		mockError     error
		expectError   bool
		expectedValue *domain.Category
	}{
		{
			name:          "Success - Find Existing Category",
			id:            validID,
			mockCategory:  expectedCategory,
			mockError:     nil,
			expectError:   false,
			expectedValue: expectedCategory,
		},
		{
			name:          "Failure - Category Not Found",
			id:            validID,
			mockCategory:  nil,
			mockError:     domain.ErrCategoryNotFound,
			expectError:   true,
			expectedValue: nil,
		},
		{
			name:          "Failure - Invalid ID Format",
			id:            "invalid-id",
			mockCategory:  nil,
			mockError:     domain.ErrInvalidCategoryID,
			expectError:   true,
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			logTestCase(tt.name)

			if primitive.IsValidObjectID(tt.id) {
				s.mockRepo.On("FindByID", s.ctx, tt.id).Return(tt.mockCategory, tt.mockError).Once()
			}

			category, err := s.useCase.GetByID(s.ctx, tt.id)

			if tt.expectError {
				s.Error(err)
				s.Nil(category)
			} else {
				s.NoError(err)
				s.Equal(tt.expectedValue, category)
			}
			s.logTestResult(tt.name, err)
		})
	}
}

func (s *CategoryTestSuite) TestGetAll() {
	expectedCategories := []domain.Category{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Category 1",
			Description: "Description 1",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Category 2",
			Description: "Description 2",
		},
	}

	tests := []struct {
		name             string
		mockCategories   []domain.Category
		mockError        error
		expectError      bool
		expectedCatCount int
	}{
		{
			name:             "Success - Get All Categories",
			mockCategories:   expectedCategories,
			mockError:        nil,
			expectError:      false,
			expectedCatCount: 2,
		},
		{
			name:             "Success - Empty Category List",
			mockCategories:   []domain.Category{},
			mockError:        nil,
			expectError:      false,
			expectedCatCount: 0,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			logTestCase(tt.name)
			s.mockRepo.On("FindAll", s.ctx).Return(tt.mockCategories, tt.mockError).Once()

			categories, err := s.useCase.GetAll(s.ctx)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(categories, tt.expectedCatCount)
				if tt.expectedCatCount > 0 {
					s.Equal(tt.mockCategories, categories)
				}
			}
			s.logTestResult(tt.name, err)
		})
	}
}

func (s *CategoryTestSuite) TestUpdate() {
	category := &domain.Category{
		ID:          primitive.NewObjectID(),
		Name:        "Updated Category",
		Description: "Updated Description",
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name        string
		category    *domain.Category
		mockError   error
		expectError bool
	}{
		{
			name:        "Success - Update Category",
			category:    category,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "Failure - Update Error",
			category:    category,
			mockError:   errors.New("update failed"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			logTestCase(tt.name)
			s.mockRepo.On("Update", s.ctx, tt.category).Return(tt.mockError).Once()

			err := s.useCase.Update(s.ctx, tt.category)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.logTestResult(tt.name, err)
		})
	}
}

func (s *CategoryTestSuite) TestDelete() {
	validID := primitive.NewObjectID().Hex()

	tests := []struct {
		name        string
		id          string
		mockError   error
		expectError bool
	}{
		{
			name:        "Success - Delete Category",
			id:          validID,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "Failure - Category Not Found",
			id:          validID,
			mockError:   domain.ErrCategoryNotFound,
			expectError: true,
		},
		{
			name:        "Failure - Invalid ID Format",
			id:          "invalid-id",
			mockError:   domain.ErrInvalidCategoryID,
			expectError: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			logTestCase(tt.name)
			if tt.mockError != domain.ErrInvalidCategoryID {
				s.mockRepo.On("Delete", s.ctx, tt.id).Return(tt.mockError).Once()
			}

			err := s.useCase.Delete(s.ctx, tt.id)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.logTestResult(tt.name, err)
		})
	}
}

func TestCategoryUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}
