package config

import (
	"github.com/aut-cic/backnet/internal/db"
)

func Default() Config {
	return Config{
		Debug: true,
		Auth: Auth{
			Username: "admin",
			Password: "@dmin123",
		},
		Database: db.Config{
			User:     "opnsense",
			Password: "opnsense@123",
			Port:     "3306",
			Host:     "127.0.0.1",
			DBName:   "radius",
		},
	}
}
