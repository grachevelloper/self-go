package book

import (
	"book-service/internal/domain/book"
	"context"
)

type Repository interface {
	Create(ctx context.Context, input CreateBookInput) (*book.Book, error)
	GetById(ctx context.Context, id int64) (*book.Book, error)
	GetAll(ctx context.Context) ([]*book.Book, error)
	Update(ctx context.Context, id int64, input UpdateBookInput) (*book.Book, error)
	Delete(ctx context.Context, id int64) error
}
