package bootstrap

import (
	"os"

	"context"

	mongo "github.com/Wuchieh/go-server-mongo"
	"github.com/Wuchieh/go-server/internal/utils/logger"
	"go.uber.org/zap"
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

	if err := mongoSetup(); err != nil {
		logger.Errorf("mongodb setup error: %v", err)
	}
	defer func() {
		err := mongo.GetClient().Disconnect(context.Background())
		if err != nil {
			logger.Errorf("mongodb setup error: %v", err)
		}
	}()
}
