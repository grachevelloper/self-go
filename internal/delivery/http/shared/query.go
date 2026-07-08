package shared

import (
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
