package apiserver

import "github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"

type Config struct {
	BindAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
