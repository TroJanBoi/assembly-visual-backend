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
	CreateUser(ctx context.Context, users *types.CreateUserRequest) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, users *types.CreateUserRequest) error {
	var user model.User

	err := r.db.WithContext(ctx).Where("email = ?", users.Email).First(&user).Error
	if err == nil {
		return fmt.Errorf("user with email %s already exists", users.Email)
	}

	if users.Name == "" {
		users.Name = ""
	}

	if users.Tel == "" {
		users.Tel = ""
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := model.User{
			Email:    users.Email,
			Password: security.HashPassword(users.Password),
			Name:     users.Name,
			Tel:      users.Tel,
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
			Password: user.Password, // Note: Password should not be returned in a real application
			Name:     user.Name,
			Tel:      user.Tel,
		})
	}
	return &userResp, nil
}
