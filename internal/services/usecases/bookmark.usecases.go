package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type BookmarkUseCase interface {
	CreateBookmarkUsecase(ctx context.Context, userID, classID int) error
	DeleteBookmarkUsecase(ctx context.Context, userID, classID int) error
	GetBookmarksByUserIDUsecase(ctx context.Context, userID int) (*[]types.ClassResponse, error)
}

type bookmarkUseCase struct {
	bookmarkRepository repository.BookmarkRepository
}

func NewBookmarkUseCase(bookmarkRepository repository.BookmarkRepository) BookmarkUseCase {
	return &bookmarkUseCase{
		bookmarkRepository: bookmarkRepository,
	}
}

func (b *bookmarkUseCase) CreateBookmarkUsecase(ctx context.Context, userID, classID int) error {
	if err := b.bookmarkRepository.CreateBookmark(ctx, userID, classID); err != nil {
		return err
	}
	return nil
}

func (b *bookmarkUseCase) DeleteBookmarkUsecase(ctx context.Context, userID, classID int) error {
	if err := b.bookmarkRepository.DeleteBookmark(ctx, userID, classID); err != nil {
		return err
	}
	return nil
}

func (b *bookmarkUseCase) GetBookmarksByUserIDUsecase(ctx context.Context, userID int) (*[]types.ClassResponse, error) {
	bookmarks, err := b.bookmarkRepository.GetBookmarksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}
