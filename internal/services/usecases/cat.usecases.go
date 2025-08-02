package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type CatUseCase interface {
	GetAllCatsUsecase(ctx context.Context) ([]types.CatResponse, error)
}

type catUseCase struct {
	catRepository repository.CatRepository
}

func NewCatUseCase(catRepository repository.CatRepository) CatUseCase {
	return &catUseCase{
		catRepository: catRepository,
	}
}

func (c *catUseCase) GetAllCatsUsecase(ctx context.Context) ([]types.CatResponse, error) {
	cats, err := c.catRepository.GetAllCats(ctx)
	if err != nil {
		return nil, err
	}
	return *cats, nil
}
