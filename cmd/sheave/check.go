package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/registry"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	GroupID: "guidance",
	Use:     "check",
	Short:   "Validate configuration against known items",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := config.GetConfigPath(cwd)
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		reg := registry.NewRegistry()
		_ = reg.DiscoverCustomItems(cwd)

		var unknown []string

		checkSelection := func(sel config.Selection, itemType string) {
			for _, pattern := range sel.Include {
				if len(reg.MatchPattern(pattern, itemType)) == 0 {
					unknown = append(unknown, fmt.Sprintf("%s (in %ss Include)", pattern, itemType))
				}
			}
			for _, pattern := range sel.Exclude {
				if len(reg.MatchPattern(pattern, itemType)) == 0 {
					unknown = append(unknown, fmt.Sprintf("%s (in %ss Exclude)", pattern, itemType))
				}
			}
		}

		checkSelection(cfg.Rules, "Rule")
		checkSelection(cfg.Commands, "Command")
		checkSelection(cfg.Templates, "Template")
		checkSelection(cfg.Workflows, "Workflow")

		if len(unknown) > 0 {
			fmt.Printf("Warning: Found %d unmatched pattern(s) in configuration:\n", len(unknown))
			for _, v := range unknown {
				fmt.Printf("  - %s\n", v)
			}
		} else {
			fmt.Println("Configuration is valid.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
