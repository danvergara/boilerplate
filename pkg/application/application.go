package application

import (
	"github.com/danvergara/boilerplate/pkg/config"
	"github.com/danvergara/boilerplate/pkg/db"
)

// Application main object
type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

// Get return an application instance that will hold our configuration and access to database
func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
