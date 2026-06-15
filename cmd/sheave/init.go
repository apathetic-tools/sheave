package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: "utility",
	Use:     "init",
	Short:   "Initialize sheave configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := config.GetConfigPath(cwd)
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("Configuration file %s already exists.\n", path)
			return nil
		}

		cfg := &config.Config{
			Rules:     config.Selection{Include: []string{"~*"}},
			Commands:  config.Selection{Include: []string{"~*"}},
			Templates: config.Selection{Include: []string{"~*"}},
			Workflows: config.Selection{Include: []string{"~*"}},
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
