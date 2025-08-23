package bootstrap

import (
	"os"

	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer logger.Sync()

	if err := redisSetup(); err != nil {
		logger.Errorf("redis setup failed: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := closeRedis(); err != nil {
			logger.Errorf("redis close failed: %v", err)
		}
	}()
}
