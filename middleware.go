package main

import (
	"fmt"
	"net/http"
)

// Middleware ...
func Middleware(fn func(http.ResponseWriter, *http.Request), o ServerOptions) http.Handler {
	next := http.Handler(http.HandlerFunc(fn))

	return defaultHeaders(next)
}

func defaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fmt.Sprintf("gblueprint"))
		next.ServeHTTP(w, r)
	})
}
