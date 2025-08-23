package bootstrap

import (
	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/utils/logger"
)

func loggerSetup() {
	logger.Setup(config.GetConfig().Log)
}
