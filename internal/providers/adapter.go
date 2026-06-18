package providers

// Adapter defines the interface for mapping generic sheave settings to
// provider-specific configurations.
type Adapter interface {
	// TranslateSettings takes a map of generic settings and returns a map
	// containing the provider-specific layout.
	TranslateSettings(settings map[string]any) (map[string]any, error)
}

// DefaultAdapter is a fallback adapter that passes settings through with no changes.
type DefaultAdapter struct{}

func (d *DefaultAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	return settings, nil
}

// GetAdapter returns the specific adapter for a given flavor.
func GetAdapter(flavor string) Adapter {
	switch flavor {
	case "claude":
		return &ClaudeAdapter{}
	default:
		return &DefaultAdapter{}
	}
}
