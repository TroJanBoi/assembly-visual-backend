package repository

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type CatRepository interface {
	GetAllCats(ctx context.Context) (*[]types.CatResponse, error)
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{
		db: db,
	}
}

func (c *catRepository) GetAllCats(ctx context.Context) (*[]types.CatResponse, error) {
	return nil, nil
}
