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

	// Translate co-author-by-ai
	if val, ok := settings["co-author-by-ai"]; ok {
		// Claude expects "includeCoAuthoredBy" (boolean)
		if b, isBool := val.(bool); isBool {
			result["includeCoAuthoredBy"] = b
		}
	}

	// Translate commands-allowed
	if val, ok := settings["commands-allowed"]; ok {
		var allowedCommands []string

		switch v := val.(type) {
		case []any:
			for _, cmd := range v {
				if s, isStr := cmd.(string); isStr {
					allowedCommands = append(allowedCommands, fmt.Sprintf("Bash(%s)", s))
				}
			}
		case []string:
			for _, s := range v {
				allowedCommands = append(allowedCommands, fmt.Sprintf("Bash(%s)", s))
			}
		}

		if len(allowedCommands) > 0 {
			result["permissions"] = map[string]any{
				"allow": allowedCommands,
			}
		}
	}

	return result, nil
}
