package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env       string         `yaml:"env" env-required:"true"`
	TokenTTL  time.Duration  `yaml:"token_ttl" env-required:"true"`
	Postrgres PostgresConfig `yaml:"postgres"`
	GRPC      GRPCConfig     `yaml:"grpc"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func LoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func Load() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("CONFIG_PARH is not set")
	}

	return LoadByPath(path)
}
