package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
		if provider.Rules.Spread == "dir" {
			rulesDir := filepath.Join(projectRoot, provider.TargetDir, provider.Rules.Path)
			if !opts.DryRun {
				_ = os.MkdirAll(rulesDir, 0755)
			}
			var changed bool
			var err error

			// Use the extension specified in config, ensuring it starts with a dot
			ext := provider.Rules.Ext
			if ext != "" && !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}

			createdRules, changed, err = writeItems(rules, rulesDir, ext, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			changed, err = removeOldFiles(rulesDir, createdRules, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed
		} else {
			// Spread file is not implemented yet for rules
			fmt.Printf("Stub: rules spread=file not implemented yet\n")
		}
	} else {
		unhandledRules = rules
	}

	// 2. Process "skills" component (maps to commands from registry)
	if provider.Skills != nil {
		if provider.Skills.Spread == "dir" {
			commandsDir := filepath.Join(projectRoot, provider.TargetDir, provider.Skills.Path)
			if !opts.DryRun {
				_ = os.MkdirAll(commandsDir, 0755)
			}
			var changed bool
			var err error

			ext := provider.Skills.Ext
			if ext != "" && !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}

			createdCommands, changed, err = writeItems(commands, commandsDir, ext, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed

			changed, err = removeOldFiles(commandsDir, createdCommands, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed
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
			stubContent := []byte(fmt.Sprintf("{\n  \"//\": \"Stub for %s\"\n}\n", name))
			if comp.Type == "dot" {
				stubContent = []byte(fmt.Sprintf("# Stub for %s\n", name))
			}
			changed, err := writeIfChanged(path, stubContent, projectRoot, opts)
			if err != nil {
				return false, err
			}
			hadChanges = hadChanges || changed
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
			// Touch empty main file if it didn't exist
			if _, err := os.Stat(outputFile); os.IsNotExist(err) {
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
