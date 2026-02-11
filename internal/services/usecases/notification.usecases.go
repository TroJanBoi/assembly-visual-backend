package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type NotificationUsecase interface {
	CreateNotificationUseCase(ctx context.Context, req *types.NotificationRequest) error
	GetNotificationsByUserIDUseCase(ctx context.Context, userID int) ([]*types.NotificationResponse, error)
	UpdateNotificationStatusUseCase(ctx context.Context, notificationID int, isRead bool) error
	DeleteNotificationUseCase(ctx context.Context, notificationID int) error
}

type notificationUsecase struct {
	notification repository.NotificationRepository
}

func NewNotificationUsecase(notificationRepo repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{
		notification: notificationRepo,
	}
}

func (u *notificationUsecase) CreateNotificationUseCase(ctx context.Context, req *types.NotificationRequest) error {
	if err := u.notification.CreateNotification(ctx, req); err != nil {
		return err
	}
	return nil
}

func (u *notificationUsecase) GetNotificationsByUserIDUseCase(ctx context.Context, userID int) ([]*types.NotificationResponse, error) {
	notifications, err := u.notification.GetNotificationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (u *notificationUsecase) UpdateNotificationStatusUseCase(ctx context.Context, notificationID int, isRead bool) error {
	if err := u.notification.UpdateNotificationStatus(ctx, notificationID, isRead); err != nil {
		return err
	}
	return nil
}

func (u *notificationUsecase) DeleteNotificationUseCase(ctx context.Context, notificationID int) error {
	if err := u.notification.DeleteNotification(ctx, notificationID); err != nil {
		return err
	}
	return nil
}
