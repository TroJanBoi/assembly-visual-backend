package usecases

import (
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type OperationUsecase interface {
	OperationAdd([]float64) (*types.OperationResponse, error)
}

type operationUsecase struct {
	operationRepository repository.OperationRepository
}

func NewOperationUsecase(operationRepository repository.OperationRepository) OperationUsecase {
	return &operationUsecase{
		operationRepository: operationRepository,
	}
}

func (u *operationUsecase) OperationAdd(values []float64) (*types.OperationResponse, error) {
	operation, err := u.operationRepository.OperationAdd(values)
	if err != nil {
		return nil, err
	}
	return operation, nil
}
