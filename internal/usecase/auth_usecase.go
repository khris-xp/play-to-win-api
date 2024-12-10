package usecase

import (
	"context"
	"play-to-win-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepo      domain.UserRepository
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewAuthUseCase(
	ur domain.UserRepository,
	accessSecret, refreshSecret string,
	accessTTL, refreshTTL time.Duration,
) domain.AuthUseCase {
	return &authUseCase{
		userRepo:      ur,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func (uc *authUseCase) Register(ctx context.Context, user *domain.User) (*domain.TokenPair, error) {
	existingUser, _ := uc.userRepo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	accessToken, err := uc.generateToken(user, uc.accessSecret, uc.accessTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.generateToken(user, uc.refreshSecret, uc.refreshTTL)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.UpdateRefreshToken(ctx, user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (uc *authUseCase) Login(ctx context.Context, email, password string) (*domain.TokenPair, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	return uc.generateTokenPair(user)
}

func (uc *authUseCase) RefreshToken(ctx context.Context, refreshToken string) (*domain.TokenPair, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(uc.refreshSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	user, err := uc.userRepo.FindByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	if user.RefreshToken != refreshToken {
		return nil, domain.ErrInvalidToken
	}

	return uc.generateTokenPair(user)
}

func (uc *authUseCase) generateTokenPair(user *domain.User) (*domain.TokenPair, error) {
	accessToken, err := uc.generateToken(user, uc.accessSecret, uc.accessTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.generateToken(user, uc.refreshSecret, uc.refreshTTL)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.UpdateRefreshToken(context.Background(), user.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *authUseCase) generateToken(user *domain.User, secret string, expiry time.Duration) (string, error) {
	claims := Claims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (uc *authUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uc.userRepo.FindByEmail(ctx, email)
}

func (uc *authUseCase) GetProfile(ctx context.Context, userID string) (*domain.UserProfile, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	return &domain.UserProfile{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}
