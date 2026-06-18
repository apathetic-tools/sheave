package sync

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/apathetic-tools/sheave/internal/registry"
)

func TestSyncToIDE(t *testing.T) {
	// Setup a temporary project directory
	tmpDir, err := os.MkdirTemp("", "sheave-sync-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	dirs := []string{filepath.Join(tmpDir, ".ai", "rules", "cursor"), filepath.Join(tmpDir, ".ai", "skills", "deep")}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}

	// 1. Base MDC file (flat)
	baseMdcPath := filepath.Join(tmpDir, ".ai", "rules", "base.mdc")
	baseMdcContent := `---
description: This is a base rule
---
Base rule content.
`
	if err := os.WriteFile(baseMdcPath, []byte(baseMdcContent), 0644); err != nil {
		t.Fatal(err)
	}

	// 2. Cursor specific MDC file (nested)
	cursorMdcPath := filepath.Join(tmpDir, ".ai", "rules", "cursor", "cursor_specific.mdc")
	cursorMdcContent := `---
description: Cursor specific
---
Cursor specific content.
`
	if err := os.WriteFile(cursorMdcPath, []byte(cursorMdcContent), 0644); err != nil {
		t.Fatal(err)
	}

	// 3. Command file (nested)
	commandMdPath := filepath.Join(tmpDir, ".ai", "skills", "deep", "my_command.md")
	commandMdContent := "---\nsheave-name: deep_command_renamed\n---\nCommand content.\n"
	if err := os.WriteFile(commandMdPath, []byte(commandMdContent), 0644); err != nil {
		t.Fatal(err)
	}

	// 4. Obsolete file in .cursor/rules/
	cursorRulesDir := filepath.Join(tmpDir, ".cursor", "rules")
	if err := os.MkdirAll(cursorRulesDir, 0755); err != nil {
		t.Fatal(err)
	}
	obsoleteFile := filepath.Join(cursorRulesDir, "obsolete.mdc")
	if err := os.WriteFile(obsoleteFile, []byte("obsolete"), 0644); err != nil {
		t.Fatal(err)
	}

	opts := Options{Quiet: true, DryRun: false}
	changed, err := SyncToIDE(tmpDir, opts)
	if err != nil {
		t.Fatalf("SyncToIDE failed: %v", err)
	}
	if !changed {
		t.Errorf("Expected changes to be true, got false")
	}

	// Verify Cursor Rules
	if _, err := os.Stat(filepath.Join(cursorRulesDir, "base.mdc")); os.IsNotExist(err) {
		t.Errorf("base.mdc was not copied to .cursor/rules")
	}
	if _, err := os.Stat(filepath.Join(cursorRulesDir, "cursor_specific.mdc")); os.IsNotExist(err) {
		t.Errorf("cursor_specific.mdc was not copied to .cursor/rules")
	}
	if _, err := os.Stat(obsoleteFile); !os.IsNotExist(err) {
		t.Errorf("obsolete.mdc was not removed from .cursor/rules")
	}

	// Verify Cursor Commands
	cursorCommandsDir := filepath.Join(tmpDir, ".cursor", "skills")
	if _, err := os.Stat(filepath.Join(cursorCommandsDir, "deep_command_renamed.md")); os.IsNotExist(err) {
		t.Errorf("deep_command_renamed.md was not copied to .cursor/skills/")
	}

	// Verify Claude output (Modular deployment)
	claudeRulesDir := filepath.Join(tmpDir, ".claude", "rules")
	if _, err := os.Stat(filepath.Join(claudeRulesDir, "base.md")); os.IsNotExist(err) {
		t.Errorf("base.md was not copied to .claude/rules")
	}
	if _, err := os.Stat(filepath.Join(claudeRulesDir, "cursor_specific.md")); os.IsNotExist(err) {
		t.Errorf("cursor_specific.md was not copied to .claude/rules")
	}

	claudeCommandsDir := filepath.Join(tmpDir, ".claude", "skills")
	if _, err := os.Stat(filepath.Join(claudeCommandsDir, "deep_command_renamed.md")); os.IsNotExist(err) {
		t.Errorf("deep_command_renamed.md was not copied to .claude/skills")
	}

	// settings.json should be generated with at least the $schema key
	settingsFile := filepath.Join(tmpDir, ".claude", "settings.json")
	if _, err := os.Stat(settingsFile); os.IsNotExist(err) {
		t.Errorf("settings.json was not generated in .claude")
	} else {
		content, err := os.ReadFile(settingsFile)
		if err != nil {
			t.Errorf("failed to read settings.json: %v", err)
		}
		if !strings.Contains(string(content), `"$schema"`) {
			t.Errorf("settings.json does not contain $schema. Content: %s", string(content))
		}
	}

	// We no longer generate empty catchall files if they don't exist
	claudeFile := filepath.Join(tmpDir, ".claude", "CLAUDE.md")
	if _, err := os.Stat(claudeFile); !os.IsNotExist(err) {
		t.Errorf("CLAUDE.md was incorrectly generated in .claude")
	}

	// Test a second run (should be no changes)
	changed, err = SyncToIDE(tmpDir, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFrontmatterOverrides(t *testing.T) {
	// Setup a temporary project directory
	tmpDir, err := os.MkdirTemp("", "sheave-frontmatter-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	if err := os.MkdirAll(filepath.Join(tmpDir, ".ai", "skills", "deep"), 0755); err != nil {
		t.Fatal(err)
	}

	// Command file with overrides
	commandMdPath := filepath.Join(tmpDir, ".ai", "skills", "deep", "my_command.md")
	commandMdContent := "---\nsheave-name: overridden_name\nsheave-family: //docs/skills\nsheave-id: overridden_id\n---\nContent.\n"
	if err := os.WriteFile(commandMdPath, []byte(commandMdContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Use a single known provider so sheave-family overrides can be tested without
	// provider layouts needing to be config-defined (adapters are now hardcoded).
	configContent := `
active_providers = ["claude"]
`
	if err := os.WriteFile(filepath.Join(tmpDir, ".ai", ".sheave.toml"), []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	opts := Options{Quiet: true, DryRun: false}
	_, err = SyncToIDE(tmpDir, opts)
	if err != nil {
		t.Fatalf("SyncToIDE failed: %v", err)
	}

	// Verify the file was written to the absolute path correctly due to sheave-family
	// and named correctly due to sheave-name.
	// It should be at /tmp/.../docs/skills/overridden_name.md (ignoring the .test_provider targetDir)
	expectedPath := filepath.Join(tmpDir, "docs", "skills", "overridden_name.md")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("File with overridden family and name was not created at expected path: %s", expectedPath)
	}

	// Verify sheave-id correctly overrides the item ID in the registry
	reg := registry.NewRegistry()
	if err := reg.DiscoverCustomItems(tmpDir); err != nil {
		t.Fatal(err)
	}
	item, ok := reg.Get("//docs/skills/overridden_id")
	if !ok {
		t.Errorf("Item not found in registry by its overridden sheave-family/sheave-id key")
	}
	if item != nil && item.BaseName != "overridden_name" {
		t.Errorf("Expected baseName to be overridden_name, got %s", item.BaseName)
	}
	if item != nil && item.Family != "//docs/skills" {
		t.Errorf("Expected family to be //docs/skills, got %s", item.Family)
	}
}

func TestSecurityJailing(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "sheave-security-test")
	defer os.RemoveAll(tmpDir)
	os.MkdirAll(filepath.Join(tmpDir, ".ai", "rules"), 0755)

	maliciousPath := filepath.Join(tmpDir, ".ai", "rules", "evil.md")
	maliciousContent := "---\nsheave-family: ../../../../../etc\n---\nEvil content.\n"
	os.WriteFile(maliciousPath, []byte(maliciousContent), 0644)

	opts := Options{Quiet: true, DryRun: false}
	_, err := SyncToIDE(tmpDir, opts)
	if err == nil {
		t.Errorf("Expected security violation error, got nil")
	}
}
