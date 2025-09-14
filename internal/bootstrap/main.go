package bootstrap

import (
	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
	}()

	if err := redisSetup(); err != nil {
		logger.Fatalf("redis setup failed: %v", err)
	}
	defer func() {
		if err := closeRedis(); err != nil {
			logger.Errorf("redis close failed: %v", err)
		}
	}()
}
