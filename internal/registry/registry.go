package registry

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed all:builtin
var builtinFS embed.FS

// Item represents a single AI guidance item (Command, Rule, Template, or Workflow)
type Item struct {
	ID          string
	Name        string
	Type        string // "Command", "Rule", "Template", "Workflow"
	Category    string
	Description string
	Builtin     bool
}

// Registry stores available items
type Registry struct {
	items map[string]*Item
}

// NewRegistry creates a new initialized registry
func NewRegistry() *Registry {
	r := &Registry{
		items: make(map[string]*Item),
	}
	r.registerBuiltins()
	return r
}

func (r *Registry) registerBuiltins() {
	fs.WalkDir(builtinFS, "builtin", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		name := d.Name()
		if strings.HasSuffix(name, ".mdc") || strings.HasSuffix(name, ".md") {
			id := strings.TrimSuffix(strings.TrimSuffix(name, ".mdc"), ".md")

			dir := filepath.Dir(path)
			var itemType string
			switch filepath.Base(dir) {
			case "rules":
				itemType = "Rule"
			case "commands":
				itemType = "Command"
			case "templates":
				itemType = "Template"
			case "workflows":
				itemType = "Workflow"
			default:
				return nil
			}

			if _, exists := r.items[id]; !exists {
				humanName := strings.ReplaceAll(id, "_", " ")

				r.Add(&Item{
					ID:          id,
					Name:        humanName,
					Type:        itemType,
					Category:    "Builtin",
					Description: "Built-in " + itemType,
					Builtin:     true,
				})
			}
		}
		return nil
	})
}

// Add registers an item
func (r *Registry) Add(i *Item) {
	r.items[i.ID] = i
}

// Get returns an item by ID
func (r *Registry) Get(id string) (*Item, bool) {
	i, ok := r.items[id]
	return i, ok
}

// List returns all registered items
func (r *Registry) List() []*Item {
	result := make([]*Item, 0, len(r.items))
	for _, i := range r.items {
		result = append(result, i)
	}
	return result
}

// DiscoverCustomItems finds items in the local .ai directories
func (r *Registry) DiscoverCustomItems(workspaceRoot string) error {
	dirs := map[string]string{
		filepath.Join(workspaceRoot, ".ai", "rules"):     "Rule",
		filepath.Join(workspaceRoot, ".ai", "commands"):  "Command",
		filepath.Join(workspaceRoot, ".ai", "templates"): "Template",
		filepath.Join(workspaceRoot, ".ai", "workflows"): "Workflow",
	}

	for dir, itemType := range dirs {
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

				if _, exists := r.items[id]; !exists {
					// Extremely simple Title casing for dummy names
					humanName := strings.ReplaceAll(id, "_", " ")

					r.Add(&Item{
						ID:          id,
						Name:        humanName,
						Type:        itemType,
						Category:    "Custom",
						Description: "Custom " + itemType + " from " + dir,
						Builtin:     false,
					})
				}
			}
		}
	}
	return nil
}
