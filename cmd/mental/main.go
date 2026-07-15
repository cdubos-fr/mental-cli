// Command mental drops thoughts into a blackhole so you can close the loop.
package main

import (
	"os"

	"github.com/cdubos-fr/mental-cli/internal/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
