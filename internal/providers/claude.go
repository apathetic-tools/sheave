package providers

import (
	"fmt"

	"github.com/apathetic-tools/sheave/internal/config"
)

// ClaudeAdapter maps generic settings to Claude's specific configuration.
type ClaudeAdapter struct{}

func (c *ClaudeAdapter) GetLayout() ProviderLayout {
	return ProviderLayout{
		TargetDir: ".claude", // https://code.claude.com/docs/en/settings
		Components: map[string]*config.ComponentConfig{
			"rules": {
				Spread: "dir",
				Path:   "rules",
				Ext:    "md",
				Type:   "md",
			},
			"skills": {
				Spread: "dir",
				Path:   "skills",
				Ext:    "md",
				Type:   "md",
			},
			// https://code.claude.com/docs/en/settings#subagent-configuration
			"agents": {
				Spread: "dir",
				Path:   "agents",
				Ext:    "md",
				Type:   "md",
			},
			// https://code.claude.com/docs/en/mcp
			"mcp": {
				Spread: "file",
				Path:   "//.mcp.json",
				Type:   "json",
				Flavor: "claude",
			},
			// https://code.claude.com/docs/en/settings#settings-files
			// https://code.claude.com/docs/en/settings#plugin-configuration
			"settings": {
				Spread: "file",
				Path:   "settings.json",
				Type:   "json",
				Flavor: "claude",
			},
			"main": {
				Spread:  "file",
				Path:    "CLAUDE.md",
				Type:    "md",
				Flavour: []string{"catchall"},
			},
		},
	}
}

func (c *ClaudeAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	result := make(map[string]any)
	result["$schema"] = "https://json.schemastore.org/claude-code-settings.json"

	if val, ok := settings["co-author-by-ai"]; ok {
		if b, isBool := val.(bool); isBool {
			result["includeCoAuthoredBy"] = b
		}
	}

	permissions := make(map[string]any)
	for _, section := range []string{"allow", "deny", "ask"} {
		var entries []string
		if val, ok := settings["shell-"+section]; ok {
			for _, s := range toStringSlice(val) {
				entries = append(entries, fmt.Sprintf("Bash(%s)", s))
			}
		}
		if val, ok := settings["skill-"+section]; ok {
			entries = append(entries, toStringSlice(val)...)
		}
		if len(entries) > 0 {
			permissions[section] = entries
		}
	}
	if len(permissions) > 0 {
		result["permissions"] = permissions
	}

	return result, nil
}

func toStringSlice(val any) []string {
	switch v := val.(type) {
	case []any:
		var out []string
		for _, item := range v {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return v
	}
	return nil
}
