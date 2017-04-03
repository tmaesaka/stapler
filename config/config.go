package config

// Config holds the stapler server settings.
type Config struct {
	Port int
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{}
}
