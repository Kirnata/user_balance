package store

type Config struct {
	DatabaseURL string `toml:"data_base_url"`
}

func NewConfig() *Config {
	return &Config{}
}
