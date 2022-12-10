package api

import (
	"citybike/pkg/db"
	"citybike/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func handleStations(r chi.Router) {
	r.Get("/page/{page}", getStations)
	r.Get("/page/{page}/{name}", getStationsByName)
	r.Post("/", addStation)
}

// @Summary Get stations by page.
// @Description GET /api/stations/{page} returns 10 * page amount of stations.
// @Tags Stations
// @Produce json
// @Param page path int true "Page number"
// @Router /api/stations/page/{page} [get]
// @Success 200 {object} types.Station
// @Failure 404 {object} types.JSONResponse
// @Failure 400 {object} types.JSONResponse
// @Failure 500 {object} types.JSONResponse
func getStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONResponse("Invalid parameter.", http.StatusBadRequest, w)
		return
	}

	stations, err := db.GetStations(pageInt, PAGE_LIMIT)
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusInternalServerError, w)
		return
	}
	if len(stations) == 0 {
		sendJSONResponse("No stations found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}

// @Summary Get stations by name.
// @Description GET api/stations/{page}/{string}
// returns 10 * page amount of stations by given name.
// Start substring of name also ok.
// @Tags Stations
// @Produce json
// @Param page path int true "Page number"
// @Param name path string true "substring/name of station"
// @Router /api/stations/page/{page}/{name} [get]
// @Success 200 {object} types.Station
// @Failure 404 {object} types.JSONResponse
// @Failure 400 {object} types.JSONResponse
// @Failure 500 {object} types.JSONResponse
func getStationsByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	page := chi.URLParam(r, "page")
	name := chi.URLParam(r, "name")
	pageInt, err := strconv.Atoi(page)

	if err != nil || pageInt < 1 {
		sendJSONResponse("Invalid parameters", http.StatusBadRequest, w)
		return
	}

	stations, err := db.GetStationsByName(pageInt, PAGE_LIMIT, name)
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusInternalServerError, w)
		return
	}
	if len(stations) == 0 {
		sendJSONResponse("No stations found", http.StatusNotFound, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}

// @Summary Add new station.
// @Description POST api/stations to add new station.
// @Tags Stations
// @Produce json
// @Router /api/stations [post]
// @Body types.Station
// @Param   addStationRequest body types.StationRequest true "New station"
// @Success 200 {object} types.Station
// @Failure 404 {object} types.JSONResponse
// @Failure 400 {object} types.JSONResponse
// @Failure 500 {object} types.JSONResponse
func addStation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var station types.StationRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&station)
	if err != nil {
		sendJSONResponse("Invalid body", http.StatusBadRequest, w)
		return
	}

	validate := validator.New()
	if err := validate.Struct(station); err != nil {
		sendJSONResponse("Invalid body", http.StatusBadRequest, w)
		return
	}

	err = db.AddStation(&station)
	if err != nil {
		sendJSONResponse(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	sendJSONResponse("Station added succesfully!", http.StatusOK, w)
}
