package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type AssignmentUseCase interface {
	// Define methods related to Assignment use cases here
	GetAssignmentsByClassIDUseCases(ctx context.Context, classID int) (*[]types.AssignmentResponse, error)
	CreateAssignmentUseCases(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) (int, error)
	GetAssignmentsByAssignmentIDUseCases(ctx context.Context, classID, assignmentID int) (*types.AssignmentResponse, error)
	EdiitAssignmentByAssignmentIDUseCases(ctx context.Context, owner, classID, assignmentID int, assignment *types.EditAssignmentRequest) error
	DeleteAssignmentByAssignmentIDUseCases(ctx context.Context, owner, classID, assignmentID int) error
}

type assignmentUseCase struct {
	// Add necessary repositories here
	assignmentRepo repository.AssignmentRepository
}

func NewAssignmentUseCase(assignmentRepo repository.AssignmentRepository) AssignmentUseCase {
	return &assignmentUseCase{assignmentRepo: assignmentRepo}
}

func (uc *assignmentUseCase) GetAssignmentsByClassIDUseCases(ctx context.Context, classID int) (*[]types.AssignmentResponse, error) {
	resp, err := uc.assignmentRepo.GetAssignmentsByClassID(ctx, classID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *assignmentUseCase) CreateAssignmentUseCases(ctx context.Context, owner int, classID int, assignment *types.CreateAssignmentRequest) (int, error) {
	id, err := uc.assignmentRepo.CreateAssignment(ctx, owner, classID, assignment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *assignmentUseCase) GetAssignmentsByAssignmentIDUseCases(ctx context.Context, classID, assignmentID int) (*types.AssignmentResponse, error) {
	resp, err := uc.assignmentRepo.GetAssignmentsByAssignmentID(ctx, classID, assignmentID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *assignmentUseCase) EdiitAssignmentByAssignmentIDUseCases(ctx context.Context, owner, classID, assignmentID int, assignment *types.EditAssignmentRequest) error {
	err := uc.assignmentRepo.EditAssignmentByAssignmentID(ctx, owner, classID, assignmentID, assignment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *assignmentUseCase) DeleteAssignmentByAssignmentIDUseCases(ctx context.Context, owner, classID, assignmentID int) error {
	err := uc.assignmentRepo.DeleteAssignmentByAssignmentID(ctx, owner, classID, assignmentID)
	if err != nil {
		return err
	}
	return nil
}
