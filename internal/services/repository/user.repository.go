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

type UserRepository interface {
	GetAllUsers(ctx context.Context) (*[]types.UserResponse, error)
	CreateUser(ctx context.Context, email, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, email, password string) error {
	var user model.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err == nil {
		return fmt.Errorf("user with email %s already exists", email)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := model.User{
			Email:    email,
			Password: security.HashPassword(password),
		}
		if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	}
	return fmt.Errorf("failed to create user: %w", err)
}

func (r *userRepository) GetAllUsers(ctx context.Context) (*[]types.UserResponse, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	var userResp []types.UserResponse
	for _, user := range users {
		userResp = append(userResp, types.UserResponse{
			ID:       int(user.ID),
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password, // Note: Password should not be returned in a real application
			Name:     user.Name,
			Avatar:   user.Avatar,
			Tel:      user.Tel,
		})
	}
	return &userResp, nil
}
