package main

import (
	"encoding/json"
	"net/http"
)

func createGif(w http.ResponseWriter, r *http.Request) {

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(CurrentVersion)
	w.Header().Set("Content-type", "application/json")
	w.Write(body)
}
