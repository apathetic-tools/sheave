package main

import (
	"fmt"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	GroupID: "guidance",
	Use:     "show",
	Short:   "Show the fully resolved configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ".sheave.toml"
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		active := cfg.Resolve()
		fmt.Printf("Active items (%d total):\n", len(active))
		for _, v := range active {
			fmt.Printf("  - %s\n", v)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
