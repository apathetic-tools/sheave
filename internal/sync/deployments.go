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

// deployFlatCombine appends all rules and commands into a single file.
func deployFlatCombine(provider config.ProviderConfig, rules, commands []*registry.Item, projectRoot string, opts Options) (bool, error) {
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

	if result.Len() == 0 {
		return false, nil
	}

	outputFile := filepath.Join(projectRoot, provider.TargetDir, provider.Filename)
	return writeIfChanged(outputFile, []byte(result.String()), projectRoot, opts)
}

// deployFolderSplit splits items into target_dir/rules and target_dir/commands.
func deployFolderSplit(provider config.ProviderConfig, rules, commands []*registry.Item, projectRoot string, opts Options) (bool, error) {
	rulesDir := filepath.Join(projectRoot, provider.TargetDir, "rules")
	commandsDir := filepath.Join(projectRoot, provider.TargetDir, "commands")

	if !opts.DryRun {
		_ = os.MkdirAll(rulesDir, 0755)
		_ = os.MkdirAll(commandsDir, 0755)
	}

	hadChanges := false

	createdRules, changed, err := writeItems(rules, rulesDir, ".mdc", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	createdCommands, changed, err := writeItems(commands, commandsDir, ".md", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(rulesDir, createdRules, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(commandsDir, createdCommands, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	return hadChanges, nil
}

// deployMemory generates subfiles and an index referencing them with ModTime timestamps.
func deployMemory(provider config.ProviderConfig, rules, commands []*registry.Item, projectRoot string, opts Options) (bool, error) {
	rulesDir := filepath.Join(projectRoot, provider.TargetDir, "rules")
	commandsDir := filepath.Join(projectRoot, provider.TargetDir, "commands")

	if !opts.DryRun {
		_ = os.MkdirAll(rulesDir, 0755)
		_ = os.MkdirAll(commandsDir, 0755)
	}

	hadChanges := false

	createdRules, changed, err := writeItems(rules, rulesDir, ".md", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	createdCommands, changed, err := writeItems(commands, commandsDir, ".md", projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(rulesDir, createdRules, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	changed, err = removeOldFiles(commandsDir, createdCommands, projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	// Generate Index File
	var result strings.Builder
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

	outputFile := filepath.Join(projectRoot, provider.TargetDir, provider.Filename)
	changed, err = writeIfChanged(outputFile, []byte(result.String()), projectRoot, opts)
	if err != nil {
		return false, err
	}
	hadChanges = hadChanges || changed

	return hadChanges, nil
}

func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
