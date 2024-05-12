package http

import "net/http"

// HandlerAddSinglePersonAndMatch Add a new user to the matching system and find any possible matches for the new user
func (s *Server) HandlerAddSinglePersonAndMatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		persons, err := s.personService.AddPersonAndMatch(nil)
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, persons, http.StatusOK)
	}
}

// HandlerRemoveSinglePerson Remove a user from the matching system so that the user cannot be matched anymore
func (s *Server) HandlerRemoveSinglePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		err := s.personService.RemovePerson("")
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, nil, http.StatusOK)
	}
}

// HandlerQuerySinglePeople Remove a user from the matching system so that the user cannot be matched anymore
func (s *Server) HandlerQuerySinglePeople() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		persons, err := s.personService.QuerySinglePeople(5)
		if err != nil {
			s.respond(r, w, err, http.StatusInternalServerError)
			return
		}

		s.respond(r, w, persons, http.StatusOK)
	}
}
