package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"static/config"
	postgresAdapter "static/internal/adapter/postgres"
	"static/internal/adapter/repository/static"
	"static/internal/infrastructure/database"
	"static/internal/infrastructure/database/postgres"
	"static/internal/ports/repository"
	"static/internal/usecase"
	staticService "static/internal/usecase/static"
)

type Server struct {
	cfg *config.Config

	staticDB *postgres.Postgres

	// repositories
	itemsRepository        repository.ItemsRepository
	pickupPointsRepository repository.PickupPointsRepository
	paymentsRepository     repository.PaymentsRepository

	// services
	itemsUseCase        usecase.ItemsUseCase
	pickupPointsUseCase usecase.PickupPointsUseCase
	paymentsUseCase     usecase.PaymentsUseCase

	router *chi.Mux
	server *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	if err := s.init(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) init() error {
	if err := s.initDB(); err != nil {
		return fmt.Errorf("init db: %v", err)
	}
	if err := database.MigrateStaticDB(s.staticDB); err != nil {
		return fmt.Errorf("migrate static db: %v", err)
	}

	s.initRepositories()
	s.initUseCases()
	s.initRouter()
	s.initHTTPServer()

	return nil
}

func (s *Server) initDB() error {
	var err error

	s.staticDB, err = postgresAdapter.Connect(s.cfg.Server.StaticData.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRepositories() {
	s.itemsRepository = static.NewItemsRepository(s.staticDB)
	s.pickupPointsRepository = static.NewPickupPointsRepository(s.staticDB)
	s.paymentsRepository = static.NewPaymentsRepository(s.staticDB)
}

func (s *Server) initUseCases() {
	s.itemsUseCase = staticService.NewItemsUseCase(s.itemsRepository)
	s.pickupPointsUseCase = staticService.NewPickupPointUseCase(s.pickupPointsRepository)
	s.paymentsUseCase = staticService.NewPaymentsUseCase(s.paymentsRepository)
}

func (s *Server) initHTTPServer() {
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.cfg.Server.Addr, s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (s *Server) Run() {
	log.Println("Server started")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
