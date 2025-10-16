package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type AssignmentUseCase interface {
	// Define methods related to Assignment use cases here
	GetAssignmentsByClassIDUseCases(ctx context.Context, owner int, classID int) (*[]types.AssignmentResponse, error)
}

type assignmentUseCase struct {
	// Add necessary repositories here
	assignmentRepo repository.AssignmentRepository
}

func NewAssignmentUseCase(assignmentRepo repository.AssignmentRepository) AssignmentUseCase {
	return &assignmentUseCase{assignmentRepo: assignmentRepo}
}

func (uc *assignmentUseCase) GetAssignmentsByClassIDUseCases(ctx context.Context, owner int, classID int) (*[]types.AssignmentResponse, error) {
	resp, err := uc.assignmentRepo.GetAssignmentsByClassID(ctx, owner, classID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
