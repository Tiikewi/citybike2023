package api

import (
	"citybike/pkg/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func handleStations(r chi.Router) {
	r.Get("/page/{page}", getStations)
	r.Get("/page/{page}/{name}", getStationsByName)
}

// @Summary Get journeys by page.
// @Description getJourneys returns 10 * page amount of journeys.
// @Tags Journeys
// @Produce json
// @Router /journeys/:{page} [get]
// @Success 200 {object} types.Journey
// @Failure 404 {object} types.ErrorResponse
// @Failure 400 {object} types.ErrorResponse
func getStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONError("Invalid parameter.", http.StatusBadRequest, w)
		return
	}

	stations, err := db.GetStations(pageInt, PAGE_LIMIT)
	if err != nil {
		sendJSONError(err.Error(), http.StatusInternalServerError, w)
		return
	}
	if len(stations) == 0 {
		sendJSONError("No stations found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}

func getStationsByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	name := chi.URLParam(r, "name")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONError("Invalid parameters", http.StatusBadRequest, w)
		return
	}

	stations, err := db.GetStationsByName(pageInt, PAGE_LIMIT, name)
	if err != nil {
		sendJSONError(err.Error(), http.StatusInternalServerError, w)
		return
	}
	if len(stations) == 0 {
		sendJSONError("No stations found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}
