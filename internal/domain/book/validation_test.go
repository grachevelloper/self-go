package book_test

import (
	"book-service/internal/domain/book"
	"book-service/internal/domain/shared"
	"errors"
	"testing"
)

func TestValidateTitle(t *testing.T) {
	t.Run("valid title", func(t *testing.T) {
		got := book.ValidateTitle("1984")
		if got != nil {
			t.Fatalf("book.ValidateTitle(%q) error = %v; want nil", "1984", got)
		}
	})

	t.Run("empty title", func(t *testing.T) {
		got := book.ValidateTitle("")

		var validationErr *shared.ValidationError
		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateTitle(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}

func TestValidateAuthor(t *testing.T) {
	t.Run("valid author", func(t *testing.T) {
		got := book.ValidateAuthor("Samuel Leroy Jackson")
		if got != nil {
			t.Fatalf("ValidateAuthor(%q) error = %v; want nil", "1984", got)
		}
	})

	t.Run("empty author", func(t *testing.T) {
		got := book.ValidateAuthor("")

		var validationErr *shared.ValidationError
		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"ValidateAuthor(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
		if validationErr.Field != "author" || validationErr.Code != "required" {
			t.Errorf(
				"book.ValidateTitle(%q) error = %+v; want Field=%q Code=%q",
				"",
				validationErr,
				"author",
				"required",
			)
		}
	})

	t.Run("lowercase author", func(t *testing.T) {
		got := book.ValidateAuthor("samuel Leroy Jackson")

		var validationErr *shared.ValidationError
		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"ValidateAthor(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}

		if validationErr.Field != "author" || validationErr.Code != "must_start_with_uppercase" {
			t.Errorf(
				"book.ValidateTitle(%q) error = %+v; want Field=%q Code=%q",
				"",
				validationErr,
				"author",
				"must_start_with_uppercase",
			)
		}
	})
}
