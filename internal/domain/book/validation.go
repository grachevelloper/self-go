package book

import (
	"book-service/internal/domain/shared"
	"strings"
)

func ValidateAuthor(author string) error {
	if author == "" {
		return &shared.ValidationError{
			Field: "author",
			Code:  "required",
		}
	}
	runes := []rune(author)
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
	if title == "" {
		return &shared.ValidationError{
			Field: "title",
			Code:  "required",
		}
	}
	return nil
}
