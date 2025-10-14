package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type ProfileUseCase interface {
	GetProfileUsecases(ctx context.Context, userID int) (*types.UserResponse, error)
	EditProfileUsecases(ctx context.Context, userID int, user *types.EditProfileRequest) error
	ChangePasswordUsecases(ctx context.Context, userID int, newPassword string) error
}

type profileUseCase struct {
	profileRepository repository.ProfileRepository
}

func NewProfileUseCase(profileRepository repository.ProfileRepository) ProfileUseCase {
	return &profileUseCase{
		profileRepository: profileRepository,
	}
}

func (p *profileUseCase) GetProfileUsecases(ctx context.Context, userID int) (*types.UserResponse, error) {
	user, err := p.profileRepository.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *profileUseCase) EditProfileUsecases(ctx context.Context, userID int, user *types.EditProfileRequest) error {
	if err := p.profileRepository.EditProfile(ctx, userID, user); err != nil {
		return err
	}
	return nil
}

func (p *profileUseCase) ChangePasswordUsecases(ctx context.Context, userID int, newPassword string) error {
	if err := p.profileRepository.ChangePassword(ctx, userID, newPassword); err != nil {
		return err
	}
	return nil
}
