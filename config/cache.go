package config

import "github.com/spf13/viper"

type CacheConfig struct {
	Host      string
	Password  string
	Port      int
	KeyPrefix string
}

func NewCacheCfg(viper *viper.Viper) CacheConfig {
	config := CacheConfig{
		Host:      viper.GetString("redis.host"),
		Password:  viper.GetString("redis.password"),
		Port:      viper.GetInt("redis.port"),
		KeyPrefix: viper.GetString("cache.key_prefix"),
	}
	return config
}
