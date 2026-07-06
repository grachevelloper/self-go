package book

import (
	"time"
)

type CreateBookRequest struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Status      string    `json:"status"`
	PublishedAt time.Time `json:"published_at"`
}

type BookResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Status      string     `json:"status"`
	PublishedAt time.Time  `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type UpdateBookRequest struct {
	Title       *string    `json:"title"`
	Author      *string    `json:"author"`
	Status      *string    `json:"status"`
	PublishedAt *time.Time `json:"published_at"`
}

type GetAllBooksRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
