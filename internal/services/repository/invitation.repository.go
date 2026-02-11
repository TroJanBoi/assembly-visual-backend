package repository

import (
	"context"
	"errors"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type InvitationRepository interface {
	SendEmailInvitation(ctx context.Context, email string, owner int, classID int) error
	GetAllInvitationsByClassID(ctx context.Context, classID int) (*[]types.InvitationResponse, error)
	GetInvitationMe(ctx context.Context, userID int) (*[]types.InvitationResponse, error)
	UpdateInvitationStatus(ctx context.Context, invitationID int, userID int, status string) error
}

type invitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) InvitationRepository {
	return &invitationRepository{db: db}
}

func (r *invitationRepository) SendEmailInvitation(ctx context.Context, email string, owner int, classID int) error {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("user with the provided email does not exist")
	}

	var class model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&class).Error; err != nil {
		return errors.New("class not found or you are not the owner")
	}

	var member model.Member
	if err := r.db.WithContext(ctx).Where("class_id = ? AND user_id = ?", classID, user.ID).First(&member).Error; err == nil {
		return errors.New("user is already a member of the class")
	}

	if class.OwnerId == int(user.ID) {
		return errors.New("cannot invite the class owner")
	}

	invitation := model.Invitation{
		InvitedEmail:  email,
		ClassID:       int(classID),
		InvitedUserID: int(user.ID),
		Status:        "pending",
	}

	if err := r.db.WithContext(ctx).Create(&invitation).Error; err != nil {
		return err
	}

	return nil
}

func (r *invitationRepository) GetAllInvitationsByClassID(ctx context.Context, classID int) (*[]types.InvitationResponse, error) {
	var invitations []model.Invitation
	if err := r.db.WithContext(ctx).Where("class_id = ?", classID).Find(&invitations).Error; err != nil {
		return nil, err
	}

	var invitationResponses []types.InvitationResponse
	for _, invitation := range invitations {
		invitationResponses = append(invitationResponses, types.InvitationResponse{
			ID:              int(invitation.ID),
			ClassID:         invitation.ClassID,
			UserID:          invitation.InvitedUserID,
			InvitationEmail: invitation.InvitedEmail,
			Status:          invitation.Status,
			Token:           invitation.Token,
			Expired:         invitation.Expired,
		})
	}

	return &invitationResponses, nil
}

func (r *invitationRepository) GetInvitationMe(ctx context.Context, userID int) (*[]types.InvitationResponse, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var invitations []model.Invitation
	if err := r.db.WithContext(ctx).Where("invited_user_id = ?", userID).Find(&invitations).Error; err != nil {
		return nil, err
	}

	var invitationResponses []types.InvitationResponse
	for _, invitation := range invitations {
		invitationResponses = append(invitationResponses, types.InvitationResponse{
			ID:              int(invitation.ID),
			ClassID:         invitation.ClassID,
			UserID:          invitation.InvitedUserID,
			InvitationEmail: invitation.InvitedEmail,
			Status:          invitation.Status,
			Token:           invitation.Token,
			Expired:         invitation.Expired,
		})
	}

	return &invitationResponses, nil
}

func (r *invitationRepository) UpdateInvitationStatus(ctx context.Context, invitationID int, userID int, status string) error {

	var invitation model.Invitation
	if err := r.db.WithContext(ctx).Where("id = ? AND invited_user_id = ?", invitationID, userID).First(&invitation).Error; err != nil {
		return errors.New("invitation not found for this user")
	}

	expireAt := invitation.CreatedAt.Add(72 * time.Hour)
	if time.Now().After(expireAt) {
		invitation.Status = "expired"

		if err := r.db.WithContext(ctx).Save(&invitation).Error; err != nil {
			return err
		}
		if err := r.db.WithContext(ctx).Delete(&invitation).Error; err != nil {
			return err
		}
		return errors.New("invitation has expired")
	}

	if invitation.InvitedUserID != userID {
		return errors.New("you are not recipient of this invitation")
	}

	invitation.Status = status

	if err := r.db.WithContext(ctx).Save(&invitation).Error; err != nil {
		return err
	}

	if status == "accepted" {
		newMember := model.Member{
			UserID:  userID,
			ClassID: invitation.ClassID,
		}
		if err := r.db.WithContext(ctx).Create(&newMember).Error; err != nil {
			return err
		}
	}

	return nil
}
