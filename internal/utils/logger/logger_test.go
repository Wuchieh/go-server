package logger_test

import (
	"os"
	"testing"

	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/utils/logger"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	os.Mkdir("./log", os.ModePerm)
	cfg := config.Log{
		Level:      "",
		Format:     "",
		OutputPath: "./log/log.log",
		ErrorPath:  "./log/error.log",
	}

	logger.Setup(cfg)

	defer logger.Sync()

	for i := 0; i < 10000; i++ {
		logger.Info("failed to fetch URL",
			zap.Any("data", config.Config{}),
		)

		logger.Error("failed to fetch URL",
			zap.Any("data", config.Config{}))
	}

}
