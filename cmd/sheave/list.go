package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/apathetic-tools/sheave/internal/registry"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	GroupID:   "guidance",
	Use:       "list [type]",
	Short:     "List available items",
	ValidArgs: []string{"commands", "rules", "templates", "workflows"},
	RunE: func(cmd *cobra.Command, args []string) error {
		reg := registry.NewRegistry()

		cwd, err := os.Getwd()
		if err == nil {
			// Discover custom items locally
			_ = reg.DiscoverCustomItems(cwd)
		}

		items := reg.List()
		if len(items) == 0 {
			fmt.Println("No items found.")
			return nil
		}

		filter := ""
		if len(args) > 0 {
			filter = strings.ToLower(args[0])
			switch filter {
			case "commands", "command":
				filter = "Command"
			case "rules", "rule":
				filter = "Rule"
			case "templates", "template":
				filter = "Template"
			case "workflows", "workflow":
				filter = "Workflow"
			default:
				return fmt.Errorf("invalid type %q. Valid types are commands, rules, templates, workflows", args[0])
			}
		}

		// Group items by Type
		grouped := make(map[string][]*registry.Item)
		for _, item := range items {
			grouped[item.Type] = append(grouped[item.Type], item)
		}

		// Types to display
		types := []string{"Command", "Rule", "Template", "Workflow"}
		if filter != "" {
			types = []string{filter}
		}

		for _, t := range types {
			groupItems := grouped[t]
			if len(groupItems) == 0 {
				continue
			}

			// Sort by ID
			sort.Slice(groupItems, func(i, j int) bool {
				return groupItems[i].ID < groupItems[j].ID
			})

			fmt.Printf("\n--- %ss ---\n\n", t)
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			_, _ = fmt.Fprintln(w, "ID\tCATEGORY\tBUILTIN\tNAME\tDESCRIPTION")
			for _, p := range groupItems {
				builtinStr := "no"
				if p.Builtin {
					builtinStr = "yes"
				}
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", p.ID, p.Category, builtinStr, p.Name, p.Description)
			}
			_ = w.Flush()
		}
		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
