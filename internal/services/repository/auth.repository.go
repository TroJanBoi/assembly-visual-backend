package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/security"
	"gorm.io/gorm"
)

type AuthRepository interface {
	// Define authentication-related methods here
	RegisterUser(ctx context.Context, newUser *types.SignUpRequest) error
	LoginUser(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) LoginUser(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error) {
	var users model.User
	err := r.db.WithContext(ctx).Where("email = ? AND password_hash = ?", user.Email, security.HashPassword(user.Password)).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.LoginResponse{}, fmt.Errorf("user with email %s does not exist", user.Email)
		}
		return &types.LoginResponse{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	tokenStr, err := security.GenerateToken(int(users.ID))
	if err != nil {
		return &types.LoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return &types.LoginResponse{Token: tokenStr}, nil
}

func (r *authRepository) RegisterUser(ctx context.Context, newUser *types.SignUpRequest) error {
	var users model.User
	err := r.db.WithContext(ctx).Where("email = ?", newUser.Email).First(&users).Error
	if err == nil {
		return fmt.Errorf("user with email %s already exists", newUser.Email)
	}

	if newUser.Name == "" {
		newUser.Name = ""
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user := model.User{
			Email:        newUser.Email,
			PasswordHash: security.HashPassword(newUser.Password),
			Name:         newUser.Name,
		}
		if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	}
	return fmt.Errorf("failed to check existing user: %w", err)
}
