package config

// Config
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"sessions_key"`
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		DatabaseURL: "postgres://postgres:mypass@localhost:5432/rest_api",
		SessionKey:  "key",
	}
}
