package config

import "github.com/aut-cic/backnet/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Auth     Auth      `koanf:"auth"`
	Database db.Config `koanf:"database"`
}

type Auth struct {
	Username string
	Password string
}
