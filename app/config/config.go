package config

import "github.com/spf13/viper"

const (
	// 超时时间
	HTTPReadTimeout  = 120
	HTTPWriteTimeout = 120
)

func GetString(key string) string {
	return viper.GetString(key)
}
