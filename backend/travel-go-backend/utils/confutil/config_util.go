package confutil

import (
	"github.com/spf13/viper"
	"strconv"
)

type AppConfig struct {
	ServiceName         string `mapstructure:"SERVICE_NAME"`
	ServerHost          string `mapstructure:"SERVER_HOST"`
	ServerPort          int    `mapstructure:"PORT"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBName              string `mapstructure:"DB_NAME"`
	DBUsername          string `mapstructure:"DB_USERNAME"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBMongoURL          string `mapstructure:"DB_MONGO_URL"`
	BaseImageURL        string `mapstructure:"BASE_IMAGE_URL"`
	UserServiceEndpoint string `mapstructure:"USER_SERVICE_ENDPOINT"`
	JwtSecretKey        string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpired          int    `mapstructure:"JWT_EXPIRED"`
	JwtRefreshSecretKet string `mapstructure:"JWT_REFRESH_SECRET_KEY"`
	JwtRefreshExpired   int    `mapstructure:"JWT_REFRESH_EXPIRED"`
	RedisHost           string `mapstructure:"REDIS_HOST"`
	RedisPort           string `mapstructure:"REDIS_PORT"`
	RedisPassword       string `mapstructure:"REDIS_PASSWORD"`
	RedisDB             int    `mapstructure:"REDIS_DB"`
	OtpExpired          int    `mapstructure:"OTP_EXPIRED"`
	GGMapApiKey         string `mapstructure:"GG_MAP_API_KEY"`
	BaseUrlService      string `mapstructure:"BASE_URL_SERVICE"`
}

func (a AppConfig) ToInt(target string) int {
	// string to int
	result, err := strconv.Atoi(target)
	if err != nil {
		panic(err)
	}

	return result
}

func LoadConfig(path string, configName string) (config interface{}, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}

	err = viper.Unmarshal(&config)
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
