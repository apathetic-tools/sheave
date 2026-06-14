package main

import (
	"os"

	"github.com/apathetic-tools/sheave/internal/project"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	GroupID: "utility",
	Use:     "project",
	Short:   "Generate a general summary of the project structure",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		return project.GenerateSummary(cwd)
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
