package bootstrap

import (
	"fmt"

	"github.com/Wuchieh/go-server/internal/config"
)

func Run() {
	initConfig()

	fmt.Println(config.GetConfig())
}
