package shared

import "book-service/internal/domain/shared/paginated"

func PaginatedResponseMapper[T any, R any](paginated *paginated.PaginatedEntity[T], items []R) PaginatedResponse[R] {
	return PaginatedResponse[R]{
		Page:    paginated.Page,
		HasNext: paginated.HasNext,
		Limit:   paginated.Limit,
		Total:   paginated.Total,
		Items:   items,
	}
}
