package repository

import (
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type OperationRepository interface {
	OperationAdd([]float64) (*types.OperationResponse, error)
}

type operationRepository struct {
	db *gorm.DB
}

func NewOperationRepository(db *gorm.DB) OperationRepository {
	return &operationRepository{db: db}
}

func (r *operationRepository) OperationAdd(values []float64) (*types.OperationResponse, error) {
	for i := range values {
		if strconv.Itoa(int(values[i])) == "" {
			return &types.OperationResponse{
				Value:      0.0,
				StatusCode: 400,
				Message:    "Invalid input: all elements must be numbers",
			}, nil
		}
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	result := &types.OperationResponse{
		Value:      sum,
		StatusCode: 200,
		Message:    "Addition successful",
	}
	return result, nil
}
