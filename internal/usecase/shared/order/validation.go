package order

import "book-service/internal/domain/shared"

func ValidateOrder(status string) error {
	switch status {
	case "asc", "desc":
		return nil
	default:
		return &shared.ValidationError{
			Field: "order",
			Code:  "order_does_not_match_orders_enum",
		}
	}

}
