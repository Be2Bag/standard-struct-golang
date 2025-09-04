package config

import (
	"github.com/spf13/viper"
)

type healthIdConfig struct {
	Url               string
	ClientId          string
	SecretKey         string
	RedirectUrl       string
	RedirectLocalhost string
	MophUrl           string
	MophClientId      string
	MophSecret        string
	Timeout           int
}

func NewHealthIdCfg(v *viper.Viper) healthIdConfig {
	healthIdconfig := healthIdConfig{
		Url:               v.GetString("health_id.url"),
		ClientId:          v.GetString("health_id.client_id"),
		SecretKey:         v.GetString("health_id.secret_key"),
		RedirectUrl:       v.GetString("health_id.redirect_url"),
		RedirectLocalhost: v.GetString("health_id.redirect_localhost"),
		MophUrl:           v.GetString("moph.moph_url"),
		MophClientId:      v.GetString("moph.client_id"),
		MophSecret:        v.GetString("moph.secret_key"),
		Timeout:           v.GetInt("health_id.timeout"),
	}
	return healthIdconfig
}
