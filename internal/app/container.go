package app

import (
	"os"
	"tutorial/internal/member"
	"tutorial/pkg/config"
	"tutorial/pkg/database"
	"tutorial/pkg/logger"

	"gorm.io/gorm"
)

type AppContainer struct {
	Config          *config.Config
	DB              *gorm.DB
	MemberContainer *member.MemberContainer
}

func NewAppContainer() (app *AppContainer, err error) {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	configType := os.Getenv("CONFIG_TYPE")
	if configType == "" {
		configType = "env"
	}

	cfg, err := config.LoadConfig(config.ConfigOptions{
		ConfigEnv:  env,
		ConfigType: configType,
	})

	if err != nil {
		return nil, err
	}

	db, err := database.LoadMySQL(cfg)

	if err != nil {
		return nil, err
	}

	err = database.LoadMigration(db)

	if err != nil {
		return nil, err
	}

	logger.LoadLogger(cfg)

	logger := logger.GetLogger()

	// test logger
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Debug("This is a debug message")

	memberContainer := member.NewMemberContainer(db)

	app = &AppContainer{
		Config:          cfg,
		DB:              db,
		MemberContainer: memberContainer,
	}

	return app, nil
}
