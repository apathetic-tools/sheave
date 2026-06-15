package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/registry"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	GroupID: "guidance",
	Use:     "enable [item_id]",
	Short:   "Enable an item by adding it to include",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
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

		item, ok := reg.Get(id)
		if !ok {
			return fmt.Errorf("item %s not found in registry", id)
		}

		var sel *config.Selection
		switch item.Type {
		case "Rule":
			sel = &cfg.Rules
		case "Command":
			sel = &cfg.Commands
		case "Template":
			sel = &cfg.Templates
		case "Workflow":
			sel = &cfg.Workflows
		default:
			return fmt.Errorf("unknown item type %s", item.Type)
		}

		// Remove from Exclude if present
		newExclude := []string{}
		for _, v := range sel.Exclude {
			if v != id {
				newExclude = append(newExclude, v)
			}
		}
		sel.Exclude = newExclude

		// Add to Include if not present
		found := false
		for _, v := range sel.Include {
			if v == id {
				found = true
				break
			}
		}

		if !found {
			sel.Include = append(sel.Include, id)
			if err := cfg.Save(path); err != nil {
				return err
			}
			fmt.Printf("Enabled item: %s (%s)\n", id, item.Type)
		} else {
			fmt.Printf("Item %s is already enabled.\n", id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
