package bootstrap

import (
	mongo "github.com/Wuchieh/go-server-mongo"
	"github.com/Wuchieh/go-server/internal/config"
)

func mongoSetup() error {
	return mongo.Setup(config.GetConfig().Mongo)
}
