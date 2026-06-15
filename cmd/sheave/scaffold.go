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

var scaffoldCmd = &cobra.Command{
	GroupID: "utility",
	Use:     "scaffold",
	Short:   "Scaffold the .ai directory structure and configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		autoAccept, _ := cmd.Flags().GetBool("yes")

		// Create directories and .gitkeep files
		dirs := []string{"rules", "commands", "templates", "workflows"}
		for _, dir := range dirs {
			dirPath := filepath.Join(cwd, ".ai", dir)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
			}

			gitkeepPath := filepath.Join(dirPath, ".gitkeep")
			if _, err := os.Stat(gitkeepPath); os.IsNotExist(err) {
				if err := os.WriteFile(gitkeepPath, []byte(""), 0644); err != nil {
					return fmt.Errorf("failed to create .gitkeep in %s: %w", dirPath, err)
				}
			}
		}

		rootPath := filepath.Join(cwd, ".sheave.toml")
		aiPath := filepath.Join(cwd, ".ai", ".sheave.toml")

		rootExists := false
		if _, err := os.Stat(rootPath); err == nil {
			rootExists = true
		}

		aiExists := false
		if _, err := os.Stat(aiPath); err == nil {
			aiExists = true
		}

		if rootExists && !aiExists {
			move := autoAccept
			if !move {
				fmt.Printf("Found existing .sheave.toml in project root. Move it into .ai/ directory? [y/N]: ")
				reader := bufio.NewReader(os.Stdin)
				response, err := reader.ReadString('\n')
				if err == nil {
					response = strings.ToLower(strings.TrimSpace(response))
					if response == "y" || response == "yes" {
						move = true
					}
				}
			}

			if move {
				if err := os.Rename(rootPath, aiPath); err != nil {
					return fmt.Errorf("failed to move config file: %w", err)
				}
				fmt.Printf("Moved .sheave.toml to .ai/.sheave.toml\n")
				return nil
			}
		}

		if aiExists || rootExists {
			fmt.Printf("Scaffolded .ai directory. Configuration file already exists.\n")
			return nil
		}

		cfg := &config.Config{
			Rules:     config.Selection{Include: []string{"*"}},
			Commands:  config.Selection{Include: []string{"*"}},
			Templates: config.Selection{Include: []string{"*"}},
			Workflows: config.Selection{Include: []string{"*"}},
		}

		if err := cfg.Save(aiPath); err != nil {
			return err
		}

		fmt.Printf("Successfully scaffolded .ai directory and initialized configuration in %s\n", aiPath)
		return nil
	},
}

func init() {
	scaffoldCmd.Flags().BoolP("yes", "y", false, "Automatically accept prompts")
	rootCmd.AddCommand(scaffoldCmd)
}
