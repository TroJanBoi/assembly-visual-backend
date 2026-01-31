package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PlaygroundRepository interface {
	Create(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error)
	GetByPlaygroundID(ctx context.Context, userID int, id int) (*types.PlaygroundResponse, error)
	GetPlaygroundByMe(ctx context.Context, userID int, req *types.PlaygroundMeRequest) (*types.PlaygroundResponse, error)
	UpdatePlaygroundByMe(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error)
	DeletePlaygroundByMe(ctx context.Context, userID int, req *types.PlaygroundMeRequest) error
}

type playgroundRepository struct {
	db *gorm.DB
}

func NewPlaygroundRepository(db *gorm.DB) PlaygroundRepository {
	return &playgroundRepository{db: db}
}

func (r *playgroundRepository) Create(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ?", playground.AssignmentID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("assignment not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find assignment: %w", err)
	}

	var existingPlayground model.Playground
	if err := r.db.WithContext(ctx).Where("assignment_id = ? AND user_id = ?", playground.AssignmentID, userID).First(&existingPlayground).Error; err == nil {
		return nil, fmt.Errorf("playground for this assignment and user already exists: %w", err)
	}

	item, _ := json.Marshal(playground.Item)

	newPlayground := model.Playground{
		AssignmentID: int(playground.AssignmentID),
		UserID:       int(userID),
		Item:         datatypes.JSON(item),
		Status:       playground.Status,
	}

	if err := r.db.WithContext(ctx).Create(&newPlayground).Error; err != nil {
		return nil, fmt.Errorf("failed to create playground: %w", err)
	}

	// var parseItems types.PlaygroundData
	// if err := json.Unmarshal(newPlayground.Item, &parseItems); err != nil {
	// 	return nil, fmt.Errorf("failed to unmarshal playground item: %w", err)
	// }
	var parseItems map[string]interface{}
	if len(newPlayground.Item) > 0 {
		if err := json.Unmarshal(newPlayground.Item, &parseItems); err != nil {
			return nil, fmt.Errorf("failed to unmarshal playground item: %w", err)
		}
	} else {
		parseItems = map[string]interface{}{}
	}

	playgroundResponse := &types.PlaygroundResponse{
		ID:           int(newPlayground.ID),
		AssignmentID: newPlayground.AssignmentID,
		UserID:       newPlayground.UserID,
		Item:         parseItems,
		Status:       playground.Status,
	}

	return playgroundResponse, nil
}

func (r *playgroundRepository) GetByPlaygroundID(ctx context.Context, userID int, id int) (*types.PlaygroundResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	var playground model.Playground
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&playground).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("playground not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find playground: %w", err)
	}

	var parseItems map[string]interface{}
	if len(playground.Item) > 0 {
		if err := json.Unmarshal(playground.Item, &parseItems); err != nil {
			return nil, fmt.Errorf("failed to unmarshal playground item: %w", err)
		}
	} else {
		parseItems = map[string]interface{}{}
	}

	playgroundResponse := &types.PlaygroundResponse{
		ID:           int(playground.ID),
		AssignmentID: playground.AssignmentID,
		UserID:       playground.UserID,
		Item:         parseItems,
		Status:       playground.Status,
	}

	return playgroundResponse, nil
}

func (r *playgroundRepository) GetPlaygroundByMe(ctx context.Context, userID int, req *types.PlaygroundMeRequest) (*types.PlaygroundResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	var playground model.Playground
	if err := r.db.WithContext(ctx).Where("assignment_id = ? AND user_id = ?", req.AssignmentID, userID).First(&playground).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("playground not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find playground: %w", err)
	}

	var parseItems map[string]interface{}
	if len(playground.Item) > 0 {
		if err := json.Unmarshal(playground.Item, &parseItems); err != nil {
			return nil, fmt.Errorf("failed to unmarshal playground item: %w", err)
		}
	} else {
		parseItems = map[string]interface{}{}
	}

	playgroundResponse := &types.PlaygroundResponse{
		ID:           int(playground.ID),
		AssignmentID: playground.AssignmentID,
		UserID:       playground.UserID,
		Item:         parseItems,
		Status:       playground.Status,
	}
	return playgroundResponse, nil
}

func (r *playgroundRepository) UpdatePlaygroundByMe(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	var existingPlayground model.Playground
	if err := r.db.WithContext(ctx).Where("assignment_id = ? AND user_id = ?", playground.AssignmentID, userID).First(&existingPlayground).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("playground not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find playground: %w", err)
	}

	item, _ := json.Marshal(playground.Item)
	existingPlayground.Item = datatypes.JSON(item)
	existingPlayground.Status = playground.Status

	if err := r.db.WithContext(ctx).Save(&existingPlayground).Error; err != nil {
		return nil, fmt.Errorf("failed to update playground: %w", err)
	}

	var parseItems map[string]interface{}
	if len(existingPlayground.Item) > 0 {
		if err := json.Unmarshal(existingPlayground.Item, &parseItems); err != nil {
			return nil, fmt.Errorf("failed to unmarshal playground item: %w", err)
		}
	} else {
		parseItems = map[string]interface{}{}
	}

	playgroundResponse := &types.PlaygroundResponse{
		ID:           int(existingPlayground.ID),
		AssignmentID: existingPlayground.AssignmentID,
		UserID:       existingPlayground.UserID,
		Item:         parseItems,
		Status:       existingPlayground.Status,
	}
	return playgroundResponse, nil
}

func (r *playgroundRepository) DeletePlaygroundByMe(ctx context.Context, userID int, req *types.PlaygroundMeRequest) error {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	var existingPlayground model.Playground
	if err := r.db.WithContext(ctx).Where("assignment_id = ? AND user_id = ?", req.AssignmentID, userID).First(&existingPlayground).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("playground not found: %w", err)
		}
		return fmt.Errorf("failed to find playground: %w", err)
	}

	if err := r.db.WithContext(ctx).Unscoped().Where("id = ?", existingPlayground.ID).Delete(&existingPlayground).Error; err != nil {
		return fmt.Errorf("failed to hard delete playground: %w", err)
	}

	return nil
}
