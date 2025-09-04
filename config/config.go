package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig        appConfig
	ServerConfig     serverConfig
	MongoConfig      mongoConfig
	CacheConfig      CacheConfig
	MophLine         mophLineConfig
	TracerConfig     TracerConfig
	ProviderConfig   providerConfig
	HealthIdConfig   healthIdConfig
	CredentialConfig credentialConfig
}

func LoadConfig(configPath string, version string) *Config {
	//กำหนด viper instance ใหม่เพื่อนำไปใช้กับ config แต่ละส่วน
	viperInstance := viper.New()
	//load file config มาจาก path
	viperInstance.SetConfigFile(configPath)
	//อ่านไฟล์ config
	errOnRead := viperInstance.ReadInConfig()
	//handle error กรณี read ไม่สำเร็จ
	if errOnRead != nil {
		logrus.Fatalln("load config file error:", errOnRead.Error())
	}

	appConfig := NewAppCfg(viperInstance, version)
	serverConfig := NewServerCfg(viperInstance)
	mongoConfig := NewMongoCfg(viperInstance)
	cacheConfig := NewCacheCfg(viperInstance)
	mophLineConfig := NewMophLineCfg(viperInstance)
	tracerConfig := NewTracerCfg(viperInstance)
	providerConfig := NewProviderCfg(viperInstance)
	healthIdConfig := NewHealthIdCfg(viperInstance)
	credentialConfig := NewCredentialCfg(viperInstance)

	config := &Config{
		AppConfig:        appConfig,
		ServerConfig:     serverConfig,
		MongoConfig:      mongoConfig,
		CacheConfig:      cacheConfig,
		MophLine:         mophLineConfig,
		TracerConfig:     tracerConfig,
		ProviderConfig:   providerConfig,
		HealthIdConfig:   healthIdConfig,
		CredentialConfig: credentialConfig,
	}

	return config
}
