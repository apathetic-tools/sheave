package providers

import "github.com/apathetic-tools/sheave/internal/config"

// ProviderLayout defines the directory structure and schema for a specific provider.
type ProviderLayout struct {
	TargetDir  string
	Components map[string]*config.ComponentConfig
}

// Adapter defines the interface for mapping generic sheave settings to
// provider-specific configurations.
type Adapter interface {
	// GetLayout returns the file layout schema for this provider.
	GetLayout() ProviderLayout

	// TranslateSettings takes a map of generic settings and returns a map
	// containing the provider-specific layout.
	TranslateSettings(settings map[string]any) (map[string]any, error)
}

// DefaultAdapter is a fallback adapter that passes settings through with no changes.
type DefaultAdapter struct{}

func (d *DefaultAdapter) GetLayout() ProviderLayout {
	return ProviderLayout{
		TargetDir:  ".default",
		Components: make(map[string]*config.ComponentConfig),
	}
}

func (d *DefaultAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	return settings, nil
}

// GetAdapter returns the specific adapter for a given flavor.
func GetAdapter(flavor string) Adapter {
	switch flavor {
	case "claude":
		return &ClaudeAdapter{}
	case "cursor":
		return &CursorAdapter{}
	case "gemini":
		return &GeminiAdapter{}
	case "openai":
		return &OpenAIAdapter{}
	default:
		return &DefaultAdapter{}
	}
}
