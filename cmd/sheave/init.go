package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: "utility",
	Use:     "init",
	Short:   "Interactively initialize Sheave configuration and scaffold the .ai directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		autoAccept, _ := cmd.Flags().GetBool("yes")

		askYesNo := func(prompt string, defaultYes bool) bool {
			if autoAccept {
				return defaultYes
			}
			def := "[Y/n]"
			if !defaultYes {
				def = "[y/N]"
			}
			fmt.Printf("%s %s: ", prompt, def)
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return defaultYes
			}
			response = strings.ToLower(strings.TrimSpace(response))
			if response == "" {
				return defaultYes
			}
			return response == "y" || response == "yes"
		}

		// 1. Scaffold Prompt
		scaffold := askYesNo("Do you want to scaffold the .ai directory structure to provide your own rules?", true)
		if scaffold {
			dirs := []string{"rules", "commands", "templates", "workflows"}
			for _, dir := range dirs {
				dirPath := filepath.Join(cwd, ".ai", dir)
				if err := os.MkdirAll(dirPath, 0755); err != nil {
					return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
				}
				gitkeepPath := filepath.Join(dirPath, ".gitkeep")
				if _, err := os.Stat(gitkeepPath); os.IsNotExist(err) {
					_ = os.WriteFile(gitkeepPath, []byte(""), 0644)
				}
			}
			fmt.Println("Scaffolded .ai/ directory structure.")
		}

		// 2. Environment Detection
		hasGo := false
		hasNode := false
		hasPython := false
		hasGithub := false

		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			hasGo = true
		}
		if _, err := os.Stat(filepath.Join(cwd, "package.json")); err == nil {
			hasNode = true
		}
		if _, err := os.Stat(filepath.Join(cwd, "requirements.txt")); err == nil {
			hasPython = true
		} else if _, err := os.Stat(filepath.Join(cwd, "pyproject.toml")); err == nil {
			hasPython = true
		}
		if _, err := os.Stat(filepath.Join(cwd, ".github", "workflows")); err == nil {
			hasGithub = true
		}

		// 3. Ask about Builtins
		includes := []string{"~*"} // Always include userland by default

		if hasGo {
			if askYesNo("\nI detected a Go project. Enable Go built-in rules?", true) {
				includes = append(includes, "#golang/*")
			}
		}
		if hasNode {
			if askYesNo("\nI detected a Node.js project. Enable Node built-in rules?", true) {
				includes = append(includes, "#node/*")
			}
		}
		if hasPython {
			if askYesNo("\nI detected a Python project. Enable Python built-in rules?", true) {
				includes = append(includes, "#python/*")
			}
		}
		if hasGithub {
			if askYesNo("\nI detected GitHub Actions. Enable CI/CD built-in rules?", true) {
				includes = append(includes, "#ci/*")
			}
		}

		if askYesNo("\nEnable general code quality built-in rules?", true) {
			includes = append(includes, "#quality/*")
		}

		cfg := &config.Config{
			Rules:     config.Selection{Include: includes},
			Skills:    config.Selection{Include: includes},
			Templates: config.Selection{Include: includes},
			Workflows: config.Selection{Include: includes},
		}

		path := config.GetConfigPath(cwd)
		// If scaffolding created .ai directory, config might need to be saved inside it
		if scaffold {
			path = filepath.Join(cwd, ".ai", ".sheave.toml")
		}

		if err := cfg.Save(path); err != nil {
			return err
		}

		fmt.Printf("\nSuccessfully configured Sheave in %s\n", path)

		synced := false
		if len(includes) > 1 {
			if askYesNo("\nDo you want to run 'sheave sync' now to apply these rules to your IDE?", true) {
				fmt.Println() // Add a blank line for readability before sync output
				err := syncCmd.RunE(cmd, args)
				if err == nil {
					synced = true
				} else {
					return err
				}
			}
		}

		if !synced {
			fmt.Printf("\nTo get started, add rules to %s or create custom rules in your .ai/ directory.\n", path)
			fmt.Println("When you are ready, run 'sheave sync' to apply them to your IDE.")
		}

		return nil
	},
}

func init() {
	initCmd.Flags().BoolP("yes", "y", false, "Automatically accept default prompts")
	rootCmd.AddCommand(initCmd)
}
