package db

type Config struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
}
