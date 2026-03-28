package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/security"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) (*[]types.UserResponse, error)
	CreateUser(ctx context.Context, users *types.CreateUserRequest) error
	UpdateUser(ctx context.Context, userID int, users *types.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userID int) error
	GetMeClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error)
	GetOwnerClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error)
	GetMeTask(ctx context.Context, userID int) (*[]types.TaskMeResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpdateUser(ctx context.Context, userID int, users *types.UpdateUserRequest) error {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return fmt.Errorf("user with ID %d does not exist", userID)
	}

	updatedData := make(map[string]interface{})
	if users.Password != "" {
		updatedData["password"] = security.HashPassword(users.Password)
	}
	if users.Name != "" {
		updatedData["name"] = users.Name
	}

	if err := r.db.WithContext(ctx).Model(&user).Updates(updatedData).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *userRepository) CreateUser(ctx context.Context, users *types.CreateUserRequest) error {
	var user model.User

	err := r.db.WithContext(ctx).Where("email = ?", users.Email).First(&user).Error
	if err == nil {
		return fmt.Errorf("user with email %s already exists", users.Email)
	}

	if users.Name == "" {
		users.Name = ""
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := model.User{
			Email:        users.Email,
			PasswordHash: security.HashPassword(users.Password),
			Name:         users.Name,
		}
		if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	}
	return fmt.Errorf("failed to create user: %w", err)
}

func (r *userRepository) DeleteUser(ctx context.Context, userID int) error {
	var existingUser model.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&existingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d does not exist", userID)
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	picture_path := existingUser.PicturePath
	if picture_path != "" {
		fileName := filepath.Base(picture_path)
		filePath := filepath.Join("uploads/users", fileName)

		if err := os.Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to delete old avatar file: %w", err)
		}
	}

	if err := r.db.WithContext(ctx).Delete(&existingUser).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) (*[]types.UserResponse, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	var userResp []types.UserResponse
	for _, user := range users {
		userResp = append(userResp, types.UserResponse{
			ID:           int(user.ID),
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Name:         user.Name,
			PicturePath:  user.PicturePath,
		})
	}
	return &userResp, nil
}

// Get me's classes joined
func (r *userRepository) GetMeClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error) {
	var member []model.Member
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&member).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve member records: %w", err)
	}

	var classIDs []int
	for _, m := range member {
		classIDs = append(classIDs, int(m.ClassID))
	}

	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve classes: %w", err)
	}

	var user model.User

	var classResp []types.ClassMeResponse
	for _, class := range classes {
		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to retrieve user: %w", err)
		}

		var countMember int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&countMember).Error; err != nil {
			return nil, fmt.Errorf("failed to count members: %w", err)
		}

		classResp = append(classResp, types.ClassMeResponse{
			ID:               int(class.ID),
			Code:             class.Code,
			Topic:            class.Topic,
			Description:      class.Description,
			OwnerID:          int(class.OwnerId),
			OwnerName:        user.Name,
			Status:           int(class.Status),
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			BannerID:         int(class.BannerID),
			MemberAmount:     countMember,
		})
	}
	return &classResp, nil
}

func (r *userRepository) GetOwnerClass(ctx context.Context, userID int) (*[]types.ClassMeResponse, error) {
	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("owner_id = ?", userID).Find(&classes).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve classes: %w", err)
	}

	var user model.User

	var classResp []types.ClassMeResponse
	for _, class := range classes {

		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to retrieve user: %w", err)
		}

		var countMember int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&countMember).Error; err != nil {
			return nil, fmt.Errorf("failed to count members: %w", err)
		}

		classResp = append(classResp, types.ClassMeResponse{
			ID:               int(class.ID),
			Code:             class.Code,
			Topic:            class.Topic,
			Description:      class.Description,
			OwnerID:          int(class.OwnerId),
			OwnerName:        user.Name,
			Status:           int(class.Status),
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			BannerID:         int(class.BannerID),
			MemberAmount:     countMember,
		})
	}
	return &classResp, nil
}

func (r *userRepository) GetMeTask(ctx context.Context, userID int) (*[]types.TaskMeResponse, error) {
	var member []model.Member // check member
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&member).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve member records: %w", err)
	}

	var classIDs []int
	for _, m := range member {
		classIDs = append(classIDs, int(m.ClassID))
	}

	var classes []model.Classroom // select classes where user is member
	if err := r.db.WithContext(ctx).Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve classes: %w", err)
	}

	// get class id list for select tasks
	var classIDList []int
	for _, class := range classes {
		classIDList = append(classIDList, int(class.ID))
	}

	// select tasks where class id in classIDList
	var assignments []model.Assignment
	if err := r.db.WithContext(ctx).Where("class_id IN ?", classIDList).Find(&assignments).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve assignments: %w", err)
	}

	var submissions []model.Submission
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&submissions).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve submissions: %w", err)
	}

	var status string
	var taskResp []types.TaskMeResponse
	for _, assignment := range assignments {

		var countSubmissionByAssignmentID int64
		if err := r.db.WithContext(ctx).Model(&model.Submission{}).Where("user_id = ? AND assignment_id = ?", userID, assignment.ID).Count(&countSubmissionByAssignmentID).Error; err != nil {
			return nil, fmt.Errorf("failed to count submissions: %w", err)
		}

		if time.Now().After(assignment.DueDate) {
			status = "overdue"
		} else if countSubmissionByAssignmentID > 0 {
			status = "completed"
		} else {
			status = "in_progress"
		}

		taskResp = append(taskResp, types.TaskMeResponse{
			ClassID:         int(assignment.ClassID),
			AssignmentID:    int(assignment.ID),
			AssignmentTitle: assignment.Title,
			Description:     assignment.Description,
			MaxAttempt:      assignment.MaxAttempt,
			DueDate:         assignment.DueDate.Format((time.RFC3339)),
			Status:          status,
		})
	}
	return &taskResp, nil
}
