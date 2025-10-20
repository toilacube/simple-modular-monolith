package app

import (
	"os"
	"tutorial/pkg/config"
	"tutorial/pkg/database"

	"gorm.io/gorm"
)

type AppContainer struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewAppContainer() (app *AppContainer, err error) {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)

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

	app = &AppContainer{
		Config: cfg,
		DB:     db,
	}

	return app, nil
}
