package bootstrap

import (
	"context"

	mongo "github.com/Wuchieh/go-server-mongo"
	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
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
