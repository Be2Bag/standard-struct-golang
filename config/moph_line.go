package config

import "github.com/spf13/viper"

type mophLineConfig struct {
	UrlLineLogin         string
	HoscodeLineNoti      string
	PasswordHashLineNoti string
	UserLineNoti         string
	UrlSendNoti          string
	Timeout              int
}

func NewMophLineCfg(v *viper.Viper) mophLineConfig {
	mophLineconfig := mophLineConfig{
		UrlLineLogin:         v.GetString("moph_line.url_login_line_noti"),
		HoscodeLineNoti:      v.GetString("moph_line.hos_code_line_noti"),
		PasswordHashLineNoti: v.GetString("moph_line.password_hash_line_noti"),
		UserLineNoti:         v.GetString("moph_line.user_line_noti"),
		UrlSendNoti:          v.GetString("moph_line.url_send_noti"),
		Timeout:              v.GetInt("moph_line.timeout"),
	}
	return mophLineconfig
}
