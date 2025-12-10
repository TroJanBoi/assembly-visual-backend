package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type PlaygroundUsecase interface {
	CreateUseCases(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error)
	GetByPlaygroundIDUseCases(ctx context.Context, userID int, id int) (*types.PlaygroundResponse, error)
	GetPlaygroundByMeUseCases(ctx context.Context, userID int, req *types.PlaygroundMeRequest) (*types.PlaygroundResponse, error)
	UpdatePlaygroundByMeUseCases(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error)
	DeletePlaygroundByMeUseCases(ctx context.Context, userID int, req *types.PlaygroundMeRequest) error
}

type playgroundUsecase struct {
	playgroundRepo repository.PlaygroundRepository
}

func NewPlaygroundUseCases(playgroundRepo repository.PlaygroundRepository) PlaygroundUsecase {
	return &playgroundUsecase{playgroundRepo: playgroundRepo}
}

func (u *playgroundUsecase) CreateUseCases(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error) {
	playgroundResponse, err := u.playgroundRepo.Create(ctx, userID, playground)
	if err != nil {
		return nil, err
	}
	return playgroundResponse, nil
}

func (u *playgroundUsecase) GetByPlaygroundIDUseCases(ctx context.Context, userID int, id int) (*types.PlaygroundResponse, error) {
	playgroundResponse, err := u.playgroundRepo.GetByPlaygroundID(ctx, userID, id)
	if err != nil {
		return nil, err
	}
	return playgroundResponse, nil
}

func (u *playgroundUsecase) GetPlaygroundByMeUseCases(ctx context.Context, userID int, req *types.PlaygroundMeRequest) (*types.PlaygroundResponse, error) {
	playgroundResponse, err := u.playgroundRepo.GetPlaygroundByMe(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return playgroundResponse, nil
}

func (u *playgroundUsecase) UpdatePlaygroundByMeUseCases(ctx context.Context, userID int, playground *types.PlaygroundRequest) (*types.PlaygroundResponse, error) {
	playgroundResponse, err := u.playgroundRepo.UpdatePlaygroundByMe(ctx, userID, playground)
	if err != nil {
		return nil, err
	}
	return playgroundResponse, nil
}

func (u *playgroundUsecase) DeletePlaygroundByMeUseCases(ctx context.Context, userID int, req *types.PlaygroundMeRequest) error {
	err := u.playgroundRepo.DeletePlaygroundByMe(ctx, userID, req)
	if err != nil {
		return err
	}
	return nil
}
