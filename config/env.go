package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config interface {
		Get(key string) string
	}

	configImpl struct{}
)

func (c *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &configImpl{}
}
