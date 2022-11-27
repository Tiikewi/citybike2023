package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleJourneys(r chi.Router) {
	r.Get("/", getJourneys)
}

func getJourneys(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	fmt.Println("page: ", page)

	w.WriteHeader(http.StatusOK)
}
