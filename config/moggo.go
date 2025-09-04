package config

import "github.com/spf13/viper"

type mongoConfig struct {
	Connection   string
	DatabaseName string
}

func NewMongoCfg(viper *viper.Viper) mongoConfig {
	//set ค่า default ให้ Env ของ Mongo
	viper.SetDefault("mongo.connection", "")
	viper.SetDefault("mongo.database", "")

	mongoConfig := mongoConfig{
		Connection:   viper.GetString("mongo.connection"),
		DatabaseName: viper.GetString("mongo.database"),
	}
	return mongoConfig
}
