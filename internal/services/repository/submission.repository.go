package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	CreateSubmission(ctx context.Context, userId int, request types.CreateSubmissionRequest) error
	UpdateSubmission(ctx context.Context, userID, submissionID int, request types.UpdateSubmissionRequest) error
	GetAllSubmissionByAssignmentID(ctx context.Context, ownerID, assignmentID int) (*[]types.SubmissionResponse, error)
	GetSubmissionByID(ctx context.Context, userID, submissionID int) (*types.SubmissionResponse, error)
	GetAllSubmissionByAssignmentIDandUserID(ctx context.Context, assignmentID, userID int) (*[]types.SubmissionResponse, error)
	UpdateGrade(ctx context.Context, userID, submissionID int, request types.UpdateGradeRequest) error
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (r *submissionRepository) CreateSubmission(ctx context.Context, userId int, request types.CreateSubmissionRequest) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}

	var assignment model.Assignment
	err = r.db.WithContext(ctx).Where("id = ?", request.AssignmentID).First(&assignment).Error
	if err != nil {
		return err
	}

	var playground model.Playground
	err = r.db.WithContext(ctx).Where("id = ?", request.PlaygroundID).First(&playground).Error
	if err != nil {
		return err
	}

	snapshot, _ := json.Marshal(request.ItemSnapshot)
	clientResult, _ := json.Marshal(request.ClientResult)
	serverResult, _ := json.Marshal(request.ServerResult)

	if request.Score > float64(assignment.Grade) {
		request.Score = float64(assignment.Grade)
	}

	newSubmission := model.Submission{
		AssignmentID:  int(request.AssignmentID),
		UserID:        int(userId),
		PlaygroundID:  int(request.PlaygroundID),
		AttemptNumber: int(request.AttemptNumber),
		ItemSnapshot:  datatypes.JSON(snapshot),
		ClientResult:  datatypes.JSON(clientResult),
		ServerResult:  datatypes.JSON(serverResult),
		Score:         request.Score,
		Status:        request.Status,
		IsVerified:    request.IsVerified,
		DurationMS:    request.DurationMS,
	}
	if err := r.db.WithContext(ctx).Create(&newSubmission).Error; err != nil {
		return err
	}
	return nil
}

func (r *submissionRepository) UpdateSubmission(ctx context.Context, userID, submissionID int, request types.UpdateSubmissionRequest) error {
	var submission model.Submission
	err := r.db.WithContext(ctx).Where("id = ?", submissionID).First(&submission).Error
	if err != nil {
		return err
	}

	var assignment model.Assignment
	err = r.db.WithContext(ctx).Where("id = ?", submission.AssignmentID).First(&assignment).Error
	if err != nil {
		return err
	}

	// อนุญาตให้แก้ไขได้ทุกฟิลด์ เจ้าของ submission และ owner ของ assignment และ member ที่เป็น TA เท่านั้น
	// Select * FROM classroom c JOIN assignment a ON a.class_id = c.id JOIN member m ON m.class_id = c.id WHERE a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?))
	err = r.db.WithContext(ctx).Table("classroom c").Select("c.id, c.owner_id").Joins("JOIN assignment a ON a.class_id = c.id").Joins("JOIN member m ON m.class_id = c.id").Joins("JOIN submission s ON s.assignment_id = a.id").Where("a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?) OR s.user_id = ?)", submission.AssignmentID, userID, userID, "ta", userID).First(&model.Classroom{}).Error
	if err != nil {
		return err
	}

	snapshot, _ := json.Marshal(request.ItemSnapshot)
	clientResult, _ := json.Marshal(request.ClientResult)
	serverResult, _ := json.Marshal(request.ServerResult)

	if request.Score > float64(assignment.Grade) {
		request.Score = float64(assignment.Grade)
	} else {
		submission.Score = request.Score
	}

	submission.AttemptNumber = int(request.AttemptNumber)
	submission.ItemSnapshot = datatypes.JSON(snapshot)
	submission.ClientResult = datatypes.JSON(clientResult)
	submission.ServerResult = datatypes.JSON(serverResult)
	submission.IsVerified = request.IsVerified
	submission.DurationMS = request.DurationMS
	submission.FeedBack = request.FeedBack
	submission.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&submission).Error; err != nil {
		return err
	}
	return nil
}

func (r *submissionRepository) GetAllSubmissionByAssignmentID(ctx context.Context, ownerID, assignmentID int) (*[]types.SubmissionResponse, error) {
	var submission []model.Submission
	err := r.db.WithContext(ctx).Where("assignment_id = ?", assignmentID).Find(&submission).Error
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", ownerID).First(&user).Error; err != nil {
		return nil, err
	}

	// อนุญาตให้แก้ไขได้ทุกฟิลด์ เจ้าของ submission และ owner ของ assignment และ member ที่เป็น TA เท่านั้น
	// Select * FROM classroom c JOIN assignment a ON a.class_id = c.id JOIN member m ON m.class_id = c.id WHERE a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?))
	err = r.db.WithContext(ctx).Table("classroom c").Select("c.id, c.owner_id").Joins("JOIN assignment a ON a.class_id = c.id").Joins("JOIN member m ON m.class_id = c.id").Where("a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?))", assignmentID, ownerID, ownerID, "ta").First(&model.Classroom{}).Error
	if err != nil {
		return nil, err
	}

	submissionResponse := make([]types.SubmissionResponse, 0, len(submission))

	for _, sub := range submission {
		var ItemSnapshot map[string]interface{}
		if len(sub.ItemSnapshot) > 0 {
			if err := json.Unmarshal(sub.ItemSnapshot, &ItemSnapshot); err != nil {
				return nil, err
			}
		} else {
			ItemSnapshot = map[string]interface{}{}
		}

		var ClientResult map[string]interface{}
		if len(sub.ClientResult) > 0 {
			if err := json.Unmarshal(sub.ClientResult, &ClientResult); err != nil {
				return nil, err
			}
		} else {
			ClientResult = map[string]interface{}{}
		}

		var ServerResult map[string]interface{}
		if len(sub.ServerResult) > 0 {
			if err := json.Unmarshal(sub.ServerResult, &ServerResult); err != nil {
				return nil, err
			}
		} else {
			ServerResult = map[string]interface{}{}
		}

		submissionResponse = append(submissionResponse, types.SubmissionResponse{
			ID:            int(sub.ID),
			AssignmentID:  sub.AssignmentID,
			UserID:        sub.UserID,
			PlaygroundID:  sub.PlaygroundID,
			AttemptNumber: sub.AttemptNumber,
			ItemSnapshot:  ItemSnapshot,
			ClientResult:  ClientResult,
			ServerResult:  ServerResult,
			Score:         sub.Score,
			Status:        sub.Status,
			IsVerified:    sub.IsVerified,
			DurationMS:    sub.DurationMS,
		})
	}
	return &submissionResponse, nil
}

func (r *submissionRepository) GetSubmissionByID(ctx context.Context, userID, submissionID int) (*types.SubmissionResponse, error) {
	var submission model.Submission
	err := r.db.WithContext(ctx).Where("id = ?", submissionID).First(&submission).Error
	if err != nil {
		return nil, err
	}

	// อนุญาตให้แก้ไขได้ทุกฟิลด์ เจ้าของ submission และ owner ของ assignment และ member ที่เป็น TA เท่านั้น
	// Select * FROM classroom c JOIN assignment a ON a.class_id = c.id JOIN member m ON m.class_id = c.id WHERE a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?))
	err = r.db.WithContext(ctx).Table("classroom c").Select("c.id, c.owner_id").Joins("JOIN assignment a ON a.class_id = c.id").Joins("JOIN member m ON m.class_id = c.id").Joins("JOIN submission s ON s.assignment_id = a.id").Where("a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?) OR s.user_id = ?)", submission.AssignmentID, userID, userID, "ta", userID).First(&model.Classroom{}).Error
	if err != nil {
		return nil, err
	}

	var ItemSnapshot map[string]interface{}
	if len(submission.ItemSnapshot) > 0 {
		if err := json.Unmarshal(submission.ItemSnapshot, &ItemSnapshot); err != nil {
			return nil, err
		}
	} else {
		ItemSnapshot = map[string]interface{}{}
	}

	var ClientResult map[string]interface{}
	if len(submission.ClientResult) > 0 {
		if err := json.Unmarshal(submission.ClientResult, &ClientResult); err != nil {
			return nil, err
		}
	} else {
		ClientResult = map[string]interface{}{}
	}

	var ServerResult map[string]interface{}
	if len(submission.ServerResult) > 0 {
		if err := json.Unmarshal(submission.ServerResult, &ServerResult); err != nil {
			return nil, err
		}
	} else {
		ServerResult = map[string]interface{}{}
	}

	submissionResponse := &types.SubmissionResponse{
		ID:            int(submission.ID),
		AssignmentID:  submission.AssignmentID,
		UserID:        submission.UserID,
		PlaygroundID:  submission.PlaygroundID,
		AttemptNumber: submission.AttemptNumber,
		ItemSnapshot:  ItemSnapshot,
		ClientResult:  ClientResult,
		ServerResult:  ServerResult,
		Score:         submission.Score,
		Status:        submission.Status,
		IsVerified:    submission.IsVerified,
		DurationMS:    submission.DurationMS,
	}
	return submissionResponse, nil
}

func (r *submissionRepository) GetAllSubmissionByAssignmentIDandUserID(ctx context.Context, assignmentID, userID int) (*[]types.SubmissionResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	var submission []model.Submission
	err := r.db.WithContext(ctx).Where("assignment_id = ? AND user_id = ?", assignmentID, userID).Find(&submission).Error
	if err != nil {
		return nil, err
	}

	submissionResponse := make([]types.SubmissionResponse, 0, len(submission))

	for _, sub := range submission {

		var ItemSnapshot map[string]interface{}
		if len(sub.ItemSnapshot) > 0 {
			if err := json.Unmarshal(sub.ItemSnapshot, &ItemSnapshot); err != nil {
				return nil, err
			}
		} else {
			ItemSnapshot = map[string]interface{}{}
		}

		var ClientResult map[string]interface{}
		if len(sub.ClientResult) > 0 {
			if err := json.Unmarshal(sub.ClientResult, &ClientResult); err != nil {
				return nil, err
			}
		} else {
			ClientResult = map[string]interface{}{}
		}

		var ServerResult map[string]interface{}
		if len(sub.ServerResult) > 0 {
			if err := json.Unmarshal(sub.ServerResult, &ServerResult); err != nil {
				return nil, err
			}
		} else {
			ServerResult = map[string]interface{}{}
		}

		submissionResponse = append(submissionResponse, types.SubmissionResponse{
			ID:            int(sub.ID),
			AssignmentID:  sub.AssignmentID,
			UserID:        sub.UserID,
			PlaygroundID:  sub.PlaygroundID,
			AttemptNumber: sub.AttemptNumber,
			ItemSnapshot:  ItemSnapshot,
			ClientResult:  ClientResult,
			ServerResult:  ServerResult,
			Score:         sub.Score,
			Status:        sub.Status,
			IsVerified:    sub.IsVerified,
			DurationMS:    sub.DurationMS,
		})
	}
	return &submissionResponse, nil
}

func (r *submissionRepository) UpdateGrade(ctx context.Context, userID, submissionID int, request types.UpdateGradeRequest) error {
	var submission model.Submission
	err := r.db.WithContext(ctx).Where("id = ?", submissionID).First(&submission).Error
	if err != nil {
		return err
	}

	var assignment model.Assignment
	err = r.db.WithContext(ctx).Where("id = ?", submission.AssignmentID).First(&assignment).Error
	if err != nil {
		return err
	}

	// อนุญาตให้แก้ไขได้ทุกฟิลด์ เจ้าของ submission และ owner ของ assignment และ member ที่เป็น TA เท่านั้น
	err = r.db.WithContext(ctx).Table("classroom c").Select("c.id, c.owner_id").Joins("JOIN assignment a ON a.class_id = c.id").Joins("JOIN member m ON m.class_id = c.id").Joins("JOIN submission s ON s.assignment_id = a.id").Where("a.id = ? AND (c.owner_id = ? OR (m.user_id = ? AND m.role = ?) OR s.user_id = ?)", submission.AssignmentID, userID, userID, "ta", userID).First(&model.Classroom{}).Error
	if err != nil {
		return err
	}

	// ถ้า score ที่ส่งมามากกว่า grade ของ assignment ให้ตั้งเป็น grade ของ assignment แทน และ ถ้าไม่ก็ให้ตั้งเป็น score ที่ส่งมา
	if request.Score > float64(assignment.Grade) {
		submission.Score = float64(assignment.Grade)
	} else {
		submission.Score = request.Score
	}
	submission.FeedBack = request.FeedBack
	submission.IsVerified = request.IsVerified
	submission.UpdatedAt = time.Now()

	submission.Status = "submitted"
	if err := r.db.WithContext(ctx).Save(&submission).Error; err != nil {
		submission.Status = "failed"
		if err := r.db.WithContext(ctx).Save(&submission).Error; err != nil {
			return err
		}
		return err
	}
	return nil
}
