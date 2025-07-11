package middlewares

import "net/http"

func ChainMiddlewares(
	h http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
