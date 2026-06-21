package delivery

import (
	"book-service/internal/config"
	"book-service/internal/delivery/middleware"
	"net/http"
)

type RouteRegistrar interface {
	RegisterRoutes(mux *http.ServeMux)
}

func NewRouter(handlers ...RouteRegistrar) http.Handler {
	mux := http.NewServeMux()

	mainMiddleware := middleware.Chain(
		middleware.CORS(config.FromEnv().Origin),
		middleware.SetHeaders(),
	)

	for _, h := range handlers {
		h.RegisterRoutes(mux)
	}

	return mainMiddleware(mux)
}
