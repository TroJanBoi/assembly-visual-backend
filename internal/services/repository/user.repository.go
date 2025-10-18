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
	UpdateUser(ctx context.Context, userID int, users *types.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userID int) error
	GetMeClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpdateUser(ctx context.Context, userID int, users *types.UpdateUserRequest) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	updatedData := make(map[string]interface{})
	if users.Password != "" {
		updatedData["password"] = security.HashPassword(users.Password)
	}
	if users.Name != "" {
		updatedData["name"] = users.Name
	}
	if users.Tel != "" {
		updatedData["tel"] = users.Tel
	}

	if err := r.db.WithContext(ctx).Model(&user).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
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

func (r *userRepository) DeleteUser(ctx context.Context, userID int) error {
	var existingUser model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&existingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d does not exist", userID)
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	if err := r.db.WithContext(ctx).Delete(&existingUser).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
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
			Password: user.Password,
			Name:     user.Name,
			Tel:      user.Tel,
		})
	}
	return &userResp, nil
}

func (r *userRepository) GetMeClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error) {
	var member []model.Member
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&member).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve member records: %w", err)
	}

	var classIDs []int
	for _, m := range member {
		classIDs = append(classIDs, int(m.ClassID))
	}

	var classes []model.Class
	if err := r.db.WithContext(ctx).Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve classes: %w", err)
	}

	var classResp []types.ClassMeResponse
	for _, class := range classes {
		classResp = append(classResp, types.ClassMeResponse{
			ID:          int(class.ID),
			Topic:       class.Topic,
			Description: class.Description,
			FavScore:    class.FavScore,
			Owner:       int(class.Owner),
		})
	}
	return &classResp, nil
}
