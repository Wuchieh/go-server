package bootstrap

import "github.com/Wuchieh/go-server/internal/utils/logger"

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
}
