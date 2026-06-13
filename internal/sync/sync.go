package sync

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Options holds configuration for the sync operation
type Options struct {
	Quiet  bool
	DryRun bool
}

// SyncToIDE syncs AI guidance files from .ai/ to IDE-specific directories.
func SyncToIDE(projectRoot string, opts Options) (bool, error) {
	aiRulesDir := filepath.Join(projectRoot, ".ai", "rules")
	aiCommandsDir := filepath.Join(projectRoot, ".ai", "commands")
	cursorRulesDir := filepath.Join(projectRoot, ".cursor", "rules")
	cursorCommandsDir := filepath.Join(projectRoot, ".cursor", "commands")
	claudeDir := filepath.Join(projectRoot, ".claude")

	if err := ensureDirectories(aiRulesDir, aiCommandsDir, cursorRulesDir, cursorCommandsDir, claudeDir, opts); err != nil {
		return false, fmt.Errorf("failed to ensure directories: %w", err)
	}

	baseMdcFiles, err := getSortedFiles(aiRulesDir, "mdc")
	if err != nil {
		return false, fmt.Errorf("failed to get base mdc files: %w", err)
	}

	createdRulesFiles := make(map[string]bool)
	hadBaseChanges, err := copyBaseMdcFiles(aiRulesDir, cursorRulesDir, projectRoot, baseMdcFiles, createdRulesFiles, opts)
	if err != nil {
		return false, err
	}

	hadCursorChanges, err := copyCursorMdcFiles(filepath.Join(aiRulesDir, "cursor"), cursorRulesDir, projectRoot, createdRulesFiles, opts)
	if err != nil {
		return false, err
	}

	createdCommandsFiles := make(map[string]bool)
	hadCommandsChanges, err := copyCommandFiles(aiCommandsDir, cursorCommandsDir, projectRoot, createdCommandsFiles, opts)
	if err != nil {
		return false, err
	}

	hadRemovalsRules, err := removeOldFiles(cursorRulesDir, createdRulesFiles, "mdc", projectRoot, opts)
	if err != nil {
		return false, err
	}

	hadRemovalsCommands, err := removeOldFiles(cursorCommandsDir, createdCommandsFiles, "md", projectRoot, opts)
	if err != nil {
		return false, err
	}

	hadClaudeChanges, err := generateClaudeFile(aiRulesDir, claudeDir, baseMdcFiles, projectRoot, opts)
	if err != nil {
		return false, err
	}

	hadAnyChanges := hadBaseChanges || hadCursorChanges || hadCommandsChanges || hadRemovalsRules || hadRemovalsCommands || hadClaudeChanges

	if !hadAnyChanges && !opts.Quiet {
		fmt.Println("No changes to make")
	}

	return hadAnyChanges, nil
}

func ensureDirectories(aiRulesDir, aiCommandsDir, cursorRulesDir, cursorCommandsDir, claudeDir string, opts Options) error {
	dirs := []string{
		aiRulesDir,
		filepath.Join(aiRulesDir, "claude"),
		filepath.Join(aiRulesDir, "cursor"),
		aiCommandsDir,
		cursorRulesDir,
		cursorCommandsDir,
		claudeDir,
	}

	for _, dir := range dirs {
		if !opts.DryRun {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

func getSortedFiles(dir, extension string) ([]string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "."+extension) {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, nil
}

func readFileContent(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(b), nil
}

func extractMdcContent(content string) string {
	// Remove YAML frontmatter between --- markers
	re := regexp.MustCompile(`(?m)^---\s*\n(?:.*?\n)*?---\s*\n`)
	return strings.TrimSpace(re.ReplaceAllString(content, ""))
}

func formatHeader(filename string) string {
	name := strings.TrimSuffix(filename, ".mdc")
	name = strings.TrimSuffix(name, ".md")
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.Title(name)
	return fmt.Sprintf("# %s\n\n", name)
}

func copyFileWithLog(source, dest, projectRoot string, opts Options) (bool, error) {
	sourceContent, err := os.ReadFile(source)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(dest); err == nil {
		destContent, err := os.ReadFile(dest)
		if err == nil && bytes.Equal(sourceContent, destContent) {
			return false, nil // No change needed
		}
	}

	if !opts.DryRun {
		if err := os.WriteFile(dest, sourceContent, 0644); err != nil {
			return false, err
		}
	}

	if !opts.Quiet {
		relSource, _ := filepath.Rel(projectRoot, source)
		relDest, _ := filepath.Rel(projectRoot, dest)
		if relSource == "" {
			relSource = source
		}
		if relDest == "" {
			relDest = dest
		}
		fmt.Printf("Copied: %s -> %s\n", relSource, relDest)
	}

	return true, nil
}

func copyBaseMdcFiles(aiRulesDir, cursorRulesDir, projectRoot string, baseMdcFiles []string, created map[string]bool, opts Options) (bool, error) {
	hadChanges := false
	for _, file := range baseMdcFiles {
		dest := filepath.Join(cursorRulesDir, filepath.Base(file))
		changed, err := copyFileWithLog(file, dest, projectRoot, opts)
		if err != nil {
			return false, err
		}
		if changed {
			hadChanges = true
		}
		created[dest] = true
	}
	return hadChanges, nil
}

func copyCursorMdcFiles(cursorSpecificDir, cursorRulesDir, projectRoot string, created map[string]bool, opts Options) (bool, error) {
	files, err := getSortedFiles(cursorSpecificDir, "mdc")
	if err != nil {
		return false, err
	}

	hadChanges := false
	for _, file := range files {
		dest := filepath.Join(cursorRulesDir, filepath.Base(file))
		changed, err := copyFileWithLog(file, dest, projectRoot, opts)
		if err != nil {
			return false, err
		}
		if changed {
			hadChanges = true
		}
		created[dest] = true
	}
	return hadChanges, nil
}

func copyCommandFiles(aiCommandsDir, cursorCommandsDir, projectRoot string, created map[string]bool, opts Options) (bool, error) {
	files, err := getSortedFiles(aiCommandsDir, "md")
	if err != nil {
		return false, err
	}

	hadChanges := false
	for _, file := range files {
		dest := filepath.Join(cursorCommandsDir, filepath.Base(file))
		changed, err := copyFileWithLog(file, dest, projectRoot, opts)
		if err != nil {
			return false, err
		}
		if changed {
			hadChanges = true
		}
		created[dest] = true
	}
	return hadChanges, nil
}

func removeOldFiles(targetDir string, created map[string]bool, extension, projectRoot string, opts Options) (bool, error) {
	files, err := getSortedFiles(targetDir, extension)
	if err != nil {
		return false, err
	}

	hadChanges := false
	for _, file := range files {
		if !created[file] {
			if !opts.DryRun {
				if err := os.Remove(file); err != nil {
					return false, err
				}
			}
			hadChanges = true
			if !opts.Quiet {
				rel, _ := filepath.Rel(projectRoot, file)
				if rel == "" {
					rel = file
				}
				fmt.Printf("Removed old file: %s\n", rel)
			}
		}
	}
	return hadChanges, nil
}

func concatenateMdcFilesForClaude(files []string) (string, error) {
	var result strings.Builder
	for _, file := range files {
		content, err := readFileContent(file)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(content) != "" {
			extracted := extractMdcContent(content)
			if extracted != "" {
				result.WriteString(formatHeader(filepath.Base(file)))
				result.WriteString(extracted)
				if !strings.HasSuffix(extracted, "\n") {
					result.WriteString("\n")
				}
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}

func concatenateMdFiles(files []string) (string, error) {
	var result strings.Builder
	for _, file := range files {
		content, err := readFileContent(file)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(content) != "" {
			result.WriteString(formatHeader(filepath.Base(file)))
			result.WriteString(content)
			if !strings.HasSuffix(content, "\n") {
				result.WriteString("\n")
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}

func generateClaudeFile(aiRulesDir, claudeDir string, baseMdcFiles []string, projectRoot string, opts Options) (bool, error) {
	baseContent, err := concatenateMdcFilesForClaude(baseMdcFiles)
	if err != nil {
		return false, err
	}

	claudeSpecificDir := filepath.Join(aiRulesDir, "claude")
	claudeMdFiles, err := getSortedFiles(claudeSpecificDir, "md")
	if err != nil {
		return false, err
	}

	claudeContent, err := concatenateMdFiles(claudeMdFiles)
	if err != nil {
		return false, err
	}

	newContent := baseContent + claudeContent
	if newContent == "" {
		return false, nil // Avoid writing an empty file if no contents
	}

	claudeOutput := filepath.Join(claudeDir, "CLAUDE.md")

	if _, err := os.Stat(claudeOutput); err == nil {
		existingContent, err := os.ReadFile(claudeOutput)
		if err == nil && string(existingContent) == newContent {
			return false, nil // No change needed
		}
	}

	if !opts.DryRun {
		if err := os.WriteFile(claudeOutput, []byte(newContent), 0644); err != nil {
			return false, err
		}
	}

	if !opts.Quiet {
		rel, _ := filepath.Rel(projectRoot, claudeOutput)
		if rel == "" {
			rel = claudeOutput
		}
		fmt.Printf("Generated: %s\n", rel)
	}

	return true, nil
}
