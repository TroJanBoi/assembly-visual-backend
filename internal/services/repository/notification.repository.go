package repository

import (
	"context"
	"encoding/json"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, req *types.NotificationRequest) error
	GetNotificationsByUserID(ctx context.Context, userID int) ([]*types.NotificationResponse, error)
	UpdateNotificationStatus(ctx context.Context, notificationID int, isRead bool) error
	DeleteNotification(ctx context.Context, notificationID int) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) CreateNotification(ctx context.Context, req *types.NotificationRequest) error {
	usr := model.User{}
	if err := r.db.WithContext(ctx).Where("id = ?", req.UserID).First(&usr).Error; err != nil {
		return err
	}

	dataJSON, err := json.Marshal(req.Data)
	if err != nil {
		return err
	}

	notification := model.Notification{
		UserID:  req.UserID,
		Type:    req.Type,
		Title:   req.Title,
		Message: req.Message,
		Data:    datatypes.JSON(dataJSON),
		IsRead:  false,
	}
	if err := r.db.WithContext(ctx).Create(&notification).Error; err != nil {
		return err
	}

	return nil
}

func (r *notificationRepository) GetNotificationsByUserID(ctx context.Context, userID int) ([]*types.NotificationResponse, error) {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&usr).Error; err != nil {
		return nil, err
	}

	var notifications []model.Notification
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error; err != nil {
		return nil, err
	}

	var result []*types.NotificationResponse
	for _, n := range notifications {
		var data map[string]interface{}
		if err := json.Unmarshal(n.Data, &data); err != nil {
			return nil, err
		}

		notificationResp := &types.NotificationResponse{
			ID:      int(n.ID),
			UserID:  n.UserID,
			Type:    n.Type,
			Title:   n.Title,
			Message: n.Message,
			Data:    data,
			IsRead:  n.IsRead,
		}
		result = append(result, notificationResp)
	}

	return result, nil
}

func (r *notificationRepository) UpdateNotificationStatus(ctx context.Context, notificationID int, isRead bool) error {
	var notification model.Notification
	if err := r.db.WithContext(ctx).Where("id = ?", notificationID).First(&notification).Error; err != nil {
		return err
	}

	notification.IsRead = isRead
	if err := r.db.WithContext(ctx).Save(&notification).Error; err != nil {
		return err
	}

	return nil
}

func (r *notificationRepository) DeleteNotification(ctx context.Context, notificationID int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", notificationID).Delete(&model.Notification{}).Error; err != nil {
		return err
	}

	return nil
}
