package repository

import (
	"context"
	"errors"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type ClassRepository interface {
	GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error)
	GetClassByID(ctx context.Context, classID int) (*types.ClassResponse, error)
	CreateClass(ctx context.Context, owner uint, class *types.CreateClassRequest) error
	UpdateClass(ctx context.Context, owner uint, classID int, class *types.UpdateClassRequest) error
	DeleteClass(ctx context.Context, owner uint, classID int) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error) {
	var classes []model.Class
	if err := r.db.WithContext(ctx).Find(&classes).Error; err != nil {
		return nil, err
	}

	var classResponses []types.ClassResponse
	for _, class := range classes {
		classResponses = append(classResponses, types.ClassResponse{
			ID:               int(class.ID),
			Topic:            class.Topic,
			Description:      class.Description,
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			FavScore:         class.FavScore,
			Owner:            uint(class.Owner),
			Status:           class.Status,
		})
	}

	return &classResponses, nil
}

func (r *classRepository) CreateClass(ctx context.Context, owner uint, class *types.CreateClassRequest) error {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	newClass := model.Class{
		Topic:            class.Topic,
		Description:      class.Description,
		GoogleCourseID:   class.GoogleCourseID,
		GoogleCourseLink: class.GoogleCourseLink,
		Owner:            owner,
		Status:           class.Status,
	}

	if err := r.db.WithContext(ctx).Create(&newClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) UpdateClass(ctx context.Context, owner uint, classID int, class *types.UpdateClassRequest) error {
	var existingClass model.Class
	err := r.db.WithContext(ctx).Where("id = ? AND owner = ?", classID, owner).First(&existingClass).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if class.Topic != "" {
		existingClass.Topic = class.Topic
	}
	if class.Description != "" {
		existingClass.Description = class.Description
	}
	if class.GoogleCourseID != "" {
		existingClass.GoogleCourseID = class.GoogleCourseID
	}
	if class.GoogleCourseLink != "" {
		existingClass.GoogleCourseLink = class.GoogleCourseLink
	}
	if class.Status != 0 {
		existingClass.Status = class.Status
	}

	if err := r.db.WithContext(ctx).Save(&existingClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) DeleteClass(ctx context.Context, owner uint, classID int) error {
	var existingClass model.Class
	err := r.db.WithContext(ctx).Where("id = ? AND owner = ?", classID, owner).First(&existingClass).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&existingClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) GetClassByID(ctx context.Context, classID int) (*types.ClassResponse, error) {
	var class model.Class
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	classResponse := types.ClassResponse{
		ID:               int(class.ID),
		Topic:            class.Topic,
		Description:      class.Description,
		GoogleCourseID:   class.GoogleCourseID,
		GoogleCourseLink: class.GoogleCourseLink,
		GoogleSyncedAt:   class.GoogleSyncedAt,
		FavScore:         class.FavScore,
		Owner:            uint(class.Owner),
		Status:           class.Status,
	}

	return &classResponse, nil
}
