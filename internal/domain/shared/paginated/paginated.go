package paginated

type PaginatedEntity[T any] struct {
	Page    int
	HasNext bool
	Items   []T
	Limit   int
	Total   int
}

type PaginationParams struct {
	Page  int
	Limit int
}

func ValidatePaginatedEntity(params PaginationParams) error {
	if err := ValidatePage(params.Page); err != nil {
		return err
	}
	return ValidateLimit(params.Limit)
}
