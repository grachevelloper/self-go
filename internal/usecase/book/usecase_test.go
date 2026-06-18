package book_test

import (
	"book-service/internal/usecase/book"
	"time"
)

var mockedBookInput = book.CreateBookInput{
	Title:       "1984",
	Author:      "Оруэл",
	PublishedAt: time.Date(1947, time.June, 18, 0, 0, 0, 0, time.UTC),
	Status:      "reading",
}

// func TestCreate(t *testing.T) {
// 	got := CreateBook()
// }
