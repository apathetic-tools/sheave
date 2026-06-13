package main

import (
	"fmt"
	"os"

	"github.com/apathetic-tools/sheave/internal/config"
	"github.com/apathetic-tools/sheave/internal/preset"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate configuration against known presets",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ".sheave.toml"
		cfg, err := config.Load(path)
		if err != nil {
			return err
		}

		reg := preset.NewRegistry()
		cwd, err := os.Getwd()
		if err == nil {
			_ = reg.DiscoverCustomPresets(cwd)
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
			fmt.Printf("Warning: Found %d unknown preset(s) in configuration:\n", len(unknown))
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
