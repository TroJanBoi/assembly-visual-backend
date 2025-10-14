package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, userID int) (*types.UserResponse, error)
	EditProfile(ctx context.Context, userID int, user *types.EditProfileRequest) error
	ChangePassword(ctx context.Context, userID int, newPassword string) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetProfile(ctx context.Context, userID int) (*types.UserResponse, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d does not exist", userID)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	userResp := &types.UserResponse{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Tel:      user.Tel,
	}
	return userResp, nil
}

func (r *profileRepository) EditProfile(ctx context.Context, userID int, user *types.EditProfileRequest) error {
	var existingUser model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&existingUser).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	updatedData := make(map[string]interface{})

	if user.Username != "" && user.Username != "string" {
		updatedData["username"] = user.Username
	}
	if user.Name != "" && user.Name != "string" {
		updatedData["name"] = user.Name
	}
	if user.Tel != "" && user.Tel != "string" {
		updatedData["tel"] = user.Tel
	}
	if user.Picture_path != "" && user.Picture_path != "string" {
		updatedData["picture_path"] = user.Picture_path
	}

	if err := r.db.WithContext(ctx).Model(&existingUser).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *profileRepository) ChangePassword(ctx context.Context, userID int, newPassword string) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	oldPassword := user.Password
	if oldPassword == newPassword {
		return fmt.Errorf("new password cannot be the same as the old password")
	}

	updatedData := map[string]interface{}{
		"password": newPassword,
	}

	if err := r.db.WithContext(ctx).Model(&user).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}
	return nil
}
