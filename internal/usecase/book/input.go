package book

import "time"

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
	Page  int
	Limit int
}
