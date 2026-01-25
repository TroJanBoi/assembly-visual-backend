package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GoogleServiceRepository interface {
	// Define methods for Google Service interactions here
	ListGoogleClassroomCourses(ctx context.Context, accessToken string) ([]byte, error)
	AssignmentGoogleClassroom(ctx context.Context, accessToken string, courseID string) ([]byte, error)
	ConfirmGoogleClassroomConnection(ctx context.Context, accessToken string, courseID string, ownerID int) error
}

type googleServiceRepository struct {
	db *gorm.DB
}

func NewGoogleServiceRepository(db *gorm.DB) GoogleServiceRepository {
	return &googleServiceRepository{db: db}
}

func (r *googleServiceRepository) ListGoogleClassroomCourses(ctx context.Context, accessToken string) ([]byte, error) {
	request, _ := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://classroom.googleapis.com/v1/courses?teacherId=me",
		nil,
	)

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to list courses: status=%d body=%s", response.StatusCode, string(body))
	}

	return io.ReadAll(response.Body)
}

func (r *googleServiceRepository) AssignmentGoogleClassroom(ctx context.Context, accessToken string, courseID string) ([]byte, error) {
	request, _ := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://classroom.googleapis.com/v1/courses/"+courseID+"/courseWork",
		nil,
	)

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to list assignments: status=%d body=%s", response.StatusCode, string(body))
	}

	return io.ReadAll(response.Body)
}

func (r *googleServiceRepository) ConfirmGoogleClassroomConnection(ctx context.Context, accessToken string, courseID string, ownerID int) error {
	request, _ := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://classroom.googleapis.com/v1/courses/"+courseID,
		nil,
	)

	request.Header.Set("Authorization", "Bearer "+accessToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get course details: status=%d body=%s", response.StatusCode, string(responseBody))
	}

	courseData := types.CourseData{}
	if err := json.Unmarshal(responseBody, &courseData); err != nil {
		return fmt.Errorf("failed to unmarshal course data: %w", err)
	}

	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", ownerID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	newClass := model.Classroom{
		GoogleCourseID:   courseID,
		OwnerId:          ownerID,
		Status:           1,
		Topic:            courseData.Name,
		GoogleCourseLink: courseData.AlternateLink,
		GoogleSyncedAt:   time.Now().Format(time.RFC3339),
		Code:             courseData.EnrollmentCode,
	}

	if err := r.db.WithContext(ctx).Create(&newClass).Error; err != nil {
		return fmt.Errorf("failed to create classroom record: %w", err)
	}

	newClassSyncLog := model.GoogleCourseSyncLog{
		ClassID:  int(newClass.ID),
		Action:   "Create",
		Response: datatypes.JSON([]byte(responseBody)),
		Status:   "success",
	}

	if err := r.db.WithContext(ctx).Create(&newClassSyncLog).Error; err != nil {
		return fmt.Errorf("failed to create google course sync log: %w", err)
	}

	requestAssignments, _ := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://classroom.googleapis.com/v1/courses/"+courseID+"/courseWork",
		nil,
	)

	requestAssignments.Header.Set("Authorization", "Bearer "+accessToken)

	responseAssignments, err := http.DefaultClient.Do(requestAssignments)
	if err != nil {
		return err
	}
	defer responseAssignments.Body.Close()

	if responseAssignments.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(responseAssignments.Body)
		return fmt.Errorf("failed to list assignments: status=%d body=%s", responseAssignments.StatusCode, string(body))
	}

	assignmentsBody, err := io.ReadAll(responseAssignments.Body)
	if err != nil {
		return fmt.Errorf("failed to read assignments response body: %w", err)
	}

	assignmentsResponse := types.CourseWorkListResponse{}
	if err := json.Unmarshal(assignmentsBody, &assignmentsResponse); err != nil {
		return fmt.Errorf("failed to unmarshal assignments data: %w", err)
	}

	for _, assignment := range assignmentsResponse.CourseWork {

		newAssignment := model.Assignment{
			ClassID:     int(newClass.ID),
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     time.Date(assignment.DueDate.Year, time.Month(assignment.DueDate.Month), assignment.DueDate.Day, 0, 0, 0, 0, time.UTC),
			MaxAttempt:  int(assignment.MaxPoints),
			Setting:     nil,
			Condition:   nil,
			Grade:       0,
		}

		if err := r.db.WithContext(ctx).Create(&newAssignment).Error; err != nil {
			return fmt.Errorf("failed to create assignment record: %w", err)
		}
	}
	return nil
}
