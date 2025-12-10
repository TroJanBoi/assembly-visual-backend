package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type AuthUseCase interface {
	// Define authentication-related use case methods here
	RegisterUserUsecase(ctx context.Context, newUser *types.SignInRequest) error
	LoginUserUsecase(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		authRepository: authRepository,
	}
}

func (a *authUseCase) RegisterUserUsecase(ctx context.Context, newUser *types.SignInRequest) error {
	if err := a.authRepository.RegisterUser(ctx, newUser); err != nil {
		return err
	}
	return nil
}

func (a *authUseCase) LoginUserUsecase(ctx context.Context, user *types.LoginRequest) (*types.LoginResponse, error) {
	if user.Email == "" || user.Password == "" {
		return &types.LoginResponse{}, nil
	}
	tokenStr, err := a.authRepository.LoginUser(ctx, user)
	if err != nil {
		return &types.LoginResponse{}, err
	}
	return &types.LoginResponse{Token: tokenStr.Token}, nil
}
