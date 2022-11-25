package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "citybike/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()

	s.MountHandlers()

	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Swagger documentation
	s.Router.Mount("/swagger", httpSwagger.WrapHandler)

	// Mount all handlers here

	s.Router.Route("/ping", handlePing)

}
