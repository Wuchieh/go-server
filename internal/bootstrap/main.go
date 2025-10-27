package bootstrap

import (
	"context"
	"fmt"
	"os/signal"

	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func Run() {
	initConfig()

	loggerSetup()
	defer func() {
		_ = logger.Sync()
	}()

	ctx, stop := signal.NotifyContext(context.Background())
	defer stop()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := serverStart(ctx); err != nil {
			logger.Error(fmt.Errorf("server start error: %v", err))
		}
		cancel()
	}()
	defer func() {
		if err := serverStop(ctx); err != nil {
			logger.Error(err)
		}
	}()

	<-ctx.Done()
}
