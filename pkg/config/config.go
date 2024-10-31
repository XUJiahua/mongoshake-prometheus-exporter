package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type ProbeConfig struct {
	// BaseURL is the base URL of the MongoShake server, e.g. http://localhost:9100
	BaseURL string `toml:"base_url"`
	// Alias is optional, if not set, use BaseURL as alias
	Alias string `toml:"alias"`
	// Timeout is the timeout for each probe, in seconds, default is 1 second
	Timeout int `toml:"timeout"`
	// Interval is the interval between each probe, in seconds, default is 5 seconds
	Interval int `toml:"interval"`
}

type Config struct {
	Debug  bool          `toml:"debug"`
	Probes []ProbeConfig `toml:"probes"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	_, err := toml.DecodeFile(path, cfg)
	return cfg, err
}

func (c *Config) Validate() error {
	for _, probe := range c.Probes {
		if probe.BaseURL == "" {
			return fmt.Errorf("base_url is required")
		}
	}
	return nil
}
