package book

import (
	"book-service/internal/domain/book"
	"context"
	"time"
)

//go:generate go tool mockgen -source=repository.go -destination=mocks/repository.go -package=mocks

type Repository interface {
	Create(ctx context.Context, entity *book.Book) (*book.Book, error)
	GetById(ctx context.Context, id string) (*book.Book, error)
	GetAll(ctx context.Context) ([]*book.Book, error)
	Update(ctx context.Context, entity *book.Book, expectedUpdatedAt *time.Time) (*book.Book, error)
	Delete(ctx context.Context, id string) error
}
