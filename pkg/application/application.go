package application

import (
	"github.com/lexscher/imagine/pkg/config"
	"github.com/lexscher/imagine/pkg/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application {
		DB: db,
		Cfg: cfg,
	}, nil
}