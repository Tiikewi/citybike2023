package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	_ "citybike/docs"
	mid "citybike/pkg/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer(log *zap.Logger) *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.RequestID)
	if log != nil {
		s.Router.Use(mid.SetLogger(log))
	}

	s.MountHandlers()

	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Swagger documentation
	s.Router.Mount("/api/swagger", httpSwagger.WrapHandler)

	// Mount all handlers here

	s.Router.Route("/api/ping", handlePing)

	s.Router.Route("/api/journeys", handleJourneys)

	s.Router.Route("/api/stations", handleStations)

}
