package handler

import (
	"github.com/amaraliou/trackr-v2/internal/storage"
	"github.com/amaraliou/trackr-v2/pkg/logger"
)

// Handler ...
type Handler struct {
	pgRepo storage.PostgresInterface
	logger logger.Logger
}

// New ...
func New(pgRepo storage.PostgresInterface, logger logger.Logger) *Handler {
	return &Handler{
		pgRepo: pgRepo,
		logger: logger,
	}
}
