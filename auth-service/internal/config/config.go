package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string      `yaml:"env"`
	LogLevel    string      `yaml:"log_level"`
	GRPC        GRPC        `yaml:"grpc"`
	UserService UserService `yaml:"user_service"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type UserService struct {
	GRPCPort    int           `yaml:"grpc_port"`
	GRPCTimeout time.Duration `yaml:"grpc_timeout"`
	GRPCRetries time.Duration `yaml:"grpc_retries"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic("config file is not exists")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file")
	}

	return &cfg
}

func fetchConfigPath() string {
	var flagPath string

	flag.StringVar(&flagPath, "config", "", "path to config file")
	flag.Parse()

	if flagPath == "" {
		flagPath = os.Getenv("config")
	}

	return flagPath
}
