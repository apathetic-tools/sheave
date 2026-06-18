package providers

import (
	"github.com/apathetic-tools/sheave/internal/config"
)

// GeminiAdapter maps generic settings to Gemini's specific configuration.
type GeminiAdapter struct{}

func (g *GeminiAdapter) GetLayout() ProviderLayout {
	return ProviderLayout{
		TargetDir: ".gemini",
		Components: map[string]*config.ComponentConfig{
			"settings": {
				Spread: "file",
				Path:   "settings.json",
				Type:   "json",
				Flavor: "gemini",
			},
			"main": {
				Spread:  "file",
				Path:    "GEMINI.md",
				Type:    "md",
				Flavour: []string{"catchall", "memory"},
			},
		},
	}
}

func (g *GeminiAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	return settings, nil
}
