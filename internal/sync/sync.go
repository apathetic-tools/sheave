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
	activeCommands := reg.Resolve("Command", cfg.Commands.Include, cfg.Commands.Exclude)

	hadChanges := false

	for name, provider := range cfg.Providers {
		var changed bool
		var err error

		switch provider.DeploymentMethod {
		case "flat-combine":
			changed, err = deployFlatCombine(provider, activeRules, activeCommands, projectRoot, opts)
		case "folder-split-command-rules":
			changed, err = deployFolderSplit(provider, activeRules, activeCommands, projectRoot, opts)
		case "memory":
			changed, err = deployMemory(provider, activeRules, activeCommands, projectRoot, opts)
		default:
			if !opts.Quiet {
				fmt.Printf("Warning: unknown deployment_method '%s' for provider '%s'\n", provider.DeploymentMethod, name)
			}
		}

		if err != nil {
			return false, fmt.Errorf("failed to sync provider %s: %w", name, err)
		}
		hadChanges = hadChanges || changed
	}

	if !hadChanges && !opts.Quiet {
		fmt.Println("No changes to make")
	}
	return hadChanges, nil
}

func writeItems(items []*registry.Item, targetDir string, extension string, projectRoot string, opts Options) (map[string]bool, bool, error) {
	created := make(map[string]bool)
	hadChanges := false

	for _, item := range items {
		filename := item.ID
		if item.Family != "" {
			filename = item.Family + "_" + item.ID
		}
		filename += extension
		dest := filepath.Join(targetDir, filename)
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
