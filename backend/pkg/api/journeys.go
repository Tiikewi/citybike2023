package api

import (
	"citybike/pkg/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const PAGE_LIMIT int = 10

func handleJourneys(r chi.Router) {
	r.Get("/page/{page}", getJourneys)
	r.Get("/page/{page}/{sort}", getJourneys)
}

// @Summary Get journeys by page.
// @Description GET /api/journeys/{page} returns 10 * page amount of journeys.
// @Tags Journeys
// @Produce json
// @Param page path int true "Page number"
// @Router /api/journeys/page/{page} [get]
// @Success 200 {object} types.Journey
// @Failure 404 {object} types.ErrorResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
func getJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONError("Invalid parameters", http.StatusBadRequest, w)
		return
	}

	sort := chi.URLParam(r, "sort")
	sortInt, err := strconv.Atoi(sort)
	if err != nil {
		sortInt = 0
	}

	journeys, err := db.GetJourneys(pageInt, PAGE_LIMIT, sortInt)
	if err != nil {
		sendJSONError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if len(journeys) == 0 {
		sendJSONError("No journeys found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(journeys)
}
