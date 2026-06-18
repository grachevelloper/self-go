package book

import (
	"book-service/internal/domain/book"
	"context"
)

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, input CreateBookInput) (*book.Book, error) {
	if err := book.ValidateAuthor(input.Author); err != nil {
		return nil, err
	}
	if err := book.ValidateTitle(input.Title); err != nil {
		return nil, err
	}

	return u.repo.Create(ctx, input)
}
func (u *UseCase) GetById(ctx context.Context, id int64) (*book.Book, error) {
	return u.repo.GetById(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) ([]*book.Book, error) {
	return u.repo.GetAll(ctx)
}

func (u *UseCase) Update(ctx context.Context, id int64, input UpdateBookInput) (*book.Book, error) {

	if input.Author != nil {
		if err := book.ValidateAuthor(*input.Author); err != nil {
			return nil, err
		}
	}

	if input.Title != nil {
		if err := book.ValidateAuthor(*input.Title); err != nil {
			return nil, err
		}
	}

	return u.repo.Update(ctx, id, input)
}

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
