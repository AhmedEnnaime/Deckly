package application

import (
	"deckly/pkg/config"
	"deckly/pkg/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	database, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}
	return &Application{
		DB:  database,
		Cfg: cfg,
	}, nil
}
