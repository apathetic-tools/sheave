package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/apathetic-tools/sheave/internal/preset"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available presets",
	RunE: func(cmd *cobra.Command, args []string) error {
		reg := preset.NewRegistry()

		cwd, err := os.Getwd()
		if err == nil {
			// Discover custom presets locally
			_ = reg.DiscoverCustomPresets(cwd)
		}

		presets := reg.List()
		if len(presets) == 0 {
			fmt.Println("No presets found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		_, _ = fmt.Fprintln(w, "ID\tCATEGORY\tBUILTIN\tNAME\tDESCRIPTION")
		for _, p := range presets {
			builtinStr := "no"
			if p.Builtin {
				builtinStr = "yes"
			}
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", p.ID, p.Category, builtinStr, p.Name, p.Description)
		}
		_ = w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
