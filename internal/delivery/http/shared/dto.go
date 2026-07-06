package shared

type PaginatedResponse[T any] struct {
	Items   []T  `json:"items"`
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	HasNext bool `json:"has_next"`
}
