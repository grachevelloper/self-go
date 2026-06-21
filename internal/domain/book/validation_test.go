package book_test

import (
	"book-service/internal/domain/book"
	"book-service/internal/domain/shared"
	"errors"
	"testing"
	"time"
)

func TestValidateTitle(t *testing.T) {
	var validationErr *shared.ValidationError

	t.Run("valid title", func(t *testing.T) {
		got := book.ValidateTitle("1984")
		if got != nil {
			t.Fatalf("book.ValidateTitle(%q) error = %v; want nil", "1984", got)
		}
	})

	t.Run("empty title", func(t *testing.T) {
		got := book.ValidateTitle("")

		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateTitle(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})

	t.Run("only space title", func(t *testing.T) {
		got := book.ValidateTitle("    ")
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

	t.Run("only space author", func(t *testing.T) {
		got := book.ValidateAuthor("    ")

		var validationErr *shared.ValidationError
		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateAuthor(%q) error = %v; want ValidationError",
				"",
				got,
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

func TestValidatePublishedAt(t *testing.T) {
	var validationErr *shared.ValidationError

	t.Run("valid publishedAt", func(t *testing.T) {
		nextYear := time.Date(time.Now().Year()-100, time.January, 1, 0, 0, 0, 0, time.UTC)

		got := book.ValidatePublishedAt(nextYear)

		if got != nil {
			t.Fatalf("book.ValidatePublishedAt(%q) error = %v; want nil", "time.Date(time.Now().Year()-100, time.January, 1, 0, 0, 0, 0, time.UTC)", got)
		}
	})
	t.Run("only current and pasted years", func(t *testing.T) {
		nextYear := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)

		got := book.ValidatePublishedAt(nextYear)

		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidatePublishedAt(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}

func TestValidateStatus(t *testing.T) {
	var validationErr *shared.ValidationError
	t.Run("valid status", func(t *testing.T) {
		got := book.ValidateStatus("reading")

		if got != nil {
			t.Fatalf("book.ValidateStatus(%q) error = %v; want nil", "reading", got)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		got := book.ValidateStatus("Walter White")

		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateStatus(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}

func TestValidateID(t *testing.T) {
	var validationErr *shared.ValidationError

	t.Run("valid title", func(t *testing.T) {
		got := book.ValidateID("y23789523845923094")
		if got != nil {
			t.Fatalf("book.ValidateID(%q) error = %v; want nil", "1984", got)
		}
	})

	t.Run("only space title", func(t *testing.T) {
		got := book.ValidateID("    ")
		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateID(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}
