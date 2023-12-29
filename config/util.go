package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Configuration Config

type Config struct {
	Redis struct {
		Host string
		Port string
	}
}

func Init() {
	config := &Config{}

	config.Load()

	Configuration = *config
}

func (c *Config) Load() {

	godotenv.Load()

	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port = os.Getenv("REDIS_PORT")
}
