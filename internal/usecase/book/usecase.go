package book

import (
	"book-service/internal/domain"
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

func (u *UseCase) Create(ctx context.Context, input CreateBookInput) (*domain.Book, error) {

	return u.repo.Create(ctx, input)
}
func (u *UseCase) GetById(ctx context.Context, id int64) (*domain.Book, error) {
	return u.repo.GetById(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) ([]*domain.Book, error) {
	return u.repo.GetAll(ctx)
}

func (u *UseCase) Update(ctx context.Context, id int64, input UpdateBookInput) (*domain.Book, error) {
	return u.repo.Update(ctx, id, input)
}

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
