package repository

import (
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	// CreateSubmission(ctx context.Context, )
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}
