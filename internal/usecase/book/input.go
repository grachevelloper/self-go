package book

import "time"

type CreateBookInput struct {
	Title       string
	Author      string
	Status      string
	PublishedAt time.Time
}
type UpdateBookInput struct {
	Title       *string
	Author      *string
	Status      *string
	PublishedAt *time.Time
}
