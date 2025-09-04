package config

import "github.com/spf13/viper"

type TracerConfig struct {
	Url    string
	Enable bool
}

func NewTracerCfg(viper *viper.Viper) TracerConfig {
	tracerConfig := TracerConfig{
		Url:    viper.GetString("tracer.url"),
		Enable: viper.GetBool("tracer.enable"),
	}

	return tracerConfig
}
