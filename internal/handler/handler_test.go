package handler

import (
	"log"
	"os"
	"testing"

	"github.com/amaraliou/trackr-v2/internal/storage/mock"
	"github.com/amaraliou/trackr-v2/pkg/logger"
)

var handler *Handler
var mockRepo *mock.Repository

func TestMain(m *testing.M) {

	logConfig := logger.Config{
		EnableConsole:     false,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: true,
		EnableFile:        false,
		FileLevel:         logger.Info,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}

	logger, err := logger.NewZapLogger(logConfig)
	if err != nil {
		log.Fatal(err)
	}

	handler = New(mockRepo, logger)
	os.Exit(m.Run())
}
