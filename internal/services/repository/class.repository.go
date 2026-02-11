package repository

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type ClassRepository interface {
	GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error)
	GetClassByID(ctx context.Context, classID int) (*types.ClassResponse, error)
	CreateClass(ctx context.Context, owner int, class *types.CreateClassRequest) error
	UpdateClass(ctx context.Context, owner int, classID int, class *types.UpdateClassRequest) error
	DeleteClass(ctx context.Context, owner int, classID int) error
	JoinClass(ctx context.Context, userID, classID int) error
	GetAllMembersByClassID(ctx context.Context, classID int) (*[]types.MemberResponse, error)
	GetAllClassPublic(ctx context.Context) (*[]types.ClassResponse, error)
	ChangePermissionMember(ctx context.Context, userID, classID int, newRole string) error
	RemoveMemberInClass(ctx context.Context, classID, userID int) error
	GetClassRecentManyIDs(ctx context.Context, limit []int) (*[]types.ClassResponse, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func generateClassCode() (string, error) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLength := 8

	b := make([]byte, codeLength)
	max := big.NewInt(int64(len(letters)))

	for i := 0; i < codeLength; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = letters[n.Int64()]
	}
	return string(b), nil
}

func (r *classRepository) GetAllClasses(ctx context.Context) (*[]types.ClassResponse, error) {
	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Find(&classes).Error; err != nil {
		return nil, err
	}

	classResponses := make([]types.ClassResponse, 0, len(classes))

	for _, class := range classes {
		var owner model.User
		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&owner).Error; err != nil {
			return nil, err
		}

		var memberCount int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&memberCount).Error; err != nil {
			return nil, err
		}
		classResponses = append(classResponses, types.ClassResponse{
			ID:               int(class.ID),
			Topic:            class.Topic,
			Description:      class.Description,
			Code:             class.Code,
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			OwnerID:          int(class.OwnerId),
			OwnerName:        owner.Name,
			Status:           class.Status,
			MemberAmount:     memberCount,
			Favorite:         class.Favorite,
			BannerID:         class.BannerID,
		})
	}

	return &classResponses, nil
}

func (r *classRepository) CreateClass(ctx context.Context, owner int, class *types.CreateClassRequest) error {
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	code, err := generateClassCode()
	if err != nil {
		return err
	}

	newClass := model.Classroom{
		Topic:            class.Topic,
		Description:      class.Description,
		GoogleCourseID:   class.GoogleCourseID,
		GoogleCourseLink: class.GoogleCourseLink,
		OwnerId:          owner,
		BannerID:         class.BannerID,
		Status:           class.Status,
		Code:             code,
	}

	if err := r.db.WithContext(ctx).Create(&newClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) UpdateClass(ctx context.Context, owner int, classID int, class *types.UpdateClassRequest) error {
	var existingClass model.Classroom
	err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&existingClass).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if class.Topic != "" {
		existingClass.Topic = class.Topic
	}
	if class.Description != "" {
		existingClass.Description = class.Description
	}
	if class.GoogleCourseID != "" {
		existingClass.GoogleCourseID = class.GoogleCourseID
	}
	if class.GoogleCourseLink != "" {
		existingClass.GoogleCourseLink = class.GoogleCourseLink
	}
	if class.Status != 0 {
		existingClass.Status = class.Status
	}
	if err := r.db.WithContext(ctx).Save(&existingClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) DeleteClass(ctx context.Context, owner int, classID int) error {
	var existingClass model.Classroom
	err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, owner).First(&existingClass).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&existingClass).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) GetClassByID(ctx context.Context, classID int) (*types.ClassResponse, error) {
	var class model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var owner model.User
	if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&owner).Error; err != nil {
		return nil, err
	}

	var memberCount int64
	if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&memberCount).Error; err != nil {
		return nil, err
	}

	classResponse := types.ClassResponse{
		ID:               int(class.ID),
		Topic:            class.Topic,
		Description:      class.Description,
		Code:             class.Code,
		GoogleCourseID:   class.GoogleCourseID,
		GoogleCourseLink: class.GoogleCourseLink,
		GoogleSyncedAt:   class.GoogleSyncedAt,
		OwnerID:          int(class.OwnerId),
		OwnerName:        owner.Name,
		MemberAmount:     memberCount,
		Status:           class.Status,
		Favorite:         class.Favorite,
		BannerID:         class.BannerID,
	}

	return &classResponse, nil
}

func (r *classRepository) JoinClass(ctx context.Context, userID, classID int) error {
	// Check if the user exists
	var users model.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Check if the class exists
	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&classes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Check if the user is the owner of the class
	if err := r.db.WithContext(ctx).Where("id = ? AND owner_id = ?", classID, userID).First(&classes).Error; err == nil {
		return errors.New("owner cannot join their own class as member")
	}

	// Check if the user is already a member of the class
	var member model.Member
	if err := r.db.WithContext(ctx).Where("user_id = ? AND class_id = ?", userID, classID).First(&member).Error; err == nil {
		return errors.New("user already joined the class")
	}

	newMember := model.Member{
		UserID:  userID,
		ClassID: classID,
		JoinAt:  time.Now(),
		Role:    "member",
	}

	if err := r.db.WithContext(ctx).Create(&newMember).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) GetAllMembersByClassID(ctx context.Context, classID int) (*[]types.MemberResponse, error) {
	var members []model.Member
	if err := r.db.WithContext(ctx).Where("class_id = ?", classID).Find(&members).Error; err != nil {
		return nil, err
	}

	var memberResponses []types.MemberResponse
	for _, member := range members {
		var user model.User
		if err := r.db.WithContext(ctx).Where("id = ?", member.UserID).First(&user).Error; err != nil {
			return nil, err
		}

		memberResponses = append(memberResponses, types.MemberResponse{
			ID:          int(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			PicturePath: user.PicturePath,
			Role:        member.Role,
			JoinAt:      member.JoinAt,
		})
	}

	return &memberResponses, nil
}

func (r *classRepository) GetAllClassPublic(ctx context.Context) (*[]types.ClassResponse, error) {
	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("status = ?", 0).Find(&classes).Error; err != nil {
		return nil, err
	}

	classResponses := make([]types.ClassResponse, 0, len(classes))

	for _, class := range classes {
		var owner model.User
		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&owner).Error; err != nil {
			return nil, err
		}

		var memberCount int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&memberCount).Error; err != nil {
			return nil, err
		}
		classResponses = append(classResponses, types.ClassResponse{
			ID:               int(class.ID),
			Topic:            class.Topic,
			Description:      class.Description,
			Code:             class.Code,
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			OwnerID:          int(class.OwnerId),
			OwnerName:        owner.Name,
			MemberAmount:     memberCount,
			Status:           class.Status,
			Favorite:         class.Favorite,
			BannerID:         class.BannerID,
		})
	}

	return &classResponses, nil
}

func (r *classRepository) ChangePermissionMember(ctx context.Context, userID, classID int, newRole string) error {
	var member model.Member
	err := r.db.WithContext(ctx).Where("user_id = ? AND class_id = ?", userID, classID).First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	member.Role = newRole

	if err := r.db.WithContext(ctx).Save(&member).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) RemoveMemberInClass(ctx context.Context, classID, userID int) error {
	var member model.Member
	err := r.db.WithContext(ctx).Where("user_id = ? AND class_id = ?", userID, classID).First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&member).Error; err != nil {
		return err
	}

	return nil
}

func (r *classRepository) GetClassRecentManyIDs(ctx context.Context, limit []int) (*[]types.ClassResponse, error) {
	var classes []model.Classroom
	if err := r.db.WithContext(ctx).Where("id IN ?", limit).Find(&classes).Error; err != nil {
		return nil, err
	}

	classResponses := make([]types.ClassResponse, 0, len(classes))

	for _, class := range classes {
		var owner model.User
		if err := r.db.WithContext(ctx).Where("id = ?", class.OwnerId).First(&owner).Error; err != nil {
			return nil, err
		}

		var memberCount int64
		if err := r.db.WithContext(ctx).Model(&model.Member{}).Where("class_id = ?", class.ID).Count(&memberCount).Error; err != nil {
			return nil, err
		}
		classResponses = append(classResponses, types.ClassResponse{
			ID:               int(class.ID),
			Topic:            class.Topic,
			Description:      class.Description,
			Code:             class.Code,
			GoogleCourseID:   class.GoogleCourseID,
			GoogleCourseLink: class.GoogleCourseLink,
			GoogleSyncedAt:   class.GoogleSyncedAt,
			OwnerID:          int(class.OwnerId),
			OwnerName:        owner.Name,
			MemberAmount:     memberCount,
			Status:           class.Status,
			Favorite:         class.Favorite,
			BannerID:         class.BannerID,
		})
	}

	return &classResponses, nil
}
