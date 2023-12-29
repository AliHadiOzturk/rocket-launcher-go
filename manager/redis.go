package manager

import (
	"fmt"
	"launcher/config"

	"github.com/redis/go-redis/v9"
)

var RedisManager Manager

type Manager struct {
	Client redis.Client
}

func NewRedisManager() {
	m := &Manager{
		Client: *redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", config.Configuration.Redis.Host, config.Configuration.Redis.Port),
		}),
	}

	RedisManager = *m
}
