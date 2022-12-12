package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` // адреc на котором запускаем веб сервер
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
