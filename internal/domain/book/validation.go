package book

import (
	"book-service/internal/domain/shared"
	"strings"
	"time"
)

func ValidateAuthor(author string) error {
	trimmedAuthor := strings.TrimSpace(author)

	if trimmedAuthor == "" {
		return &shared.ValidationError{
			Field: "author",
			Code:  "required",
		}
	}
	runes := []rune(trimmedAuthor)
	firstLetter := string(runes[0])

	if firstLetter != strings.ToUpper(firstLetter) {
		return &shared.ValidationError{
			Field: "author",
			Code:  "must_start_with_uppercase",
		}
	}
	return nil
}

func ValidateTitle(title string) error {
	trimmedTitle := strings.TrimSpace(title)

	if trimmedTitle == "" {
		return &shared.ValidationError{
			Field: "title",
			Code:  "required",
		}
	}
	return nil
}

func ValidatePublishedAt(publishedAt time.Time) error {
	if publishedAt.After(time.Now()) {
		return &shared.ValidationError{
			Field: "publishedAt",
			Code:  "must_be_current_time_or_less",
		}
	}
	return nil
}

func ValidateStatus(status string) error {
	switch status {
	case "reading", "in_wishlist", "finished":
		return nil
	default:
		return &shared.ValidationError{
			Field: "status",
			Code:  "status_does_not_match_book_status_enum",
		}
	}

}

func ValidateID(id string) error {
	trimmedID := strings.TrimSpace(id)

	if trimmedID == "" {
		return &shared.ValidationError{
			Field: "id",
			Code:  "required",
		}
	}
	return nil
}
