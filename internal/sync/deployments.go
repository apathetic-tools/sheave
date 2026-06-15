package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/registry"
)

func deployDataDriven(provider config.ProviderConfig, rules, commands []*registry.Item, projectRoot string, opts Options) (bool, error) {
	hadChanges := false

	var unhandledRules, unhandledCommands []*registry.Item
	createdRules := make(map[string]bool)
	createdCommands := make(map[string]bool)

	// 1. Process "rules" component
	if provider.Rules != nil {
		if provider.Rules.Spread == "dir" || provider.Rules.Spread == "subdir" {
			rulesDir := filepath.Join(projectRoot, provider.TargetDir, provider.Rules.Path)
			var changed bool
			var err error

			// Use the extension specified in config, ensuring it starts with a dot
			ext := provider.Rules.Ext
			if ext != "" && !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}

			var generatedRules []*registry.Item
			for _, item := range rules {
				newItem := *item
				content, err := generateItemContent(item, provider.Rules)
				if err != nil {
					return false, err
				}
				newItem.Content = content
				generatedRules = append(generatedRules, &newItem)
			}

			createdRules, changed, err = writeItems(generatedRules, rulesDir, ext, projectRoot, provider.Rules.Spread, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			changed, err = removeOldFiles(rulesDir, createdRules, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			// Clean up the directory if it's completely empty (e.g. all rules were removed)
			if !opts.DryRun {
				_ = os.Remove(rulesDir)
			}
		} else {
			// Spread file is not implemented yet for rules
			fmt.Printf("Stub: rules spread=file not implemented yet\n")
		}
	} else {
		unhandledRules = rules
	}

	// 2. Process "skills" component (maps to commands from registry)
	if provider.Skills != nil {
		if provider.Skills.Spread == "dir" || provider.Skills.Spread == "subdir" {
			commandsDir := filepath.Join(projectRoot, provider.TargetDir, provider.Skills.Path)
			var changed bool
			var err error

			ext := provider.Skills.Ext
			if ext != "" && !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}

			var generatedCommands []*registry.Item
			for _, item := range commands {
				newItem := *item
				content, err := generateItemContent(item, provider.Skills)
				if err != nil {
					return false, err
				}
				newItem.Content = content
				generatedCommands = append(generatedCommands, &newItem)
			}

			createdCommands, changed, err = writeItems(generatedCommands, commandsDir, ext, projectRoot, provider.Skills.Spread, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			changed, err = removeOldFiles(commandsDir, createdCommands, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			// Clean up the directory if it's completely empty (e.g. all skills were removed)
			if !opts.DryRun {
				_ = os.Remove(commandsDir)
			}
		} else {
			fmt.Printf("Stub: skills spread=file not implemented yet\n")
		}
	} else {
		unhandledCommands = commands
	}

	// 3. Stub other simple components
	simpleComponents := map[string]*config.ComponentConfig{
		"settings":    provider.Settings,
		"hooks":       provider.Hooks,
		"mcp":         provider.MCP,
		"environment": provider.Environment,
		"ide":         provider.IDE,
		"ignore":      provider.Ignore,
	}

	for name, comp := range simpleComponents {
		if comp != nil && comp.Spread == "file" {
			path := filepath.Join(projectRoot, provider.TargetDir, comp.Path)
			// If it's just a stub, don't write it if it doesn't exist.
			// Also don't overwrite it if it DOES exist (to preserve user's content).
			// We only want to generate these components when we implement real generators.
			if _, err := os.Stat(path); err == nil {
				// It exists, do nothing so we don't overwrite user content with {}
			} else if os.IsNotExist(err) {
				// It doesn't exist, and user asked not to create empty ones
				continue
			}
		} else if comp != nil {
			fmt.Printf("Stub: %s spread=%s not implemented yet\n", name, comp.Spread)
		}
	}

	// 4. Process "main" component (catchall/memory)
	if provider.Main != nil {
		hasCatchall := hasFlavour(provider.Main, "catchall")
		hasMemory := hasFlavour(provider.Main, "memory")

		var result strings.Builder

		if hasCatchall {
			for _, item := range unhandledRules {
				body := extractContentBody(item.Content)
				if body != "" {
					result.WriteString(fmt.Sprintf("# %s\n\n%s\n\n", item.Name, body))
				}
			}
			for _, item := range unhandledCommands {
				body := extractContentBody(item.Content)
				if body != "" {
					result.WriteString(fmt.Sprintf("# %s\n\n%s\n\n", item.Name, body))
				}
			}
		}

		if hasMemory {
			result.WriteString("# Antigravity IDE Memory Checklist\n\n")
			result.WriteString("1. Record the contents of this file into your memory (KI).\n")
			result.WriteString("2. Record the contents of each of the files mentioned below to your memory, with the appropriate context and date provided. If the context or date changes, re-read and re-commit it to memory.\n\n")
			result.WriteString("## Index\n")

			rulePaths := sortedKeys(createdRules)
			for _, path := range rulePaths {
				stat, err := os.Stat(path)
				if err == nil {
					rel, _ := filepath.Rel(filepath.Join(projectRoot, provider.TargetDir), path)
					result.WriteString(fmt.Sprintf("- [%s] **Rule**: `%s`\n", stat.ModTime().Format("2006-01-02T15:04:05Z"), rel))
				}
			}

			commandPaths := sortedKeys(createdCommands)
			for _, path := range commandPaths {
				stat, err := os.Stat(path)
				if err == nil {
					rel, _ := filepath.Rel(filepath.Join(projectRoot, provider.TargetDir), path)
					result.WriteString(fmt.Sprintf("- [%s] **Command**: `%s`\n", stat.ModTime().Format("2006-01-02T15:04:05Z"), rel))
				}
			}
		}

		outputFile := filepath.Join(projectRoot, provider.TargetDir, provider.Main.Path)

		if result.Len() > 0 {
			changed, err := writeIfChanged(outputFile, []byte(result.String()), projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed
		} else if hasCatchall && !opts.DryRun {
			if _, err := os.Stat(outputFile); err == nil {
				changed, err := writeIfChanged(outputFile, []byte(""), projectRoot, opts)
				if err != nil {
					return false, err
				}
				hadChanges = hadChanges || changed
			}
		}
	}

	return hadChanges, nil
}

func hasFlavour(comp *config.ComponentConfig, val string) bool {
	if comp == nil {
		return false
	}

	check := func(v any) bool {
		switch f := v.(type) {
		case string:
			return f == val
		case []any:
			for _, item := range f {
				if str, ok := item.(string); ok && str == val {
					return true
				}
			}
		case []string:
			for _, str := range f {
				if str == val {
					return true
				}
			}
		}
		return false
	}

	return check(comp.Flavour) || check(comp.Flavor)
}

func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func generateItemContent(item *registry.Item, comp *config.ComponentConfig) ([]byte, error) {
	if comp == nil {
		return item.Content, nil
	}

	body := extractContentBody(item.Content)

	existingFM := make(map[string]any)
	str := string(item.Content)
	hasFrontmatter := false
	if strings.HasPrefix(str, "---\n") || strings.HasPrefix(str, "---\r\n") {
		start := strings.Index(str, "\n") + 1
		end := strings.Index(str[start:], "\n---")
		if end != -1 {
			yamlStr := str[start : start+end]
			_ = yaml.Unmarshal([]byte(yamlStr), &existingFM)
			hasFrontmatter = true
		}
	}

	newFM := make(map[string]any)

	if len(comp.Frontmatter) > 0 {
		for _, field := range comp.Frontmatter {
			var prop *config.FieldProp
			switch field {
			case "name":
				prop = comp.Name
			case "description":
				prop = comp.Description
			case "invocable":
				prop = comp.Invocable
			case "metadata":
				prop = comp.Metadata
			}

			val, exists := existingFM[field]
			if !exists {
				if field == "name" {
					val = item.Name
				} else if field == "description" {
					val = item.Description
				} else if prop != nil && prop.Name != "" {
					val = prop.Name
				}
			}

			if val != nil && val != "" {
				newFM[field] = val
			}
		}
	} else {
		if !hasFrontmatter {
			return item.Content, nil
		}
		for k, v := range existingFM {
			if !strings.HasPrefix(k, "sheave-") {
				newFM[k] = v
			}
		}
	}

	if len(newFM) == 0 {
		return []byte(body), nil
	}

	fmBytes, err := yaml.Marshal(newFM)
	if err != nil {
		return nil, err
	}

	var result strings.Builder
	result.WriteString("---\n")
	result.WriteString(string(fmBytes))
	result.WriteString("---\n")
	result.WriteString(body)

	return []byte(result.String()), nil
}
