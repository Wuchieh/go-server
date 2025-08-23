package bootstrap

import (
	redis "github.com/Wuchieh/go-server-redis"
	"github.com/Wuchieh/go-server/internal/config"
	redis2 "github.com/redis/go-redis/v9"
)

func redisSetup() error {
	return redis.Setup(config.GetConfig().Redis)
}

func closeRedis() error {
	return redis.Use(func(client *redis2.Client) error {
		return client.Close()
	}, true)
}
