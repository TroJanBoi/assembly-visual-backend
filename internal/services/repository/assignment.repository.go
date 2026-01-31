package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AssignmentRepository interface {
	// Define methods related to Assignment repository here
	GetAssignmentsByClassID(ctx context.Context, classID int) (*[]types.AssignmentResponse, error)
	CreateAssignment(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) (int, error)
	GetAssignmentsByAssignmentID(ctx context.Context, classID, assignmentID int) (*types.AssignmentResponse, error)
	EditAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int, assignment *types.EditAssignmentRequest) error
	DeleteAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int) error
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

func (r *assignmentRepository) GetAssignmentsByClassID(ctx context.Context, classID int) (*[]types.AssignmentResponse, error) {
	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ?", classID).Find(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	if len(classes) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	dueDate := time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339) // Example due date set to one week from now := 2023-10-01T15:04:05Z07:00

	var assignments []model.Assignment
	if err := r.db.WithContext(ctx).Where("class_id = ?", classID).Find(&assignments).Error; err != nil {
		return nil, err
	}

	var assignmentResponses []types.AssignmentResponse
	for _, assignment := range assignments {

		var settings map[string]interface{}
		if len(assignment.Setting) > 0 {
			if err := json.Unmarshal(assignment.Setting, &settings); err != nil {
				return nil, err
			}
		} else {
			settings = map[string]interface{}{}
		}

		var condition map[string]interface{}
		if len(assignment.Condition) > 0 {
			if err := json.Unmarshal(assignment.Condition, &condition); err != nil {
				return nil, err
			}
		} else {
			condition = map[string]interface{}{}
		}

		assignmentResponses = append(assignmentResponses, types.AssignmentResponse{
			ID:          int(assignment.ID),
			ClassID:     classID,
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     dueDate,
			MaxAttempt:  assignment.MaxAttempt,
			Settings:    settings,
			Condition:   condition,
			Grade:       assignment.Grade,
		})
	}

	return &assignmentResponses, nil
}

func (r *assignmentRepository) CreateAssignment(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) (int, error) {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, err
	}

	settingBytes, _ := json.Marshal(assignment.Settings)
	conditionBytes, _ := json.Marshal(assignment.Condition)

	newAssignment := model.Assignment{
		ClassID:     classID,
		Title:       assignment.Title,
		Description: assignment.Description,
		DueDate:     time.Now().Add(7 * 24 * time.Hour), // Example due date set to one week from now
		MaxAttempt:  assignment.MaxAttempt,
		Setting:     datatypes.JSON(settingBytes),
		Condition:   datatypes.JSON(conditionBytes),
		Grade:       assignment.Grade,
	}

	if err := r.db.WithContext(ctx).Create(&newAssignment).Error; err != nil {
		return 0, err
	}
	return int(newAssignment.ID), nil
}

func (r *assignmentRepository) GetAssignmentsByAssignmentID(ctx context.Context, classID, assignmentID int) (*types.AssignmentResponse, error) {
	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	dueDate := assignment.DueDate.Format(time.RFC3339)

	// var settings map[string]interface{}
	// if err := json.Unmarshal(assignment.Setting, &settings); err != nil {
	// 	return nil, err
	// }
	// var condition map[string]interface{}
	// if err := json.Unmarshal(assignment.Condition, &condition); err != nil {
	// 	return nil, err
	// }

	var settings map[string]interface{}
	if len(assignment.Setting) > 0 {
		if err := json.Unmarshal(assignment.Setting, &settings); err != nil {
			return nil, err
		}
	} else {
		settings = map[string]interface{}{}
	}

	var condition map[string]interface{}
	if len(assignment.Condition) > 0 {
		if err := json.Unmarshal(assignment.Condition, &condition); err != nil {
			return nil, err
		}
	} else {
		condition = map[string]interface{}{}
	}

	assignmentResponse := &types.AssignmentResponse{
		ID:          int(assignment.ID),
		ClassID:     classID,
		Title:       assignment.Title,
		Description: assignment.Description,
		DueDate:     dueDate,
		MaxAttempt:  assignment.MaxAttempt,
		Settings:    settings,
		Condition:   condition,
		Grade:       assignment.Grade,
	}

	return assignmentResponse, nil
}

func (r *assignmentRepository) EditAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int, assignment *types.EditAssignmentRequest) error {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var existingAssignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&existingAssignment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if assignment.Title != "" {
		existingAssignment.Title = assignment.Title
	}
	if assignment.Description != "" {
		existingAssignment.Description = assignment.Description
	}
	if assignment.MaxAttempt != 0 {
		existingAssignment.MaxAttempt = assignment.MaxAttempt
	}
	if assignment.Grade != 0 {
		existingAssignment.Grade = assignment.Grade
	}

	settingBytes, _ := json.Marshal(assignment.Setting)
	conditionBytes, _ := json.Marshal(assignment.Condition)

	existingAssignment.Setting = datatypes.JSON(settingBytes)
	existingAssignment.Condition = datatypes.JSON(conditionBytes)

	if err := r.db.WithContext(ctx).Save(&existingAssignment).Error; err != nil {
		return err
	}
	return nil

}

func (r *assignmentRepository) DeleteAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int) error {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).Delete(&model.Assignment{}).Error; err != nil {
		return err
	}
	return nil
}
