package bootstrap

import (
	"os"

	"github.com/Wuchieh/go-server/internal/utils/logger"
	"go.uber.org/zap"
)

func Run() {
	initConfig()

	loggerSetup()
	defer logger.Sync()

	if err := databaseSetup(); err != nil {
		logger.Errorf("database setup fail: %v", err)
		os.Exit(1)
	}
	defer func() {
		err := databaseClose()
		if err != nil {
			logger.Error("database close fail", zap.Error(err))
		}
	}()
}
