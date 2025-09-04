package app

import (
	"standard-struct-golang/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type App struct {
	Config *config.Config
	Router *fiber.App
	log    *logrus.Entry
}

func NewApp(config *config.Config) *App {
	log := logrus.New()
	log.SetLevel(config.AppConfig.LogLevel)
	return &App{
		Config: config,
		log:    log.WithField("package", "app"),
	}
}
