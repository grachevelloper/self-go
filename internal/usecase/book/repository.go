package book

import (
	"book-service/internal/domain"
	"context"
)

type Repository interface {
	Create(ctx context.Context, book domain.Book) error
	GetById(ctx context.Context, id int64) (*domain.Book, error)
	GetAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, id int64) (*domain.Book, error)
	Delete(ctx context.Context, id int64) (*domain.Book, error)
}
