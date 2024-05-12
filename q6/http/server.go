package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/rs/cors"

	person "tinder"

	"github.com/go-chi/chi/v5"
)

// Server is an HTTP server which embeds a chi router
type Server struct {
	router        *chi.Mux
	personService person.PersonService
}

// NewServer creates and configures a new Server instance
func NewServer(personService person.PersonService) *Server {
	s := Server{
		router:        chi.NewRouter(),
		personService: personService,
	}

	s.router.Use(corsMiddleware().Handler)
	s.routes()
	return &s
}

func corsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

// ErrorResponse is responsible to include multiple errors
type ErrorResponse struct {
	Errors []string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func (s *Server) handleError(w http.ResponseWriter, err error, status int) {
	errs := []string{err.Error()}
	for e := errors.Unwrap(err); e != nil; e = errors.Unwrap(e) {
		errs = append(errs, e.Error())
	}
	s.respond(w, ErrorResponse{Errors: errs}, status)
}
