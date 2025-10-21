package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		MemberPort string `mapstructure:"member_port"`
		MoviePort  string `mapstructure:"movie_port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Driver   string `mapstructure:"driver"`
	} `mapstructure:"database"`

	Logger struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logger"`

	JWT struct {
		SecretKey         string `mapstructure:"secret_key"`
		ExpirationMinutes int    `mapstructure:"expiration_minutes"`
	} `mapstructure:"jwt"`
}

type ConfigOptions struct {
	ConfigEnv  string // environment, for example local, dev, prod
	ConfigType string // "env", "yaml", or empty (tries both)
}

var (
	globalConfig *Config
)

func GetConfig() *Config {
	return globalConfig
}

func LoadConfig(cfgOpts ConfigOptions) (*Config, error) {
	v := viper.New()

	setDefaultConfig(v)

	v.AddConfigPath(".")
	if err := tryLoadConfigFile(v, cfgOpts); err != nil {
	}

	if err := v.Unmarshal(&globalConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return globalConfig, nil
}

func tryLoadConfigFile(v *viper.Viper, cfgOpts ConfigOptions) error {
	switch cfgOpts.ConfigType {
	case "env":
		return tryLoadEnvFile(v, cfgOpts.ConfigEnv)
	case "yaml":
		return tryLoadYAMLFile(v, cfgOpts.ConfigEnv)
	default:
		if err := tryLoadEnvFile(v, cfgOpts.ConfigEnv); err == nil {
			return nil
		}
		return tryLoadYAMLFile(v, cfgOpts.ConfigEnv)
	}
}

func tryLoadEnvFile(v *viper.Viper, env string) error {
	if env != "" {
		v.SetConfigName(fmt.Sprintf(".env.%s", env))
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err == nil {
			return nil
		}
	}

	// fallback to default .env.local
	v.SetConfigName(".env.local")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("no .env file found")
	}
	return nil
}

func tryLoadYAMLFile(v *viper.Viper, env string) error {
	if env != "" {
		v.SetConfigName(fmt.Sprintf("config.%s", env))
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err == nil {
			return nil
		}
	}

	// fallback to default config.local.yaml
	v.SetConfigName("config.local")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("no yaml file found")
	}
	return nil
}

func setDefaultConfig(v *viper.Viper) {
	v.SetDefault("server.member_port", "6363")
	v.SetDefault("server.movie_port", "8889")

	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", "3308")
	v.SetDefault("database.user", "root")
	v.SetDefault("database.password", "password")
	v.SetDefault("database.name", "tutorial_db")
	v.SetDefault("database.driver", "mysql")

	v.SetDefault("logger.level", "info")
	v.SetDefault("logger.format", "console")
}
