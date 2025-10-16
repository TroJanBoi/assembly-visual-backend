package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type ClassUseCase interface {
	// Define methods related to Class use cases here
	GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error)
	GetClassByIDUseCases(ctx context.Context, classID int) (*types.ClassResponse, error)
	CreateClassUseCases(ctx context.Context, owner uint, class *types.CreateClassRequest) error
	UpdateClassUseCases(ctx context.Context, owner uint, classID int, class *types.UpdateClassRequest) error
	DeleteClassUseCases(ctx context.Context, owner uint, classID int) error
}

type classUseCase struct {
	classRepo repository.ClassRepository
}

func NewClassUseCase(classRepo repository.ClassRepository) ClassUseCase {
	return &classUseCase{classRepo: classRepo}
}

func (uc *classUseCase) GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error) {
	resp, err := uc.classRepo.GetAllClasses(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *classUseCase) GetClassByIDUseCases(ctx context.Context, classID int) (*types.ClassResponse, error) {
	resp, err := uc.classRepo.GetClassByID(ctx, classID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *classUseCase) CreateClassUseCases(ctx context.Context, owner uint, class *types.CreateClassRequest) error {
	err := uc.classRepo.CreateClass(ctx, owner, class)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) UpdateClassUseCases(ctx context.Context, owner uint, classID int, class *types.UpdateClassRequest) error {
	err := uc.classRepo.UpdateClass(ctx, owner, classID, class)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) DeleteClassUseCases(ctx context.Context, owner uint, classID int) error {
	err := uc.classRepo.DeleteClass(ctx, owner, classID)
	if err != nil {
		return err
	}
	return nil
}
