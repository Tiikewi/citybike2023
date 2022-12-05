package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestGetRouter ensures that the router contains all expected routes
func TestGetRouter(t *testing.T) {
	log, _ := zap.NewProduction(zap.WithCaller(false))
	r := CreateNewServer(log)

	testcases := map[string]struct {
		method string
		path   string
	}{
		"GET /swagger": {
			method: http.MethodGet,
			path:   "/api/swagger/",
		},
		"GET /api/journeys/page/1": {
			method: http.MethodGet,
			path:   "/api/journeys/page/1",
		},
		"GET /api/journeys/page/1/1": {
			method: http.MethodGet,
			path:   "/api/journeys/page/1/1",
		},
		"GET /api/stations/page/1": {
			method: http.MethodGet,
			path:   "/api/stations/page/1",
		},
		"GET /api/stations/page/1/test": {
			method: http.MethodGet,
			path:   "/api/stations/page/1/test",
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			got := r.Router.Match(chi.NewRouteContext(), test.method, test.path)
			assert.Equal(t, true, got, fmt.Sprintf("not found: %s '%s'", test.method, test.path))
		})
	}
}
