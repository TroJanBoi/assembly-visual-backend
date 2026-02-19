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
	CreateClassUseCases(ctx context.Context, owner int, class *types.CreateClassRequest) error
	UpdateClassUseCases(ctx context.Context, owner int, classID int, class *types.UpdateClassRequest) error
	DeleteClassUseCases(ctx context.Context, owner int, classID int) error
	JoinClassUseCases(ctx context.Context, userID, classID int) error
	GetAllMembersByClassID(ctx context.Context, classID int) (*[]types.MemberResponse, error)
	GetAllClassPublicUseCases(ctx context.Context) (*[]types.ClassResponse, error)
	ChangePermissionMemberUseCases(ctx context.Context, userID, classID int, newRole string) error
	RemoveMemberInClassUseCases(ctx context.Context, classID, userID int) error
	GetClassRecentManyIDsUseCases(ctx context.Context, limit []int) (*[]types.ClassResponse, error)
	JoinWithCodeUseCases(ctx context.Context, userID int, code string) error
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

func (uc *classUseCase) CreateClassUseCases(ctx context.Context, owner int, class *types.CreateClassRequest) error {
	err := uc.classRepo.CreateClass(ctx, owner, class)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) UpdateClassUseCases(ctx context.Context, owner int, classID int, class *types.UpdateClassRequest) error {
	err := uc.classRepo.UpdateClass(ctx, owner, classID, class)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) DeleteClassUseCases(ctx context.Context, owner int, classID int) error {
	err := uc.classRepo.DeleteClass(ctx, owner, classID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) JoinClassUseCases(ctx context.Context, userID, classID int) error {
	err := uc.classRepo.JoinClass(ctx, userID, classID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) GetAllMembersByClassID(ctx context.Context, classID int) (*[]types.MemberResponse, error) {
	resp, err := uc.classRepo.GetAllMembersByClassID(ctx, classID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *classUseCase) GetAllClassPublicUseCases(ctx context.Context) (*[]types.ClassResponse, error) {
	resp, err := uc.classRepo.GetAllClassPublic(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *classUseCase) ChangePermissionMemberUseCases(ctx context.Context, userID, classID int, newRole string) error {
	err := uc.classRepo.ChangePermissionMember(ctx, userID, classID, newRole)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) RemoveMemberInClassUseCases(ctx context.Context, classID, userID int) error {
	err := uc.classRepo.RemoveMemberInClass(ctx, classID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *classUseCase) GetClassRecentManyIDsUseCases(ctx context.Context, limit []int) (*[]types.ClassResponse, error) {
	resp, err := uc.classRepo.GetClassRecentManyIDs(ctx, limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *classUseCase) JoinWithCodeUseCases(ctx context.Context, userID int, code string) error {
	err := uc.classRepo.JoinWithCode(ctx, userID, code)
	if err != nil {
		return err
	}
	return nil
}
