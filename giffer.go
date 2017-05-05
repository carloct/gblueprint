package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", healthCheck)
	r.Post("/video", createGifFromVideo)

	http.ListenAndServe(":8080", r)

}
