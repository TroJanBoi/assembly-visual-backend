package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
)

type SystemsUseCase interface {
	NOP(ctx context.Context) error
	HLT(ctx context.Context) (bool, error)
	LABELS(ctx context.Context, label string) (map[string]string, error)
}

type systemsUseCase struct {
	systemsRepo repository.SystemsRepository
}

func NewSystemsUseCase(systemsRepo repository.SystemsRepository) SystemsUseCase {
	return &systemsUseCase{systemsRepo: systemsRepo}
}

func (uc *systemsUseCase) NOP(ctx context.Context) error {
	err := uc.systemsRepo.NOP(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (uc *systemsUseCase) HLT(ctx context.Context) (bool, error) {
	halted, err := uc.systemsRepo.HLT(ctx)
	if err != nil {
		return false, err
	}
	return halted, nil
}

func (uc *systemsUseCase) LABELS(ctx context.Context, label string) (map[string]string, error) {
	labels, err := uc.systemsRepo.LABELS(ctx, label)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
