package project

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var IgnoreDirs = map[string]bool{
	".git":          true,
	".mypy_cache":   true,
	".ruff_cache":   true,
	".pytest_cache": true,
	".venv":         true,
	"__pycache__":   true,
	"node_modules":  true,
	"dist":          true,
	"build":         true,
}

func shouldIgnore(relPath string) bool {
	parts := strings.Split(relPath, string(os.PathSeparator))
	for _, part := range parts {
		if IgnoreDirs[part] {
			return true
		}
	}
	return false
}

func formatPath(relPath string, isDir bool) string {
	depth := strings.Count(relPath, string(os.PathSeparator))
	indent := strings.Repeat("  ", depth)

	if isDir {
		return fmt.Sprintf("%s📁 %s", indent, relPath)
	}
	return fmt.Sprintf("%s📄 %s", indent, relPath)
}

// GenerateSummary recursively walks the directory and prints a formatted project structure.
func GenerateSummary(rootDir string) error {
	absRoot, err := filepath.Abs(rootDir)
	if err != nil {
		return err
	}

	fmt.Printf("📦 Project structure under: %s\n", absRoot)

	var ignores []string
	for k := range IgnoreDirs {
		ignores = append(ignores, k)
	}
	sort.Strings(ignores)
	fmt.Printf("🧹 Ignoring: %s\n", strings.Join(ignores, ", "))
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Formatted tree:")

	var paths []string
	count := 0

	err = filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == absRoot {
			return nil
		}

		relPath, err := filepath.Rel(absRoot, path)
		if err != nil {
			return err
		}

		if shouldIgnore(relPath) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		paths = append(paths, relPath)
		return nil
	})

	if err != nil {
		return err
	}

	// WalkDir visits in lexical order, so it's already sorted.
	for _, relPath := range paths {
		info, err := os.Stat(filepath.Join(absRoot, relPath))
		if err != nil {
			continue
		}
		fmt.Println(formatPath(relPath, info.IsDir()))
		count++
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Printf("✅ Done. Printed %d visible entries.\n", count)

	return nil
}
