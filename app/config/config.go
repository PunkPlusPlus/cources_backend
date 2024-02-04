package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Host string `env:"HOST" env-default:"0.0.0.0"`
		Port string `env:"PORT" env-default:"10000"`
	}
	App struct {
		LogLevel string
	}
	Storage struct {
		Url string `env:"DATABASE_URL" env-default:""`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("config collected")
		instance = &Config{}
		err := cleanenv.ReadEnv(instance)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return instance
}
