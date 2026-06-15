package sync

import (
	"os"
	"path/filepath"
	"testing"
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
	commandMdContent := "Command content.\n"
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
	if _, err := os.Stat(filepath.Join(cursorRulesDir, "cursor_cursor_specific.mdc")); os.IsNotExist(err) {
		t.Errorf("cursor_cursor_specific.mdc was not copied to .cursor/rules")
	}
	if _, err := os.Stat(obsoleteFile); !os.IsNotExist(err) {
		t.Errorf("obsolete.mdc was not removed from .cursor/rules")
	}

	// Verify Cursor Commands
	cursorCommandsDir := filepath.Join(tmpDir, ".cursor", "skills")
	if _, err := os.Stat(filepath.Join(cursorCommandsDir, "deep_my_command.md")); os.IsNotExist(err) {
		t.Errorf("deep_my_command.md was not copied to .cursor/skills/")
	}

	// Verify Claude output (Modular deployment)
	claudeRulesDir := filepath.Join(tmpDir, ".claude", "rules")
	if _, err := os.Stat(filepath.Join(claudeRulesDir, "base.md")); os.IsNotExist(err) {
		t.Errorf("base.md was not copied to .claude/rules")
	}
	if _, err := os.Stat(filepath.Join(claudeRulesDir, "cursor_cursor_specific.md")); os.IsNotExist(err) {
		t.Errorf("cursor_cursor_specific.md was not copied to .claude/rules")
	}

	claudeCommandsDir := filepath.Join(tmpDir, ".claude", "skills")
	if _, err := os.Stat(filepath.Join(claudeCommandsDir, "deep_my_command.md")); os.IsNotExist(err) {
		t.Errorf("deep_my_command.md was not copied to .claude/skills")
	}

	// We no longer generate empty stub files, so settings.json should not exist
	settingsFile := filepath.Join(tmpDir, ".claude", "settings.json")
	if _, err := os.Stat(settingsFile); !os.IsNotExist(err) {
		t.Errorf("settings.json was incorrectly generated in .claude")
	}

	claudeFile := filepath.Join(tmpDir, ".claude", "CLAUDE.md")
	claudeContent, err := os.ReadFile(claudeFile)
	if err != nil {
		t.Fatalf("Failed to read CLAUDE.md: %v", err)
	}

	if len(claudeContent) > 0 {
		t.Errorf("Expected CLAUDE.md to be empty, got %d bytes", len(claudeContent))
	}

	// Test a second run (should be no changes)
	changed, err = SyncToIDE(tmpDir, opts)
	if err != nil {
		t.Fatalf("SyncToIDE failed on second run: %v", err)
	}
	if changed {
		t.Errorf("Expected changes to be false on second run, got true")
	}
}
