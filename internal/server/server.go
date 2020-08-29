package server

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/amaraliou/trackr-v2/internal/handler"
	"github.com/amaraliou/trackr-v2/internal/storage/postgres"
	"github.com/amaraliou/trackr-v2/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// Server ...
type Server struct {
	Logger  logger.Logger
	Handler *handler.Handler
	Router  *chi.Mux
	Server  *http.Server
}

// NewServer ...
func NewServer() (*Server, error) {

	// To edit once config is set up
	config := postgres.NewConfig(
		os.Getenv("ENVIRONMENT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		5432,
	)

	server := Server{}

	// Initialize logger
	logConfig := logger.Config{
		EnableConsole:     true,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: true,
		EnableFile:        true,
		FileLevel:         logger.Info,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}
	logger, err := logger.NewZapLogger(logConfig)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	server.Logger = logger

	// Initialize Postgres
	pgConn, err := postgres.NewConnection(config, logger)
	if err != nil {
		return nil, err
	}

	// Initialize repos
	pgRepo, err := postgres.NewRepository(pgConn)
	if err != nil {
		return nil, err
	}

	// Initialize handler
	handler := handler.New(pgRepo, server.Logger)
	server.Handler = handler

	// Initialize router
	server.NewRouter(handler)

	// Initialize server
	server.Server = &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
		Handler: server.Router,
	}

	return &server, nil
}

// ListenAndServe ...
func (server *Server) ListenAndServe() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
