package store

type Config struct {
	DatabaseURL string `toml:"database_url"` // строка подключения к бд
}

func NewConfig() *Config {
	return &Config{}
}
