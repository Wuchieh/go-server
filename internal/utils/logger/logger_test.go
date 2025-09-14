package logger_test

import (
	"os"
	"testing"

	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func TestLogger(t *testing.T) {
	os.Mkdir("./log", os.ModePerm)
	cfg := config.Log{
		Level:      "",
		Format:     "json",
		OutputPath: "./log/log.log",
		ErrorPath:  "./log/error.log",
	}

	logger.Setup(cfg)

	defer logger.Sync()

	logger.Info("failed to fetch URL")

	logger.Error("failed to fetch URL")
}
