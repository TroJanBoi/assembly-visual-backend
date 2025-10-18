package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type InvitationUseCase interface {
	SendEmailInvitationUseCases(ctx context.Context, email string, owner int, classID int) error
	GetAllInvitationsByClassIDUseCases(ctx context.Context, classID int) (*[]types.InvitationResponse, error)
	GetInvitationMeUseCases(ctx context.Context, userID int) (*[]types.InvitationResponse, error)
	UpdateInvitationStatusUseCases(ctx context.Context, invitationID int, userID int, status string) error
}

type invitationUseCase struct {
	invitationRepo repository.InvitationRepository
}

func NewInvitationUseCase(invitationRepo repository.InvitationRepository) InvitationUseCase {
	return &invitationUseCase{invitationRepo: invitationRepo}
}

func (uc *invitationUseCase) SendEmailInvitationUseCases(ctx context.Context, email string, owner int, classID int) error {
	err := uc.invitationRepo.SendEmailInvitation(ctx, email, owner, classID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *invitationUseCase) GetAllInvitationsByClassIDUseCases(ctx context.Context, classID int) (*[]types.InvitationResponse, error) {
	resp, err := uc.invitationRepo.GetAllInvitationsByClassID(ctx, classID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *invitationUseCase) GetInvitationMeUseCases(ctx context.Context, userID int) (*[]types.InvitationResponse, error) {
	resp, err := uc.invitationRepo.GetInvitationMe(ctx, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *invitationUseCase) UpdateInvitationStatusUseCases(ctx context.Context, invitationID int, userID int, status string) error {
	err := uc.invitationRepo.UpdateInvitationStatus(ctx, invitationID, userID, status)
	if err != nil {
		return err
	}
	return nil
}
