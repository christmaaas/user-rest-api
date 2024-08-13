package config

import (
	"os"
	"sync"
	"user-rest-api/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Connection ConnectionConfig `yaml:"listen"`
	Storage    StorageConfig    `yaml:"db"`
}

type ConnectionConfig struct {
	BindIP string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type StorageConfig struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `env:"DB_PASSWORD"`
}

var (
	instance *Config = &Config{}
	once     sync.Once
)

const (
	yamlConfigPath = "configs/config.yml"
	envConfigPath  = ".env"
)

func GetConfig() *Config {
	once.Do(func() {
		logger := logger.GetLogger()
		logger.Info("read application config")

		if err := cleanenv.ReadConfig(yamlConfigPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			logger.Info(help)
			logger.Fatal(err)
		}

		if err := godotenv.Load(envConfigPath); err != nil {
			logger.Fatal(err)
		}
		instance.Storage.Password = os.Getenv("DB_PASSWORD")
	})

	return instance
}
