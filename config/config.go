package config

import (
	"errors"
)

// Config holds the stapler server settings.
type Config struct {
	MapPath string // Path to the route definition file
	Port    int    // TCP port number to listen on
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) Validate() error {
	if len(cfg.MapPath) == 0 {
		return errors.New("--map option is required")
	}

	return nil
}
