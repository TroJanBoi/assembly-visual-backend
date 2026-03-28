package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type UserUseCase interface {
	CreateUserUsecase(ctx context.Context, user *types.CreateUserRequest) error
	GetAllUsersUsecase(ctx context.Context) (*[]types.UserResponse, error)
	UpdateUsersUsecase(ctx context.Context, userID int, user *types.UpdateUserRequest) error
	DeleteUserUsecase(ctx context.Context, userID int) error
	GetMeClassUsecase(ctx context.Context, userID int) (*[]types.ClassMeResponse, error)
	GetOwnerClassUsecase(ctx context.Context, userID int) (*[]types.ClassMeResponse, error)
	GetMeTaskUsecase(ctx context.Context, userID int) (*[]types.TaskMeResponse, error)
}
type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) CreateUserUsecase(ctx context.Context, user *types.CreateUserRequest) error {
	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) GetAllUsersUsecase(ctx context.Context) (*[]types.UserResponse, error) {
	users, err := u.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) UpdateUsersUsecase(ctx context.Context, userID int, user *types.UpdateUserRequest) error {
	if err := u.userRepository.UpdateUser(ctx, userID, user); err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) DeleteUserUsecase(ctx context.Context, userID int) error {
	if err := u.userRepository.DeleteUser(ctx, userID); err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) GetMeClassUsecase(ctx context.Context, userID int) (*[]types.ClassMeResponse, error) {
	classes, err := u.userRepository.GetMeClass(ctx, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (u *userUseCase) GetOwnerClassUsecase(ctx context.Context, userID int) (*[]types.ClassMeResponse, error) {
	classes, err := u.userRepository.GetOwnerClass(ctx, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (u *userUseCase) GetMeTaskUsecase(ctx context.Context, userID int) (*[]types.TaskMeResponse, error) {
	tasks, err := u.userRepository.GetMeTask(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
