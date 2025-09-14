package bootstrap

import "github.com/Wuchieh/go-server/internal/utils/logger"

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
	}()
}
