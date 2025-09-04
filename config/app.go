package config

import (
	"standard-struct-golang/appconst"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type appConfig struct {
	Name       string
	Version    string
	Mode       string
	StorageDir string
	PrefixPath string
	LogLevel   logrus.Level
	Env        string
}

func (appCfg *appConfig) IsDebug() bool {
	return appCfg.LogLevel == logrus.DebugLevel
}

func NewAppCfg(viper *viper.Viper, version string) appConfig {
	//set ค่า default ให้ Env ของ Main app
	viper.SetDefault("app.name", "your project")
	viper.SetDefault("app.version", version)
	viper.SetDefault("app.mode", appconst.EnvDevelopment)
	viper.SetDefault("app.storage", "./storage")
	viper.SetDefault("app.prefix.path", "/api")
	viper.SetDefault("app.log.level", "info")
	appConfig := appConfig{
		Name:       viper.GetString("app.name"),
		Version:    viper.GetString("app.version"),
		Mode:       viper.GetString("app.mode"),
		PrefixPath: viper.GetString("app.prefix.path"),
		StorageDir: viper.GetString("app.storage"),
		Env:        viper.GetString("app.env"),
	}
	logrusLevel, errOnSetLogLevel := logrus.ParseLevel(viper.GetString("app.log.level"))
	if errOnSetLogLevel == nil {
		appConfig.LogLevel = logrusLevel
	}

	return appConfig
}
