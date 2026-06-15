package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Selection represents a scoped include/exclude list
type Selection struct {
	Include []string `toml:"include,omitempty"`
	Exclude []string `toml:"exclude,omitempty"`
}

// Config represents the schema of .sheave.toml
type Config struct {
	Rules     Selection `toml:"rules,omitempty"`
	Commands  Selection `toml:"commands,omitempty"`
	Templates Selection `toml:"templates,omitempty"`
	Workflows Selection `toml:"workflows,omitempty"`
}

// GetConfigPath returns the path to the configuration file.
// It checks .ai/.sheave.toml first, then .sheave.toml.
// If neither exists, it defaults to .ai/.sheave.toml if the .ai dir exists, else .sheave.toml.
func GetConfigPath(projectRoot string) string {
	aiPath := filepath.Join(projectRoot, ".ai", ".sheave.toml")
	if _, err := os.Stat(aiPath); err == nil {
		return aiPath
	}
	rootPath := filepath.Join(projectRoot, ".sheave.toml")
	if _, err := os.Stat(rootPath); err == nil {
		return rootPath
	}
	if _, err := os.Stat(filepath.Join(projectRoot, ".ai")); err == nil {
		return aiPath
	}
	return rootPath
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
