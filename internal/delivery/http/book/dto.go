package book

import (
	"book-service/internal/domain"
)

type CreateBookInput struct {
	Title  string `json:title`
	Author string `json:authors`
	Status string `json:status`
}

type UpdateBookInput struct {
	Title  *string `json:title`
	Author *string `json:authors`
	Status *string `json:status`
}

type UpdateBookOutput struct {
	*domain.Book
}
