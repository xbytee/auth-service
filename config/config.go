package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	App App `yaml:"app"`
}

type App struct {
	Mode   string `yaml:"mode"`
	Logger Logger `yaml:"logger"`
	Grpc   Grpc   `yaml:"grpc"`
}

type Logger struct {
	Mode string `yaml:"mode"`
}

type Grpc struct {
	Port    int `yaml:"port"`
	Timeout int `yaml:"timeout"`
}

func ReadConfig(path string) Config {
	if _, err := os.Stat(path); err != nil {
		panic("Err not exist config.yaml file: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Err read config file: " + err.Error())
	}

	return cfg
}
