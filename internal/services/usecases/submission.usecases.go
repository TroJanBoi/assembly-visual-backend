package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type SubmissionUseCase interface {
	CreateSubmissionUseCase(ctx context.Context, userId int, request types.CreateSubmissionRequest) error
	UpdateSubmissionUseCase(ctx context.Context, userID, submissionID int, request types.UpdateSubmissionRequest) error
	GetAllSubmissionByAssignmentIDUseCase(ctx context.Context, ownerID, assignmentID int) (*[]types.SubmissionResponse, error)
	GetSubmissionByIDUseCase(ctx context.Context, userID, submissionID int) (*types.SubmissionResponse, error)
	GetAllSubmissionByAssignmentIDandUserIDUseCase(ctx context.Context, assignmentID, userID int) (*[]types.SubmissionResponse, error)
	UpdateGradeUseCase(ctx context.Context, userID, submissionID int, request types.UpdateGradeRequest) error
}

type submissionUseCase struct {
	submissionRepo repository.SubmissionRepository
}

func NewSubmissionUseCase(submissionRepo repository.SubmissionRepository) SubmissionUseCase {
	return &submissionUseCase{submissionRepo: submissionRepo}
}

func (uc *submissionUseCase) CreateSubmissionUseCase(ctx context.Context, userId int, request types.CreateSubmissionRequest) error {
	err := uc.submissionRepo.CreateSubmission(ctx, userId, request)
	if err != nil {
		return err
	}
	return nil
}

func (uc *submissionUseCase) UpdateSubmissionUseCase(ctx context.Context, userID, submissionID int, request types.UpdateSubmissionRequest) error {
	err := uc.submissionRepo.UpdateSubmission(ctx, userID, submissionID, request)
	if err != nil {
		return err
	}
	return nil
}

func (uc *submissionUseCase) GetAllSubmissionByAssignmentIDUseCase(ctx context.Context, ownerID, assignmentID int) (*[]types.SubmissionResponse, error) {
	resp, err := uc.submissionRepo.GetAllSubmissionByAssignmentID(ctx, ownerID, assignmentID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *submissionUseCase) GetSubmissionByIDUseCase(ctx context.Context, userID, submissionID int) (*types.SubmissionResponse, error) {
	resp, err := uc.submissionRepo.GetSubmissionByID(ctx, userID, submissionID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *submissionUseCase) GetAllSubmissionByAssignmentIDandUserIDUseCase(ctx context.Context, assignmentID, userID int) (*[]types.SubmissionResponse, error) {
	resp, err := uc.submissionRepo.GetAllSubmissionByAssignmentIDandUserID(ctx, assignmentID, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *submissionUseCase) UpdateGradeUseCase(ctx context.Context, userID, submissionID int, request types.UpdateGradeRequest) error {
	err := uc.submissionRepo.UpdateGrade(ctx, userID, submissionID, request)
	if err != nil {
		return err
	}
	return nil
}
