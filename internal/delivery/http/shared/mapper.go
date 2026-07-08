package shared

import "book-service/internal/usecase/shared/paginated"

func PaginatedResponseMapper[T any, R any](paginated *paginated.New[T], items []R) PaginatedResponse[R] {
	return PaginatedResponse[R]{
		Page:    paginated.Page,
		HasNext: paginated.HasNext,
		Limit:   paginated.Limit,
		Total:   paginated.Total,
		Items:   items,
	}
}
