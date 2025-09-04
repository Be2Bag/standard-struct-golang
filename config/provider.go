package config

import (
	"github.com/spf13/viper"
)

type providerConfig struct {
	Url       string
	Redirect  string
	ClientId  string
	SecretKey string
	Timeout   int
}

func NewProviderCfg(v *viper.Viper) providerConfig {

	providerconfig := providerConfig{
		Url:       v.GetString("provider.url"),
		Redirect:  v.GetString("provider.redirect"),
		ClientId:  v.GetString("provider.client_id"),
		SecretKey: v.GetString("provider.secret_key"),
		Timeout:   v.GetInt("provider.timeout"),
	}
	return providerconfig
}
