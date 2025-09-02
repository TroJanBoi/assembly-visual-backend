package repository

import (
	"context"
	"fmt"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"golang.org/x/oauth2"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

type ClassroomRepository interface {
	ListCourse(ctx context.Context, accessToken string) ([]*classroom.Course, error)
	ListStudents(ctx context.Context, accessToken string, courseId string) ([]*classroom.Student, error)
	GetAllAssignments(ctx context.Context, accessToken string, courseId string) ([]*classroom.CourseWork, error)
}

type classroomRepository struct{}

func NewClassroomRepository() ClassroomRepository {
	return &classroomRepository{}
}

func (r *classroomRepository) svc(ctx context.Context, accessToken string) (*classroom.Service, error) {
	ts := conf.GetGoogleOAuthConfig().TokenSource(ctx, &oauth2.Token{AccessToken: accessToken})
	return classroom.NewService(ctx, option.WithTokenSource(ts))
}

func (c *classroomRepository) ListCourse(ctx context.Context, accessToken string) ([]*classroom.Course, error) {
	service, err := c.svc(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create classroom service: %w", err)
	}
	response, err := service.Courses.List().PageSize(50).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to list classrooms: %w", err)
	}
	return response.Courses, nil
}

func (c *classroomRepository) ListStudents(ctx context.Context, accessToken, courseId string) ([]*classroom.Student, error) {
	service, err := c.svc(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create classroom service: %w", err)
	}

	var allStudents []*classroom.Student
	pageToken := ""
	for {
		call := service.Courses.Students.List(courseId).PageSize(100)
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}
		res, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("failed to list students: %w", err)
		}

		allStudents = append(allStudents, res.Students...)

		if res.NextPageToken == "" {
			break
		}
		pageToken = res.NextPageToken
	}

	return allStudents, nil
}

func (c *classroomRepository) GetAllAssignments(ctx context.Context, accessToken, courseId string) ([]*classroom.CourseWork, error) {
	service, err := c.svc(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create classroom service: %w", err)
	}
	var GetAllAssignments []*classroom.CourseWork
	pageToken := ""
	for {
		assignment := service.Courses.CourseWork.List(courseId).PageSize(100)
		if pageToken != "" {
			assignment = assignment.PageToken(pageToken)
		}

		response, err := assignment.Do()
		if err != nil {
			return nil, fmt.Errorf("failed to list assignments: %w", err)
		}
		GetAllAssignments = append(GetAllAssignments, response.CourseWork...)
		if response.NextPageToken == "" {
			break
		}
		pageToken = response.NextPageToken
	}
	return GetAllAssignments, nil
}
