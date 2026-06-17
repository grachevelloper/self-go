package book

import "time"

type CreateBookRequest struct {
	Title       string    `json:"title"`
	Author      string    `json:"authors"`
	Status      string    `json:"status"`
	PublishedAt time.Time `json:"published_at"`
}

type UpdateBookRequest struct {
	Title       *string    `json:"title"`
	Author      *string    `json:"authors"`
	Status      *string    `json:"status"`
	PublishedAt *time.Time `json:"published_at"`
}
