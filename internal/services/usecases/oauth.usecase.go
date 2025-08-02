package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
)

type OAuthUseCase interface {
	HandleOAuthUseCase(ctx context.Context, code string) (string, error)
}

type oauthUseCase struct {
	repo repository.OAuthRepository
}

func NewOAuthUseCase(repo repository.OAuthRepository) OAuthUseCase {
	return &oauthUseCase{
		repo: repo,
	}
}

func (o *oauthUseCase) HandleOAuthUseCase(ctx context.Context, code string) (string, error) {
	return o.repo.HandleOAuth(ctx, code)
}
