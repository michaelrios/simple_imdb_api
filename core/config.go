package core

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Config struct {
	LoggingConfig
	ServerConfig
	MongoConfig
}

type LoggingConfig struct {
	LoggingLevel int `env:"LOGGING_LEVEL" envDefault:"5"`
}

type ServerConfig struct {
	Addr string `env:"SERVER_ADDR" envDefault:":8080"`
}

func MustSetConfigs() *Config {
	config := Config{}

	if err := env.Parse(&config.ServerConfig); err != nil {
		logrus.WithError(err).Fatal("Failed to import server configs")
	}

	if err := env.Parse(&config.LoggingConfig); err != nil {
		logrus.WithError(err).Fatal("Failed to import logging configs")
	}

	if err := env.Parse(&config.MongoConfig); err != nil {
		logrus.WithError(err).Fatal("Failed to import mongo configs")
	}

	logrus.Info(config)

	return &config
}
