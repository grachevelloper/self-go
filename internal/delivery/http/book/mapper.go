package book

import (
	"book-service/internal/domain/book"
)

func bookResponseMapper(book *book.Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
		Status:      string(book.Status),
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   *book.UpdatedAt,
	}
}
