package shared

import (
	"book-service/internal/domain/shared"
	"book-service/internal/usecase/shared/order"
	"net/http"
	"strconv"
)

func ParseIntQuery(r *http.Request, name string) (int, error) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return 0, nil
	}

	return strconv.Atoi(value)
}

func ParseSortOrder(value string) (order.New, error) {
	if value == "" {
		return order.Desc, nil
	}
	sortOrder := order.New(value)

	switch sortOrder {
	case order.Desc, order.Asc:
		return sortOrder, nil
	default:
		return "", &shared.ValidationError{
			Field: "order",
			Code:  "invalid",
		}
	}
}
