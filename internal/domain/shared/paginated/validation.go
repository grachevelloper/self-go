package paginated

import (
	"book-service/internal/domain/shared"
)

func ValidateLimit(limit int) error {
	if limit < 0 {
		return &shared.ValidationError{
			Field: "limit",
			Code:  "must be > 0",
		}
	}
	return nil
}

func ValidatePage(page int) error {
	if page < 0 {
		return &shared.ValidationError{
			Field: "page",
			Code:  "must be > 0",
		}
	}
	return nil
}
