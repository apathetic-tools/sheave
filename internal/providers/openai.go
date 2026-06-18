package providers

import (
	"github.com/apathetic-tools/sheave/internal/config"
)

// OpenAIAdapter maps generic settings to OpenAI's specific configuration.
type OpenAIAdapter struct{}

func (o *OpenAIAdapter) GetLayout() ProviderLayout {
	return ProviderLayout{
		TargetDir: ".openai",
		Components: map[string]*config.ComponentConfig{
			"main": {
				Spread:  "file",
				Path:    "OPENAI.md",
				Type:    "md",
				Flavour: []string{"catchall", "memory"},
			},
		},
	}
}

func (o *OpenAIAdapter) TranslateSettings(settings map[string]any) (map[string]any, error) {
	return settings, nil
}
