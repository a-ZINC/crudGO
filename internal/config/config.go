package config

import (
	"log"
	"os"
	"slices"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env:"ADDRESS" env-default:"localhost:8001"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"prod"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func LoadConfig() *Config {
	var configPath string

	configPath = os.Getenv("CON_PATH")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		args := os.Args
		if len(args) > 1 {
			if index := slices.Index(args, "CONFIG_PATH") + 1; index > 0 {
				configPath = args[index]
			}
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				log.Fatal("Config file not found")
			}
		}
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
