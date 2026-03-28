package repository

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, userID, classID int) error
	DeleteBookmark(ctx context.Context, userID, classID int) error
	GetBookmarksByUserID(ctx context.Context, userID int) (*[]types.ClassResponse, error)
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) BookmarkRepository {
	return &bookmarkRepository{db: db}
}

func (r *bookmarkRepository) CreateBookmark(ctx context.Context, userID, classID int) error {
	bookmark := model.BookMark{
		UserID:  userID,
		ClassID: classID,
	}

	if err := r.db.WithContext(ctx).Create(&bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookmarkRepository) DeleteBookmark(ctx context.Context, userID, classID int) error {
	if err := r.db.WithContext(ctx).Where("user_id = ? AND class_id = ?", userID, classID).Delete(&model.BookMark{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookmarkRepository) GetBookmarksByUserID(ctx context.Context, userID int) (*[]types.ClassResponse, error) {
	var bookmarks []model.BookMark
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	var classIDs []int
	for _, bookmark := range bookmarks {
		classIDs = append(classIDs, bookmark.ClassID)
	}

	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		return nil, err
	}

	classResponses := make([]types.ClassResponse, 0, len(classes))

	for _, class := range classes {
		var owner model.User
		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&owner).Error; err != nil {
			return nil, err
		}

		var memberCount int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&memberCount).Error; err != nil {
			return nil, err
		}
		classResponses = append(classResponses, types.ClassResponse{
			ID:               int(class.ID),
			Topic:            class.Topic,
			Description:      class.Description,
			Code:             class.Code,
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			OwnerID:          int(class.OwnerId),
			OwnerName:        owner.Name,
			MemberAmount:     memberCount,
			Status:           class.Status,
			BannerID:         class.BannerID,
		})
	}

	return &classResponses, nil
}
