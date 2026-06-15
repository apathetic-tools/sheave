package registry

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed all:builtin
var builtinFS embed.FS

// Item represents a single AI guidance item (Command, Rule, Template, or Workflow)
type Item struct {
	ID          string
	Family      string
	Name        string
	Type        string // "Command", "Rule", "Template", "Workflow"
	Category    string
	Description string
	Builtin     bool
	Content     []byte
}

type frontmatter struct {
	ID     string `yaml:"sheave-id"`
	Family string `yaml:"sheave-family"`
}

func parseFrontmatter(content []byte) (string, string) {
	str := string(content)
	if !strings.HasPrefix(str, "---\n") && !strings.HasPrefix(str, "---\r\n") {
		return "", ""
	}

	start := strings.Index(str, "\n") + 1
	end := strings.Index(str[start:], "\n---")
	if end == -1 {
		return "", ""
	}
	end += start

	var fm frontmatter
	if err := yaml.Unmarshal([]byte(str[start:end]), &fm); err != nil {
		return "", ""
	}
	return fm.ID, fm.Family
}

// Registry stores available items
type Registry struct {
	builtins map[string]*Item
	customs  map[string]*Item
}

// NewRegistry creates a new initialized registry
func NewRegistry() *Registry {
	r := &Registry{
		builtins: make(map[string]*Item),
		customs:  make(map[string]*Item),
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

			var content []byte
			family := ""
			if b, err := fs.ReadFile(builtinFS, path); err == nil {
				content = b
				fmID, fmFamily := parseFrontmatter(b)
				if fmID != "" {
					id = fmID
				}
				family = fmFamily
			}

			// We need a unique key for builtins in case a local item overrides it.
			// Let's use `registryKey = family + "/" + id` but since we want local to override built-in,
			// they will share the exact same key.
			registryKey := id
			if family != "" {
				registryKey = family + "/" + id
			}

			if _, exists := r.builtins[registryKey]; !exists {
				humanName := strings.ReplaceAll(id, "_", " ")

				r.AddBuiltin(&Item{
					ID:          id,
					Family:      family,
					Name:        humanName,
					Type:        itemType,
					Category:    "Builtin",
					Description: "Built-in " + itemType,
					Builtin:     true,
					Content:     content,
				})
			}
		}
		return nil
	})
}

// AddBuiltin registers a built-in item
func (r *Registry) AddBuiltin(i *Item) {
	key := i.ID
	if i.Family != "" {
		key = i.Family + "/" + i.ID
	}
	r.builtins[key] = i
}

// AddCustom registers a custom userland item
func (r *Registry) AddCustom(i *Item) {
	key := i.ID
	if i.Family != "" {
		key = i.Family + "/" + i.ID
	}
	r.customs[key] = i
}

// Get returns an item by ID (custom overrides builtin)
func (r *Registry) Get(id string) (*Item, bool) {
	if i, ok := r.customs[id]; ok {
		return i, true
	}
	i, ok := r.builtins[id]
	return i, ok
}

// List returns all registered items, with customs overriding builtins
func (r *Registry) List() []*Item {
	merged := make(map[string]*Item)
	for k, v := range r.builtins {
		merged[k] = v
	}
	for k, v := range r.customs {
		merged[k] = v
	}

	result := make([]*Item, 0, len(merged))
	for _, i := range merged {
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

				family := ""
				var content []byte
				if b, err := os.ReadFile(filepath.Join(dir, name)); err == nil {
					content = b
					fmID, fmFamily := parseFrontmatter(b)
					if fmID != "" {
						id = fmID
					}
					family = fmFamily
				}

				// Local items overwrite in the customs map
				humanName := strings.ReplaceAll(id, "_", " ")

				r.AddCustom(&Item{
					ID:          id,
					Family:      family,
					Name:        humanName,
					Type:        itemType,
					Category:    "Custom",
					Description: "Custom " + itemType + " from " + dir,
					Builtin:     false,
					Content:     content,
				})
			}
		}
	}
	return nil
}
