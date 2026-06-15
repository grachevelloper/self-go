package handler

import "context"

type BookService interface {
	CreateBook(ctx context.Context, input CreateBookInput) (CreateBookOutput, error)
}
