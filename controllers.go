package main

import (
	"net/http"
)

func indexController(w http.ResponseWriter, r *http.Request) {

	//View.Render(w, "home.jet", nil, nil)

}

func createGif(w http.ResponseWriter, r *http.Request) {

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
