package main

import (
	"os"

	"github.com/apathetic-tools/sheave/internal/sync"
	"github.com/spf13/cobra"
)

var (
	quiet  bool
	dryRun bool
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync AI guidance files from .ai/ to IDE-specific directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := sync.Options{
			Quiet:  quiet,
			DryRun: dryRun,
		}

		projectRoot, err := os.Getwd()
		if err != nil {
			return err
		}

		_, err = sync.SyncToIDE(projectRoot, opts)
		return err
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Suppress output messages")
	syncCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simulate changes without modifying files")
	rootCmd.AddCommand(syncCmd)
}
