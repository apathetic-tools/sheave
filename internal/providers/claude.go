package providers

import "fmt"

// ClaudeAdapter maps generic settings to Claude's specific configuration.
type ClaudeAdapter struct{}

func (c *ClaudeAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	result := make(map[string]any)

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
