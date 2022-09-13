package config

import "github.com/aut-cic/backnet/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Database db.Config `koanf:"database"`
}
