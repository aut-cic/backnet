package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New creates a new postgres connection and tests it.
func New(cfg Config) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(url), new(gorm.Config))
	if err != nil {
		return nil, fmt.Errorf("cannot open a database connection %w", err)
	}

	underDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("underlying database connection is not *db.DB %w", err)
	}

	if err := underDB.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping the database %w", err)
	}

	return db, nil
}
