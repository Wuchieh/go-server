package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := serverStart(ctx); err != nil {
			logger.Fatal(err)
		}
	}()
	defer func() {
		if err := serverStart(ctx); err != nil {
			logger.Error(err)
		}
		cancel()
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc
}
