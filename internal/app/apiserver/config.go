package apiserver

import "github.com/Kirnata/user_balance/internal/app/store"

// Config ...
type Config struct {
	BindAddr string        `toml:"bind_addr"` // адреc на котором запускаем веб сервер
	LogLevel string        `toml:"log_level"` // уровень логирования
	Store    *store.Config `toml:"store"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "error",
		Store:    store.NewConfig(),
	}
}
