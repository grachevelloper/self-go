package paginated_test

import (
	"book-service/internal/domain/shared"
	"book-service/internal/domain/shared/paginated"
	"errors"
	"testing"
)

func TestValidateLimit(t *testing.T) {
	var validationErr *shared.ValidationError

	t.Run("valid limit", func(t *testing.T) {
		got := paginated.ValidateLimit(67)
		if got != nil {
			t.Fatalf("book.ValidateLimit(%q) error = %v; want nil", "67", got)
		}
	})

	t.Run("negative limit", func(t *testing.T) {
		got := paginated.ValidateLimit(-3)

		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidateLimit(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}

func TestValidatePage(t *testing.T) {
	var validationErr *shared.ValidationError

	t.Run("valid page", func(t *testing.T) {
		got := paginated.ValidatePage(67)
		if got != nil {
			t.Fatalf("book.ValidatePage(%q) error = %v; want nil", "67", got)
		}
	})

	t.Run("negative page", func(t *testing.T) {
		got := paginated.ValidatePage(-3)

		if !errors.As(got, &validationErr) {
			t.Fatalf(
				"book.ValidatePage(%q) error = %v; want ValidationError",
				"",
				got,
			)
		}
	})
}
