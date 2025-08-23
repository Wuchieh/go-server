package bootstrap

import (
	orm "github.com/Wuchieh/go-server-orm"
	"github.com/Wuchieh/go-server/internal/config"
)

func databaseSetup() error {
	return orm.Setup(config.GetConfig().Database)
}

func databaseClose() error {
	return orm.Close()
}
