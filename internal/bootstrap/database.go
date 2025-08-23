package bootstrap

import (
	orm "github.com/Wuchieh/go-server-orm"
	"github.com/Wuchieh/go-server/internal/config"
)

func setupDatabase() error {
	return orm.Setup(config.GetConfig().Database)
}

func closeDatabase() error {
	return orm.Close()
}
