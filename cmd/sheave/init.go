package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize sheave configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ".sheave.toml"
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("Configuration file %s already exists.\n", path)
			return nil
		}

		cfg := &config.Config{
			Select: []string{"ALL"},
		}

		if err := cfg.Save(path); err != nil {
			return err
		}

		fmt.Printf("Initialized configuration in %s\n", path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
