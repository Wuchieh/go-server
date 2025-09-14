package bootstrap

import (
	"context"

	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/route"
	"github.com/Wuchieh/go-server/internal/utils/server"
)

var ser *server.Server

func serverStart(ctx context.Context) error {
	ser = server.New()
	route.Route(ser)
	cfg := config.GetDefault().Server
	return ser.RunWithConfig(ctx, &cfg)
}

func serverStop(ctx context.Context) error {
	return ser.Stop(ctx)
}
