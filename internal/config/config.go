package config

import (
	"os"
	"sync"
	"user-rest-api/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Listen struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Database struct {
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `env:"DB_PASSWORD"`
	} `yaml:"db"`
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
		instance.Database.Password = os.Getenv("DB_PASSWORD")
	})

	return instance
}
