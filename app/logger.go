package app

import "github.com/sirupsen/logrus"

func (app *App) NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(app.Config.AppConfig.LogLevel)
	return l
}
