package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/registry"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	GroupID: "guidance",
	Use:     "show",
	Short:   "Show the fully resolved configuration",
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

		activeRules := reg.Resolve("Rule", cfg.Rules.Include, cfg.Rules.Exclude)
		activeCommands := reg.Resolve("Skill", cfg.Skills.Include, cfg.Skills.Exclude)
		activeTemplates := reg.Resolve("Template", cfg.Templates.Include, cfg.Templates.Exclude)
		activeWorkflows := reg.Resolve("Workflow", cfg.Workflows.Include, cfg.Workflows.Exclude)

		total := len(activeRules) + len(activeCommands) + len(activeTemplates) + len(activeWorkflows)

		fmt.Printf("Active items (%d total):\n", total)

		printGroup := func(name string, items []*registry.Item) {
			if len(items) > 0 {
				fmt.Printf("\n--- %ss ---\n", name)
				for _, v := range items {
					key := v.ID
					if v.Family != "" {
						key = v.Family + "/" + v.ID
					}
					fmt.Printf("  - %s\n", key)
				}
			}
		}

		printGroup("Rule", activeRules)
		printGroup("Skill", activeCommands)
		printGroup("Template", activeTemplates)
		printGroup("Workflow", activeWorkflows)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
