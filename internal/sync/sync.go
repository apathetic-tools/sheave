package sync

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/providers"
	"github.com/apathetic-tools/sheave/internal/registry"
)

type Options struct {
	Quiet  bool
	DryRun bool
}

func SyncToIDE(projectRoot string, opts Options) (bool, error) {
	reg := registry.NewRegistry()
	if err := reg.DiscoverCustomItems(projectRoot); err != nil {
		return false, fmt.Errorf("failed to discover custom items: %w", err)
	}

	path := config.GetConfigPath(projectRoot)
	cfg, err := config.Load(path)
	if err != nil {
		return false, fmt.Errorf("failed to load config: %w", err)
	}

	activeRules := reg.Resolve("Rule", cfg.Rules.Include, cfg.Rules.Exclude)
	activeCommands := reg.Resolve("Skill", cfg.Skills.Include, cfg.Skills.Exclude)
	activeSettings := reg.Resolve("Setting", cfg.Settings.Include, cfg.Settings.Exclude)

	hadChanges := false

	for _, name := range cfg.ActiveProviders {
		adapter := providers.GetAdapter(name)
		layout := adapter.GetLayout()

		var changed bool
		var err error

		changed, err = deployDataDriven(name, layout, activeRules, activeCommands, activeSettings, projectRoot, opts)

		if err != nil {
			return false, fmt.Errorf("failed to sync provider %s: %w", name, err)
		}
		hadChanges = hadChanges || changed
	}

	if !hadChanges && !opts.Quiet {
		fmt.Println("No changes detected.")
	}
	return hadChanges, nil
}

func writeItems(items []*registry.Item, targetDir string, extension string, projectRoot string, spread string, opts Options) (map[string]bool, bool, error) {
	created := make(map[string]bool)
	hadChanges := false

	idCounts := make(map[string]int)
	if spread == "dir" {
		for _, item := range items {
			idCounts[item.BaseName]++
		}
	}

	for _, item := range items {
		filename := item.BaseName + extension
		var dest string

		if strings.HasPrefix(item.Family, "//") || strings.HasPrefix(item.Family, "/") {
			familyPath := strings.TrimPrefix(item.Family, "//")
			familyPath = strings.TrimPrefix(familyPath, "/")
			dest = filepath.Join(projectRoot, familyPath, filename)
		} else if item.IsFamilyOverridden {
			dest = filepath.Join(targetDir, item.Family, filename)
		} else {
			if item.Family != "" {
				if spread == "subdir" {
					dest = filepath.Join(targetDir, item.Family, filename)
				} else if spread == "dir" {
					if idCounts[item.BaseName] > 1 {
						parts := strings.Split(item.Family, "/")
						lastDir := parts[len(parts)-1]
						filename = lastDir + "-" + item.BaseName + extension
					}
					dest = filepath.Join(targetDir, filename)
				} else {
					dest = filepath.Join(targetDir, filename)
				}
			} else {
				dest = filepath.Join(targetDir, filename)
			}
		}

		cleanDest := filepath.Clean(dest)
		cleanRoot := filepath.Clean(projectRoot)

		if !strings.HasPrefix(cleanDest, cleanRoot+string(filepath.Separator)) && cleanDest != cleanRoot {
			return nil, false, fmt.Errorf("security violation: path %s escapes project root %s", cleanDest, cleanRoot)
		}

		dest = cleanDest
		created[dest] = true

		changed, err := writeIfChanged(dest, item.Content, projectRoot, opts)
		if err != nil {
			return nil, false, err
		}
		if changed {
			hadChanges = true
		}
	}
	return created, hadChanges, nil
}

func writeIfChanged(dest string, newContent []byte, projectRoot string, opts Options) (bool, error) {
	if _, err := os.Stat(dest); err == nil {
		existingContent, err := os.ReadFile(dest)
		if err == nil && bytes.Equal(existingContent, newContent) {
			return false, nil
		}
	}

	if !opts.DryRun {
		if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
			return false, err
		}
		if err := os.WriteFile(dest, newContent, 0644); err != nil {
			return false, err
		}
	}

	if !opts.Quiet {
		relDest, _ := filepath.Rel(projectRoot, dest)
		if relDest == "" {
			relDest = dest
		}
		fmt.Printf("Generated: %s\n", relDest)
	}

	return true, nil
}

func removeOldFiles(targetDir string, created map[string]bool, projectRoot string, opts Options) (bool, error) {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return false, nil
	}

	hadChanges := false
	err := filepath.WalkDir(targetDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !created[path] {
			if !opts.DryRun {
				_ = os.Remove(path)
			}
			hadChanges = true
			if !opts.Quiet {
				rel, _ := filepath.Rel(projectRoot, path)
				if rel == "" {
					rel = path
				}
				fmt.Printf("Removed old file: %s\n", rel)
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return hadChanges, nil
}

func extractContentBody(content []byte) string {
	str := string(content)
	re := regexp.MustCompile(`(?m)^---\s*\n(?:.*?\n)*?---\s*\n`)
	return strings.TrimSpace(re.ReplaceAllString(str, ""))
}
