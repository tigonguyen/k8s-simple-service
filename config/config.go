// Package config provides data structs and methods to store the
//application's configuration details
package config

import (
	"github.com/caarlos0/env/v6"
)

// Config is responsible for holding the application configuration
// variables. Each configuration point is also exported as an environment
// variable.
type Config struct {
	Port        uint   `env:"PORT" envDefault:"8000"`
	PostgresURL string `env:"POSTGRES_URL" envDefault:"postgres://user:pass@127.0.0.1/simple-service"`
}

// NewConfig returns a config.Config object filled with the environment
// variables values.
func NewConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic("Failed to parse env vars: " + err.Error())
	}

	return cfg
}
