package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// ServerOptions ...
type ServerOptions struct {
	Port    int
	Address string
}

// Server ...
func Server(o ServerOptions) error {
	addr := o.Address + ":" + strconv.Itoa(o.Port)
	handler := routes()

	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}

func routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/health", healthCheck)
	r.Post("/gif", createGif)

	return r
}
