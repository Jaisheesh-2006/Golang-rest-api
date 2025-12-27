package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		//* if no CONFIG_PATH is set, we will check for any flags passed to the application

		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("No configuration file path provided. Set CONFIG_PATH environment variable or use -config flag")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at this path: %s", configPath)
	}

	//* Now we can proceed to load the configuration from the specified path
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return &cfg
}
