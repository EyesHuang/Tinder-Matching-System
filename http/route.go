package http

func (s *Server) routes() {
	s.router.Post("/persons", s.HandlerAddSinglePersonAndMatch())
	s.router.Delete("/persons", s.HandlerRemoveSinglePerson())
	s.router.Get("/persons", s.HandlerQuerySinglePeople())
}
