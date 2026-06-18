package providers

import (
	"github.com/apathetic-tools/sheave/internal/config"
)

// CursorAdapter maps generic settings to Cursor's specific configuration.
type CursorAdapter struct{}

func (c *CursorAdapter) GetLayout() ProviderLayout {
	return ProviderLayout{
		TargetDir: ".cursor",
		Components: map[string]*config.ComponentConfig{
			// https://cursor.com/docs/skills
			"skills": {
				Spread:      "dir",
				Path:        "skills",
				Ext:         "md",
				Type:        "md",
				Frontmatter: []string{"name", "description", "paths", "invocable", "metadata"},
				Name:        &config.FieldProp{Required: true},
				Description: &config.FieldProp{Required: true},
				Invocable:   &config.FieldProp{Name: "disable-model-invocation"},
				Metadata:    &config.FieldProp{Type: "kv"},
			},
			// https://cursor.com/docs/rules
			"rules": {
				Spread:      "dir",
				Path:        "rules",
				Ext:         "mdc",
				Type:        "md",
				Frontmatter: []string{"description", "alwaysApply"},
			},
			// https://cursor.com/docs/hooks
			"hooks": {
				Spread: "file",
				Path:   "hooks.json",
				Type:   "json",
				Flavor: "cursor",
			},
			// https://cursor.com/docs/mcp
			"mcp": {
				Spread: "file",
				Path:   "mcp.json",
				Type:   "json",
				Flavor: "cursor",
			},
			"environment": {
				Spread: "file",
				Path:   "environment.json",
				Type:   "json",
				Flavor: "cursor",
			},
			"ide": {
				Spread: "file",
				Path:   "ide.json",
				Type:   "json",
				Flavor: "cursor",
			},
			"ignore": {
				Spread: "file",
				Path:   ".ignorecursor",
				Type:   "dot",
			},
			// https://cursor.com/docs/rules#agentsmd
			"main": {
				Spread:  "file",
				Path:    "AGENTS.md",
				Type:    "md",
				Flavour: []string{"catchall"},
			},
		},
	}
}

func (c *CursorAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	return settings, nil
}
