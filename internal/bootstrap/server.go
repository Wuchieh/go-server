package bootstrap

import (
	"context"

	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/route"
	"github.com/Wuchieh/go-server/internal/utils/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ser *server.Server

func addSwagger(r gin.IRouter) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		func(c *ginSwagger.Config) {
			c.URL = "/swagger/doc.json"
		},
	))
}

func serverStart(ctx context.Context) error {
	ser = server.New()
	route.Route(ser)
	addSwagger(ser)
	cfg := config.GetDefault().Server
	return ser.RunWithConfig(ctx, &cfg)
}

func serverStop(ctx context.Context) error {
	return ser.Stop(ctx)
}
