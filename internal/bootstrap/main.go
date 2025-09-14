package bootstrap

import (
	"context"

	mongo "github.com/Wuchieh/go-server-mongo"
	"github.com/Wuchieh/go-server/internal/utils/logger"
)
import (
	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
	}()

	if err := databaseSetup(); err != nil {
		logger.Fatal("database setup fail:", err)
	}
	defer func() {
		err := databaseClose()
		if err != nil {
			logger.Error("database close fail:", err)
		}
	}()

	if err := mongoSetup(); err != nil {
		logger.Fatalf("mongodb setup error: %v", err)
	}
	defer func() {
		err := mongo.GetClient().Disconnect(context.Background())
		if err != nil {
			logger.Errorf("mongodb disconnect error: %v", err)
		}
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
