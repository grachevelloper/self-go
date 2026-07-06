package book

import (
	"book-service/internal/domain/book"
	"book-service/internal/domain/shared/paginated"
	"context"
	"time"
)

//go:generate go tool mockgen -source=repository.go -destination=mocks/repository.go -package=mocks

type Repository interface {
	Create(ctx context.Context, entity *book.Book) (*book.Book, error)
	GetById(ctx context.Context, id string) (*book.Book, error)
	GetAll(ctx context.Context, paginateParams paginated.PaginationParams) (*paginated.PaginatedEntity[book.Book], error)
	Update(ctx context.Context, id string, params book.UpdateBookParams, expectedUpdatedAt *time.Time) (*book.Book, error)
	Delete(ctx context.Context, id string) error
}
