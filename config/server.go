package config

import (
	"standard-struct-golang/appconst"
	"standard-struct-golang/packages/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type serverConfig struct {
	ListenIP     string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	ServerHeader string
	ProxyHeader  string
	EnableCORS   bool
}

func NewServerCfg(viper *viper.Viper) serverConfig {
	//set ค่า default ให้ Env ของ Server
	viper.SetDefault("server.listen", appconst.ServerListen)
	viper.SetDefault("server.port", appconst.ServerPort)
	viper.SetDefault("server.timeout.read", appconst.ServerTimeoutRead)
	viper.SetDefault("server.timeout.write", appconst.ServerTimeoutWrite)
	viper.SetDefault("server.timeout.idle", appconst.ServerTimeoutIdle)
	viper.SetDefault("server.header", viper.GetString("app.name"))
	viper.SetDefault("server.proxy.header", fiber.HeaderXForwardedFor)
	viper.SetDefault("server.enable.cors", "false")

	serverConfig := serverConfig{
		ListenIP:     viper.GetString("server.listen"),
		Port:         viper.GetString("server.port"),
		ReadTimeout:  util.ParseDuration(viper.GetString("server.timeout.read")),
		WriteTimeout: util.ParseDuration(viper.GetString("server.timeout.write")),
		IdleTimeout:  util.ParseDuration(viper.GetString("server.timeout.idle")),
		ServerHeader: viper.GetString("server.header"),
		ProxyHeader:  viper.GetString("server.proxy.header"),
		EnableCORS:   viper.GetBool("server.enable.cors"),
	}
	return serverConfig
}
