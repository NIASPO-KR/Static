package server

import (
	"github.com/go-chi/chi/v5"

	"static/internal/handlers/items"
	"static/internal/handlers/payments"
	"static/internal/handlers/pickuppoints"
)

func (s *Server) initRouter() {
	s.router = chi.NewRouter()

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/items", s.registerItemsRoutes)
		r.Route("/pickupPoints", s.registerPickupPointsRoutes)
		r.Route("/payments", s.registerPaymentsRoutes)
	})
}

func (s *Server) registerItemsRoutes(r chi.Router) {
	r.Get("/", items.GetItems(s.itemsUseCase))
}

func (s *Server) registerPickupPointsRoutes(r chi.Router) {
	r.Get("/", pickuppoints.GetPickupPoints(s.pickupPointsUseCase))
}

func (s *Server) registerPaymentsRoutes(r chi.Router) {
	r.Get("/", payments.GetPayments(s.paymentsUseCase))
}
