package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type UserUseCase interface {
	CreateUserUsecase(ctx context.Context, user *types.CreateUserRequest) error
	GetAllUsersUsecase(ctx context.Context) (*[]types.UserResponse, error)
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
