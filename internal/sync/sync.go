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

	cursorRulesDir := filepath.Join(projectRoot, ".cursor", "rules")
	cursorCommandsDir := filepath.Join(projectRoot, ".cursor", "commands")
	claudeDir := filepath.Join(projectRoot, ".claude")

	if !opts.DryRun {
		_ = os.MkdirAll(cursorRulesDir, 0755)
		_ = os.MkdirAll(cursorCommandsDir, 0755)
		_ = os.MkdirAll(claudeDir, 0755)
	}

	hadChanges := false

	createdRules, changed, err := writeItems(activeRules, cursorRulesDir, ".mdc", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	createdCommands, changed, err := writeItems(activeCommands, cursorCommandsDir, ".md", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(cursorRulesDir, createdRules, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(cursorCommandsDir, createdCommands, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = generateClaudeFile(claudeDir, activeRules, activeCommands, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

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

func generateClaudeFile(claudeDir string, rules, commands []*registry.Item, projectRoot string, opts Options) (bool, error) {
	var result strings.Builder

	for _, item := range rules {
		body := extractContentBody(item.Content)
		if body != "" {
			result.WriteString(fmt.Sprintf("# %s\n\n%s\n\n", item.Name, body))
		}
	}

	for _, item := range commands {
		body := extractContentBody(item.Content)
		if body != "" {
			result.WriteString(fmt.Sprintf("# %s\n\n%s\n\n", item.Name, body))
		}
	}

	outputFile := filepath.Join(claudeDir, "CLAUDE.md")

	if result.Len() == 0 {
		return false, nil
	}

	return writeIfChanged(outputFile, []byte(result.String()), projectRoot, opts)
}
