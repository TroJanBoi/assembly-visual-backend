package repository

import (
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type OperationRepository interface {
	OperationAdd([]float64) (*types.OperationResponse, error)
	OperationSub([]float64) (*types.OperationResponse, error)
	OperationMul([]float64) (*types.OperationResponse, error)
	OperationDiv([]float64) (*types.OperationResponse, error)
}

type operationRepository struct {
	db *gorm.DB
}

func NewOperationRepository(db *gorm.DB) OperationRepository {
	return &operationRepository{db: db}
}

func (r *operationRepository) OperationAdd(values []float64) (*types.OperationResponse, error) {
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

func (o *operationRepository) OperationSub(values []float64) (*types.OperationResponse, error) {
	sum := values[0]
	for _, v := range values[1:] {
		sum -= v
	}
	result := &types.OperationResponse{
		Value:      sum,
		StatusCode: 200,
		Message:    "Subtraction successful",
	}
	return result, nil
}

func (o *operationRepository) OperationMul(values []float64) (*types.OperationResponse, error) {
	sum := values[0]
	for _, v := range values[1:] {
		sum *= v
	}
	result := &types.OperationResponse{
		Value:      sum,
		StatusCode: 200,
		Message:    "Multiplication successful",
	}
	return result, nil
}

func (o *operationRepository) OperationDiv(values []float64) (*types.OperationResponse, error) {
	sum := values[0]
	for _, v := range values[1:] {
		if v == 0 {
			return &types.OperationResponse{
				Value:      0,
				StatusCode: 400,
				Message:    "Division by zero is not allowed",
			}, nil
		}
		sum /= v
	}
	result := &types.OperationResponse{
		Value:      sum,
		StatusCode: 200,
		Message:    "Division successful",
	}
	return result, nil
}
