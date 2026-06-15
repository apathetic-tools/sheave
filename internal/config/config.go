package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apathetic-tools/sheave/registry"
	"github.com/pelletier/go-toml/v2"
)

// Selection represents a scoped include/exclude list
type Selection struct {
	Include []string `toml:"include,omitempty"`
	Exclude []string `toml:"exclude,omitempty"`
}

type FieldProp struct {
	Required bool   `toml:"required,omitempty"`
	Name     string `toml:"name,omitempty"`
	Type     string `toml:"type,omitempty"`
}

type ComponentConfig struct {
	Spread      string   `toml:"spread"`
	Path        string   `toml:"path"`
	Ext         string   `toml:"ext,omitempty"`
	Type        string   `toml:"type,omitempty"`
	Flavor      any      `toml:"flavor,omitempty"`
	Flavour     any      `toml:"flavour,omitempty"`
	Frontmatter []string `toml:"frontmatter,omitempty"`

	Name        *FieldProp `toml:"name,omitempty"`
	Description *FieldProp `toml:"description,omitempty"`
	Invocable   *FieldProp `toml:"invocable,omitempty"`
	Metadata    *FieldProp `toml:"metadata,omitempty"`
}

// ProviderConfig represents a deployment target for AI
type ProviderConfig struct {
	TargetDir   string           `toml:"target_dir"`
	Rules       *ComponentConfig `toml:"rules,omitempty"`
	Skills      *ComponentConfig `toml:"skills,omitempty"`
	Settings    *ComponentConfig `toml:"settings,omitempty"`
	Main        *ComponentConfig `toml:"main,omitempty"`
	Hooks       *ComponentConfig `toml:"hooks,omitempty"`
	MCP         *ComponentConfig `toml:"mcp,omitempty"`
	Environment *ComponentConfig `toml:"environment,omitempty"`
	IDE         *ComponentConfig `toml:"ide,omitempty"`
	Ignore      *ComponentConfig `toml:"ignore,omitempty"`
}

// Config represents the schema of .sheave.toml
type Config struct {
	ActiveProviders []string                  `toml:"active_providers,omitempty"`
	Rules           Selection                 `toml:"rules,omitempty"`
	Skills          Selection                 `toml:"skills,omitempty"`
	Templates       Selection                 `toml:"templates,omitempty"`
	Workflows       Selection                 `toml:"workflows,omitempty"`
	Providers       map[string]ProviderConfig `toml:"providers,omitempty"`
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

// Load reads and parses the configuration file at the given path, merging it with the defaults.
// If the file does not exist, it returns the default configuration.
func Load(path string) (*Config, error) {
	// 1. Load Defaults
	defaultBytes, err := registry.FS.ReadFile(".sheave.toml")
	if err != nil {
		return nil, fmt.Errorf("failed to read default config from registry: %w", err)
	}

	var baseCfg Config
	if err := toml.Unmarshal(defaultBytes, &baseCfg); err != nil {
		return nil, fmt.Errorf("failed to parse default TOML: %w", err)
	}

	// 2. Load User Config
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &baseCfg, nil
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var userCfg Config
	if err := toml.Unmarshal(b, &userCfg); err != nil {
		return nil, fmt.Errorf("failed to parse TOML in %s: %w", path, err)
	}

	// 3. Merge configs
	if len(userCfg.ActiveProviders) > 0 {
		baseCfg.ActiveProviders = userCfg.ActiveProviders
	}
	if len(userCfg.Rules.Include) > 0 {
		baseCfg.Rules.Include = userCfg.Rules.Include
	}
	if len(userCfg.Rules.Exclude) > 0 {
		baseCfg.Rules.Exclude = userCfg.Rules.Exclude
	}
	if len(userCfg.Skills.Include) > 0 {
		baseCfg.Skills.Include = userCfg.Skills.Include
	}
	if len(userCfg.Skills.Exclude) > 0 {
		baseCfg.Skills.Exclude = userCfg.Skills.Exclude
	}
	if len(userCfg.Templates.Include) > 0 {
		baseCfg.Templates.Include = userCfg.Templates.Include
	}
	if len(userCfg.Templates.Exclude) > 0 {
		baseCfg.Templates.Exclude = userCfg.Templates.Exclude
	}
	if len(userCfg.Workflows.Include) > 0 {
		baseCfg.Workflows.Include = userCfg.Workflows.Include
	}
	if len(userCfg.Workflows.Exclude) > 0 {
		baseCfg.Workflows.Exclude = userCfg.Workflows.Exclude
	}

	if baseCfg.Providers == nil && len(userCfg.Providers) > 0 {
		baseCfg.Providers = make(map[string]ProviderConfig)
	}
	for k, v := range userCfg.Providers {
		baseCfg.Providers[k] = v
	}

	return &baseCfg, nil
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
