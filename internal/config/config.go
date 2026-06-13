package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config represents the schema of .sheave.toml
type Config struct {
	Select       []string `toml:"select,omitempty"`
	Ignore       []string `toml:"ignore,omitempty"`
	ExtendSelect []string `toml:"extend-select,omitempty"`
}

// Load reads and parses the configuration file at the given path.
// If the file does not exist, it returns an empty configuration.
func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{}, nil
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var cfg Config
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse TOML in %s: %w", path, err)
	}

	return &cfg, nil
}

// Save writes the configuration to the given path.
func (c *Config) Save(path string) error {
	b, err := toml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := os.WriteFile(path, b, 0644); err != nil {
		return fmt.Errorf("failed to write config to %s: %w", path, err)
	}
	return nil
}

// Resolve returns the final set of active rules after applying select, extend-select, and ignore.
func (c *Config) Resolve() []string {
	active := make(map[string]bool)

	// Base selection
	for _, s := range c.Select {
		active[s] = true
	}

	// Extensions
	for _, s := range c.ExtendSelect {
		active[s] = true
	}

	// Ignores
	for _, s := range c.Ignore {
		delete(active, s)
	}

	result := make([]string, 0, len(active))
	for k := range active {
		result = append(result, k)
	}
	return result
}
