package config

import (
	"github.com/spf13/viper"
	"gitlab.com/virtual-travel/travel-go-backend/utils/confutil"
)

type Config struct {
	confutil.AppConfig
	OtpExpired int `mapstructure:"OTP_EXPIRED"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, configName string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}

	err = viper.Unmarshal(&config.AppConfig)
	if err != nil {
		panic(err)
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
		return
	}
	return
}
