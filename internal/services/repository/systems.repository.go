package repository

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type SystemsRepository interface {
	NOP(ctx context.Context) error
	HLT(ctx context.Context) (bool, error)
	LABELS(ctx context.Context, label string) (map[string]string, error)
}

type systemsRepository struct {
	db *gorm.DB
}

func NewSystemsRepository(db *gorm.DB) SystemsRepository {
	return &systemsRepository{db: db}
}

func (r *systemsRepository) NOP(ctx context.Context) error {
	log.Println("System NOP Invoked")
	time.Sleep(1 * time.Second)
	return nil
}

func (r *systemsRepository) HLT(ctx context.Context) (bool, error) {
	log.Println("System Halted Check Invoked")
	return true, nil
}

func (r *systemsRepository) LABELS(ctx context.Context, label string) (map[string]string, error) {
	log.Println("System Labels Check Invoked")
	return map[string]string{"label": label}, nil
}
