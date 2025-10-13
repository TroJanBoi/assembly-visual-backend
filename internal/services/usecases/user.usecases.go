package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type UserUseCase interface {
	CreateUserUsecase(ctx context.Context, email, password string) error
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

func (u *userUseCase) CreateUserUsecase(ctx context.Context, email, password string) error {
	if err := u.userRepository.CreateUser(ctx, email, password); err != nil {
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
