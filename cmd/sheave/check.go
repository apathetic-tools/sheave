package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/registry"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate configuration against known items",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ".sheave.toml"
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		reg := registry.NewRegistry()
		cwd, err := os.Getwd()
		if err == nil {
			_ = reg.DiscoverCustomItems(cwd)
		}

		active := cfg.Resolve()
		var unknown []string
		for _, v := range active {
			if v == "ALL" {
				continue
			}
			if _, ok := reg.Get(v); !ok {
				unknown = append(unknown, v)
			}
		}

		if len(unknown) > 0 {
			fmt.Printf("Warning: Found %d unknown item(s) in configuration:\n", len(unknown))
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
