package server

import "github.com/go-chi/chi/v5"

func (s *Server) initRouter() {
	s.router = chi.NewRouter()
}
