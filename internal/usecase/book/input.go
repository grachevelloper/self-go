package book

import (
	"book-service/internal/usecase/shared/order"
	"time"
)

type CreateBookInput struct {
	Title       string
	Author      string
	Status      string
	PublishedAt time.Time
}
type UpdateBookInput struct {
	ID          string
	Title       *string
	Author      *string
	Status      *string
	PublishedAt *time.Time
}

type GetAllBooksInput struct {
	Page      int
	Limit     int
	SortField BookSortField
	Order     order.New
}

type BookSortField string

const (
	Title     BookSortField = "title"
	CreatedAt BookSortField = "created_at"
)
