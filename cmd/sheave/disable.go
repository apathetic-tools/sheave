package main

import (
	"fmt"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable [item_id]",
	Short: "Disable an item by adding it to ignore",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		path := ".sheave.toml"
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		// Ensure it's not in extend-select
		newExtend := []string{}
		for _, v := range cfg.ExtendSelect {
			if v != id {
				newExtend = append(newExtend, v)
			}
		}
		cfg.ExtendSelect = newExtend

		// Add to ignore if not already there
		found := false
		for _, v := range cfg.Ignore {
			if v == id {
				found = true
				break
			}
		}

		if !found {
			cfg.Ignore = append(cfg.Ignore, id)
			if err := cfg.Save(path); err != nil {
				return err
			}
			fmt.Printf("Disabled item: %s\n", id)
		} else {
			fmt.Printf("Item %s is already disabled.\n", id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
