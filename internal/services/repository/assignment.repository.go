package repository

import (
	"context"
	"errors"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type AssignmentRepository interface {
	// Define methods related to Assignment repository here
	GetAssignmentsByClassID(ctx context.Context, owner int, classID int) (*[]types.AssignmentResponse, error)
	CreateAssignment(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) error
	GetAssignmentsByAssignmentID(ctx context.Context, owner, classID, assignmentID int) (*types.AssignmentResponse, error)
	EditAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int, assignment *types.EditAssignmentRequest) error
	DeleteAssignmentByAssignmentID(ctx context.Context, owner, classID, assignmentID int) error
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

func (r *assignmentRepository) GetAssignmentsByClassID(ctx context.Context, owner int, classID int) (*[]types.AssignmentResponse, error) {
	var classes []model.Class
	if err := r.db.WithContext(ctx).Where("owner = ? AND id = ?", owner, classID).Find(&classes).Error; err != nil {
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
		assignmentResponses = append(assignmentResponses, types.AssignmentResponse{
			ID:          int(assignment.ID),
			ClassID:     classID,
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     dueDate,
			MaxAttempt:  assignment.MaxAttempt,
			Grade:       assignment.Grade,
			Settings: types.AssignmentSettings{
				GradePolicy: types.GradePolicy{
					Mode: "auto",
					Weight: types.WeightPolicy{
						TestCaseWeight:         0.5,
						NumberOfNodeUsedWeight: 0.5,
					},
				},
				TestCasePolicy: types.TestCasePolicy{
					VisibleToStudent: true,
				},
				FEBehavior: types.FEBehavior{
					LockAfterSubmit:      true,
					AllowResumitAfterDue: false,
				},
			},
			Condition: types.AssignmentCondition{
				System: map[string]int{
					"label": 2,
				},
				DataMovement: map[string]int{
					"mov": 2,
				},
				Arithmetic: map[string]int{
					"add": 1,
					"sub": 1,
				},
				ComparisonAndConditional: map[string]int{
					"cmp": 1,
					"jmp": 1,
				},
			},
		})
	}

	return &assignmentResponses, nil
}

func (r *assignmentRepository) CreateAssignment(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) error {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Class
	if err := r.db.WithContext(ctx).Where("id = ? AND owner = ?", classID, owner).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	newAssignment := model.Assignment{
		ClassID:     classID,
		Title:       assignment.Title,
		Description: assignment.Description,
		DueDate:     time.Now().Add(7 * 24 * time.Hour), // Example due date set to one week from now
		MaxAttempt:  assignment.MaxAttempt,
		Grade:       assignment.Grade,
		Setting:     nil,
		Condition:   nil,
	}

	if err := r.db.WithContext(ctx).Create(&newAssignment).Error; err != nil {
		return err
	}
	return nil
}

func (r *assignmentRepository) GetAssignmentsByAssignmentID(ctx context.Context, owner, classID, assignmentID int) (*types.AssignmentResponse, error) {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
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

	assignmentResponse := &types.AssignmentResponse{
		ID:          int(assignment.ID),
		ClassID:     classID,
		Title:       assignment.Title,
		Description: assignment.Description,
		DueDate:     dueDate,
		MaxAttempt:  assignment.MaxAttempt,
		Grade:       assignment.Grade,
		Settings: types.AssignmentSettings{
			GradePolicy: types.GradePolicy{
				Mode: "auto",
				Weight: types.WeightPolicy{
					TestCaseWeight:         0.5,
					NumberOfNodeUsedWeight: 0.5,
				},
			},
			TestCasePolicy: types.TestCasePolicy{
				VisibleToStudent: true,
			},
			FEBehavior: types.FEBehavior{
				LockAfterSubmit:      true,
				AllowResumitAfterDue: false,
			},
		},
		Condition: types.AssignmentCondition{
			System: map[string]int{
				"label": 2,
			},
			DataMovement: map[string]int{
				"mov": 2,
			},
			Arithmetic: map[string]int{
				"add": 1,
				"sub": 1,
			},
			ComparisonAndConditional: map[string]int{
				"cmp": 1,
				"jmp": 1,
			},
		},
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

	var classes model.Class
	if err := r.db.WithContext(ctx).Where("id = ? AND owner = ?", classID, owner).First(&classes).Error; err != nil {
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

	var classes model.Class
	if err := r.db.WithContext(ctx).Where("id = ? AND owner = ?", classID, owner).First(&classes).Error; err != nil {
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
