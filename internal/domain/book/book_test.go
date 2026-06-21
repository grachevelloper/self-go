package book_test

import (
	"book-service/internal/domain/book"
	"testing"
	"time"
)

func TestNewBook(t *testing.T) {
	publishedAt := time.Date(1949, time.June, 8, 0, 0, 0, 0, time.UTC)
	createdAt := time.Date(2026, time.June, 21, 12, 0, 0, 0, time.UTC)

	got, err := book.NewBook(book.NewBookParams{
		ID:          "book-id",
		Title:       "1984",
		Author:      "George Orwell",
		Status:      book.BookStatusDraft,
		PublishedAt: publishedAt,
		CreatedAt:   createdAt,
	})
	if err != nil {
		t.Fatalf("NewBook() error = %v", err)
	}

	if got.ID() != "book-id" {
		t.Errorf("ID() = %q, want %q", got.ID(), "book-id")
	}
	if got.Title() != "1984" {
		t.Errorf("Title() = %q, want %q", got.Title(), "1984")
	}
	if got.Author() != "George Orwell" {
		t.Errorf("Author() = %q, want %q", got.Author(), "George Orwell")
	}
	if got.Status() != book.BookStatusDraft {
		t.Errorf("Status() = %q, want %q", got.Status(), book.BookStatusDraft)
	}
	if !got.PublishedAt().Equal(publishedAt) {
		t.Errorf("PublishedAt() = %v, want %v", got.PublishedAt(), publishedAt)
	}
	if !got.CreatedAt().Equal(createdAt) {
		t.Errorf("CreatedAt() = %v, want %v", got.CreatedAt(), createdAt)
	}
	if got.UpdatedAt() != nil {
		t.Errorf("UpdatedAt() = %v, want nil", got.UpdatedAt())
	}
}

func TestNewBookRejectsInvalidData(t *testing.T) {
	valid := book.NewBookParams{
		ID:          "book-id",
		Title:       "1984",
		Author:      "George Orwell",
		Status:      book.BookStatusDraft,
		PublishedAt: time.Date(1949, time.June, 8, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
	}

	tests := []struct {
		name   string
		mutate func(*book.NewBookParams)
	}{
		{name: "empty id", mutate: func(p *book.NewBookParams) { p.ID = " " }},
		{name: "empty title", mutate: func(p *book.NewBookParams) { p.Title = " " }},
		{name: "empty author", mutate: func(p *book.NewBookParams) { p.Author = " " }},
		{name: "invalid status", mutate: func(p *book.NewBookParams) { p.Status = "unknown" }},
		{name: "future publication date", mutate: func(p *book.NewBookParams) { p.PublishedAt = time.Now().AddDate(1, 0, 0) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := valid
			tt.mutate(&params)

			if _, err := book.NewBook(params); err == nil {
				t.Fatal("NewBook() error = nil, want validation error")
			}
		})
	}
}

func TestBookUpdated(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	entity, err := book.NewBook(book.NewBookParams{
		ID:          "book-id",
		Title:       "Old title",
		Author:      "Author",
		Status:      book.BookStatusDraft,
		PublishedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   createdAt,
	})
	if err != nil {
		t.Fatalf("NewBook() error = %v", err)
	}

	newTitle := "New title"
	updatedAt := time.Date(2026, time.June, 21, 12, 0, 0, 0, time.UTC)
	updated, err := entity.Updated(book.UpdateBookParams{
		Title:     &newTitle,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		t.Fatalf("Updated() error = %v", err)
	}

	if entity.Title() != "Old title" {
		t.Errorf("original Title() = %q, want %q", entity.Title(), "Old title")
	}
	if updated.Title() != newTitle {
		t.Errorf("updated Title() = %q, want %q", updated.Title(), newTitle)
	}
	if updated.Author() != entity.Author() || updated.Status() != entity.Status() || !updated.PublishedAt().Equal(entity.PublishedAt()) {
		t.Error("Updated() changed fields absent from input")
	}
	if got := updated.UpdatedAt(); got == nil || !got.Equal(updatedAt) {
		t.Errorf("UpdatedAt() = %v, want %v", got, updatedAt)
	}
}
