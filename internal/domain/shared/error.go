package shared

type ValidationError struct {
	Field string
	Code  string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Code
}
