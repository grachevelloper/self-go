package repository

import (
	usecasebook "book-service/internal/usecase/book"
	"book-service/internal/usecase/shared/order"
	"testing"
)

func TestBookOrderingSQL(t *testing.T) {
	t.Run("maps allowed sort options to SQL fragments", func(t *testing.T) {
		sortField, sortOrder, err := bookOrderingSQL(usecasebook.Title, order.Asc)
		if err != nil {
			t.Fatalf("bookOrderingSQL() error = %v", err)
		}

		if sortField != "title" || sortOrder != "ASC" {
			t.Fatalf("bookOrderingSQL() = %q, %q; want title, ASC", sortField, sortOrder)
		}
	})

	t.Run("rejects unknown sort field", func(t *testing.T) {
		_, _, err := bookOrderingSQL(usecasebook.BookSortField("67"), order.Asc)
		if err == nil {
			t.Fatal("bookOrderingSQL() error = nil, want error")
		}
	})

	t.Run("rejects unknown order", func(t *testing.T) {
		_, _, err := bookOrderingSQL(usecasebook.Title, order.New("67"))
		if err == nil {
			t.Fatal("bookOrderingSQL() error = nil, want error")
		}
	})
}
