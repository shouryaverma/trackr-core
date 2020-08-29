package handler

import (
	"github.com/amaraliou/trackr-core/internal/storage"
	"github.com/amaraliou/trackr-core/pkg/logger"
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
