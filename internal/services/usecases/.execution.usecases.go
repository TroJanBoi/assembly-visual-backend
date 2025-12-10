package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type ExecutionUsecase interface {
	ExecutionPlaygroundUseCases(ctx context.Context, userID int, playgroundID int) (*types.ExecutionState, error)
}

type executionUsecase struct {
	executionRepo repository.ExecutionRepository
}

func NewExecutionUseCases(executionRepo repository.ExecutionRepository) ExecutionUsecase {
	return &executionUsecase{executionRepo: executionRepo}
}

func (u *executionUsecase) ExecutionPlaygroundUseCases(ctx context.Context, userID int, playgroundID int) (*types.ExecutionState, error) {
	executionState, err := u.executionRepo.ExecutionPlayground(ctx, userID, playgroundID)
	if err != nil {
		return nil, err
	}
	return executionState, nil
}
