package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, userID int) (*types.UserResponse, error)
	EditProfile(ctx context.Context, userID int, user *types.EditProfileRequest) error
	ChangePassword(ctx context.Context, userID int, newPassword string) error
	DeleteProfile(ctx context.Context, userID int) error
	UploadAvatar(ctx context.Context, userID int, avatarURL string) error
	GetAvatar(ctx context.Context, userID int) (string, error)
	ChangeAvatar(ctx context.Context, userID int, avatarURL string) error
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
		ID:           int(user.ID),
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         user.Name,
		PicturePath:  user.PicturePath,
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
	if user.PicturePath != "" && user.PicturePath != "string" {
		updatedData["picture_path"] = user.PicturePath
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

	oldPassword := user.PasswordHash
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

func (r *profileRepository) DeleteProfile(ctx context.Context, userID int) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	if user.PicturePath != "" {
		fileName := filepath.Base(user.PicturePath)
		filePath := filepath.Join("uploads/users", fileName)

		if err := os.Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to delete avatar file: %w", err)
		}
	}

	if err := r.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *profileRepository) UploadAvatar(ctx context.Context, userID int, avatarURL string) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	if err := r.db.WithContext(ctx).Model(&user).Update("picture_path", avatarURL).Error; err != nil {
		return fmt.Errorf("failed to upload avatar: %w", err)
	}
	return nil
}

func (r *profileRepository) GetAvatar(ctx context.Context, userID int) (string, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("user with ID %d does not exist", userID)
	}

	return user.PicturePath, nil
}

func (r *profileRepository) ChangeAvatar(ctx context.Context, userID int, avatarURL string) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	oldPath := user.PicturePath
	if oldPath != "" {
		fileName := filepath.Base(oldPath)
		filePath := filepath.Join("uploads/users", fileName)

		if err := os.Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to delete old avatar file: %w", err)
		}
	}

	if err := r.db.WithContext(ctx).Model(&user).Update("picture_path", avatarURL).Error; err != nil {
		return fmt.Errorf("failed to change avatar: %w", err)
	}
	return nil
}
