package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"google.golang.org/api/classroom/v1"
)

type ClassroomUseCase interface {
	ListCourseUsecase(ctx context.Context, accessToken string) ([]*classroom.Course, error)
	ListStudentsUsecase(ctx context.Context, accessToken string, courseId string) ([]*classroom.Student, error)
	GetAllAssignmentsUsecase(ctx context.Context, accessToken string, courseId string) ([]*classroom.CourseWork, error)
}

type classroomUseCase struct {
	repo repository.ClassroomRepository
}

func NewClassroomUseCase(repo repository.ClassroomRepository) ClassroomUseCase {
	return &classroomUseCase{
		repo: repo,
	}
}

func (c *classroomUseCase) ListCourseUsecase(ctx context.Context, accessToken string) ([]*classroom.Course, error) {
	courses, err := c.repo.ListCourse(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *classroomUseCase) ListStudentsUsecase(ctx context.Context, accessToken string, courseId string) ([]*classroom.Student, error) {
	students, err := c.repo.ListStudents(ctx, accessToken, courseId)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (c *classroomUseCase) GetAllAssignmentsUsecase(ctx context.Context, accessToken string, courseId string) ([]*classroom.CourseWork, error) {
	assignments, err := c.repo.GetAllAssignments(ctx, accessToken, courseId)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}