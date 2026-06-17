package book

import (
	"book-service/internal/domain"
	"context"
)

type Handler struct {
	serivce Service
}

type Service interface {
	Create(ctx context.Context, input CreateBookRequest) error
	GetById(ctx context.Context, id int64) (*domain.Book, error)
	GetAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, id int64, input UpdateBookRequest) (*domain.Book, error)
	Delete(ctx context.Context, id int64) error
}
