// Package cli wires up the mental-cli cobra commands.
package cli

import (
	"github.com/spf13/cobra"
)

// version is set at build time via -ldflags.
var version = "dev"

// NewRootCmd builds the root "mental" command with all subcommands attached.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "mental",
		Short:   "Drop a thought in the blackhole and close the loop.",
		Long:    "mental helps you dump anxious or looping thoughts out of your head\nand into a log, so you can get back to work.",
		Version: version,
	}

	for _, a := range actions {
		root.AddCommand(newActionCmd(a))
	}
	root.AddCommand(newPlopCmd())
	root.AddCommand(newLogCmd())
	root.AddCommand(newStatsCmd())

	return root
}
