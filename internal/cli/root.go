// Package cli wires up the mental-cli cobra commands.
package cli

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

// version is set at build time via -ldflags (GoReleaser). When absent —
// e.g. `go install .../cmd/mental@version` — fall back to the module
// version Go itself embeds in the binary's build info.
var version = "dev"

func init() {
	if version != "dev" {
		return
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		version = info.Main.Version
	}
}

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
	root.AddCommand(newStimCmd())
	root.AddCommand(newFocusCmd())
	root.AddCommand(newCheckinCmd())
	root.AddCommand(newLogCmd())
	root.AddCommand(newStatsCmd())

	return root
}
