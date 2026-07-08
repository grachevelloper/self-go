package paginated

import (
	"book-service/internal/usecase/shared/order"
)

type New[T any] struct {
	Page    int
	HasNext bool
	Items   []T
	Limit   int
	Total   int
}

type PaginationParams[T any] struct {
	Page      int
	Limit     int
	Order     order.New
	SortField T
}

func ValidatePaginatedEntity[T any](params PaginationParams[T]) error {
	if err := ValidatePage(params.Page); err != nil {
		return err
	}
	return ValidateLimit(params.Limit)
}
