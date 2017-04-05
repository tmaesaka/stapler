package config

// Config holds the stapler server settings.
type Config struct {
	MapPath string // Path to the route definition file
	Port    int    // TCP port number to listen on
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{}
}
