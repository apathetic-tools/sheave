package main

import (
	"fmt"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable [preset_id]",
	Short: "Enable a preset by adding it to extend-select",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		path := ".sheave.toml"
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		// Ensure it's not in ignore
		newIgnore := []string{}
		for _, v := range cfg.Ignore {
			if v != id {
				newIgnore = append(newIgnore, v)
			}
		}
		cfg.Ignore = newIgnore

		// Add to extend-select if not already there
		found := false
		for _, v := range cfg.ExtendSelect {
			if v == id {
				found = true
				break
			}
		}
		// Also check select
		for _, v := range cfg.Select {
			if v == id {
				found = true
				break
			}
		}

		if !found {
			cfg.ExtendSelect = append(cfg.ExtendSelect, id)
			if err := cfg.Save(path); err != nil {
				return err
			}
			fmt.Printf("Enabled preset: %s\n", id)
		} else {
			fmt.Printf("Preset %s is already enabled.\n", id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
