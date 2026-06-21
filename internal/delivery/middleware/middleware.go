package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for middleware := len(middlewares) - 1; middleware >= 0; middleware-- {
			next = middlewares[middleware](next)
		}
		return next
	}
}
