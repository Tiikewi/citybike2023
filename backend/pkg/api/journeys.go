package api

import (
	"citybike/pkg/db"
	"citybike/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const PAGE_LIMIT int = 10

func handleJourneys(r chi.Router) {
	r.Get("/page/{page}", getJourneys)
}

// @Summary Get journeys by page.
// @Description getJourneys returns 10 * page amount of journeys.
// @Tags Journeys
// @Produce json
// @Router /journeys/:{page} [get]
// @Success 200 {object} types.Journey
// @Failure 404 {object} types.ErrorResponse
// @Failure 400 {object} types.ErrorResponse
func getJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := types.ErrorResponse{Message: "Invalid parameter.", StatusCode: 400}
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	journeys := db.GetJourneys(pageInt, PAGE_LIMIT)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(journeys)
}
