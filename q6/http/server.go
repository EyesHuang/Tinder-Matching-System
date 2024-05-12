package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/rs/cors"

	person "tinder"

	"github.com/go-chi/chi/v5"
)

// Pinger is an interface that represents any type that can respond to a PingContext.
type Pinger interface {
	PingContext(ctx context.Context) error
}

// Server is an HTTP server which embeds a chi router
type Server struct {
	router        *chi.Mux
	personService person.PersonService
}

// NewServer creates and configures a new Server instance
func NewServer(personService person.PersonService, p ...Pinger) *Server {
	s := Server{
		router:        chi.NewRouter(),
		personService: personService,
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	s.router.Use(corsMiddleware.Handler)
	s.routes(p...)
	return &s
}

// ErrorResponse is responsible to include multiple errors
type ErrorResponse struct {
	Errors []string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(r *http.Request, w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	var body interface{}
	switch v := data.(type) {
	case error:
		res := ErrorResponse{
			Errors: []string{},
		}

		for ce := v; ce != nil; ce = errors.Unwrap(ce) {
			res.Errors = append(res.Errors, ce.Error())
		}
		body = res
	case *[]person.Person:
		body = data
	default:
		body = nil
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Fatalln(err)
		}
	}
}
