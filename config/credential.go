package config

import "github.com/spf13/viper"

type credentialConfig struct {
	AuthEncryptRegisterKey string
}

func NewCredentialCfg(viper *viper.Viper) credentialConfig {
	config := credentialConfig{
		AuthEncryptRegisterKey: viper.GetString("credential.auth_encrypt.register_key"),
	}
	return config
}
