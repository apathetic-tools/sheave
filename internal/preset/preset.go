package preset

import (
	"os"
	"path/filepath"
	"strings"
)

// Preset represents a single AI guidance preset
type Preset struct {
	ID          string
	Name        string
	Category    string
	Description string
	Builtin     bool
}

// Registry stores available presets
type Registry struct {
	presets map[string]*Preset
}

// NewRegistry creates a new initialized registry
func NewRegistry() *Registry {
	r := &Registry{
		presets: make(map[string]*Preset),
	}
	r.registerBuiltins()
	return r
}

func (r *Registry) registerBuiltins() {
	// Add some dummy built-ins for Phase 1 as proof of concept
	r.Add(&Preset{
		ID:          "clap",
		Name:        "Clap",
		Category:    "Encouragement",
		Description: "Provides AI encouragement in the right direction",
		Builtin:     true,
	})
	r.Add(&Preset{
		ID:          "mentat",
		Name:        "Mentat",
		Category:    "Role",
		Description: "Adopts the persona of a Dune mentat",
		Builtin:     true,
	})
}

// Add registers a preset
func (r *Registry) Add(p *Preset) {
	r.presets[p.ID] = p
}

// Get returns a preset by ID
func (r *Registry) Get(id string) (*Preset, bool) {
	p, ok := r.presets[id]
	return p, ok
}

// List returns all registered presets
func (r *Registry) List() []*Preset {
	result := make([]*Preset, 0, len(r.presets))
	for _, p := range r.presets {
		result = append(result, p)
	}
	return result
}

// DiscoverCustomPresets finds presets in the local .ai directories
func (r *Registry) DiscoverCustomPresets(workspaceRoot string) error {
	dirs := []string{
		filepath.Join(workspaceRoot, ".ai", "rules"),
		filepath.Join(workspaceRoot, ".ai", "commands"),
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if strings.HasSuffix(name, ".mdc") || strings.HasSuffix(name, ".md") {
				id := strings.TrimSuffix(strings.TrimSuffix(name, ".mdc"), ".md")

				if _, exists := r.presets[id]; !exists {
					// Extremely simple Title casing for dummy names
					humanName := strings.ReplaceAll(id, "_", " ")

					r.Add(&Preset{
						ID:          id,
						Name:        humanName,
						Category:    "Custom",
						Description: "Custom preset from " + dir,
						Builtin:     false,
					})
				}
			}
		}
	}
	return nil
}
