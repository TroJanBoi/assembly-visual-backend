package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
)

type GoogleServiceUsecase interface {
	ListGoogleClassroomCoursesUsecase(ctx context.Context, userID int) ([]byte, error)
	AssignmentGoogleClassroomUsecase(ctx context.Context, userID int, courseID string) ([]byte, error)
	ConfirmGoogleClassroomConnectionUsecase(ctx context.Context, userID int, courseID string, ownerID int) error
}

type googleServiceUsecase struct {
	googleServiceRepo repository.GoogleServiceRepository
	oauthRepository   repository.OAuthRepository
}

func NewGoogleServiceUsecase(googleServiceRepo repository.GoogleServiceRepository, oauthRepository repository.OAuthRepository) GoogleServiceUsecase {
	return &googleServiceUsecase{googleServiceRepo: googleServiceRepo, oauthRepository: oauthRepository}
}

func (uc *googleServiceUsecase) ListGoogleClassroomCoursesUsecase(ctx context.Context, userID int) ([]byte, error) {
	accessToken, err := uc.oauthRepository.RefreshGoogleToken(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp, err := uc.googleServiceRepo.ListGoogleClassroomCourses(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *googleServiceUsecase) AssignmentGoogleClassroomUsecase(ctx context.Context, userID int, courseID string) ([]byte, error) {
	accessToken, err := uc.oauthRepository.RefreshGoogleToken(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp, err := uc.googleServiceRepo.AssignmentGoogleClassroom(ctx, accessToken, courseID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *googleServiceUsecase) ConfirmGoogleClassroomConnectionUsecase(ctx context.Context, userID int, courseID string, ownerID int) error {
	accessToken, err := uc.oauthRepository.RefreshGoogleToken(ctx, userID)
	if err != nil {
		return err
	}

	err = uc.googleServiceRepo.ConfirmGoogleClassroomConnection(ctx, accessToken, courseID, ownerID)
	if err != nil {
		return err
	}
	return nil
}
