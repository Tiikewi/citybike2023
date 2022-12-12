package api

import (
	"citybike/pkg/db"
	"citybike/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

const PAGE_LIMIT int = 10

func handleJourneys(r chi.Router) {
	r.Get("/page/{page}", getJourneys)
	r.Get("/page/{page}/{sort}", getJourneys)
	r.Post("/", addJourney)
}

// @Summary Get journeys by page.
// @Description GET /api/journeys/{page} returns 10 * page amount of journeys.
// @Tags Journeys
// @Produce json
// @Param page path int true "Page number"
// @Router /api/journeys/page/{page} [get]
// @Success 200 {object} types.Journey
// @Failure 404 {object} types.JSONResponse
// @Failure 400 {object} types.JSONResponse
// @Failure 500 {object} types.JSONResponse
func getJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONResponse("Invalid parameters", http.StatusBadRequest, w)
		return
	}

	sort := chi.URLParam(r, "sort")
	sortInt, err := strconv.Atoi(sort)
	if err != nil {
		sortInt = 0
	}

	journeys, err := db.GetJourneys(pageInt, PAGE_LIMIT, sortInt)
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if len(journeys) == 0 {
		sendJSONResponse("No journeys found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(journeys)
}

// @Summary Add new journey.
// @Description POST api/journeys to add new journey.
// @Tags Journeys
// @Produce json
// @Router /api/journeys [post]
// @Param   addJourneyRequest body types.JourneyRequest true "New journey"
// @Success 200 {object} types.JSONResponse
// @Failure 400 {object} types.JSONResponse
// @Failure 500 {object} types.JSONResponse
func addJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var journey types.JourneyRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&journey)
	if err != nil {
		sendJSONResponse("Invalid body", http.StatusBadRequest, w)
		return
	}

	validate := validator.New()
	if err := validate.Struct(journey); err != nil {
		sendJSONResponse(err.Error(), http.StatusBadRequest, w)
		return
	}

	_, err = time.Parse(time.RFC3339, journey.DepTime+"Z")
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusBadRequest, w)
		return
	}
	_, err = time.Parse(time.RFC3339, journey.RetTime+"Z")
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	// sendJSONResponse("Station added succesfully!", http.StatusOK, w)
	json.NewEncoder(w).Encode(journey)
}
