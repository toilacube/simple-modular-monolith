package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MemberServerPort string `mapstructure:"MEMBER_SERVER_PORT"`
	MovieServerPort  string `mapstructure:"MOVIE_SERVER_PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	DBDriver         string `mapstructure:"DB_DRIVER"`
}

func LoadConfig(env string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(fmt.Sprintf(".env.%s", env))
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// fallback to .env.local if file not found
			viper.SetConfigName(".env.local")
			if err := viper.ReadInConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
