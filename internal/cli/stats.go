package cli

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

type actionCount struct {
	label string
	total int
	week  int
}

func newStatsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Compter les pensées archivées par type",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return printStats(cmd.OutOrStdout())
		},
	}
}

// printStats writes the per-action counts table to out. Shared by the
// stats command and checkin, which shows it right after logging.
func printStats(out io.Writer) error {
	entries, err := store.All()
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		_, _ = fmt.Fprintln(out, "Aucune pensée archivée pour le moment.")
		return nil
	}

	weekAgo := time.Now().AddDate(0, 0, -7)
	counts := map[string]*actionCount{}
	for _, e := range entries {
		c, ok := counts[e.Action]
		if !ok {
			c = &actionCount{label: e.Action}
			counts[e.Action] = c
		}
		c.total++
		if e.Time.After(weekAgo) {
			c.week++
		}
	}

	list := make([]*actionCount, 0, len(counts))
	for _, c := range counts {
		list = append(list, c)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].total != list[j].total {
			return list[i].total > list[j].total
		}
		return list[i].label < list[j].label
	})

	_, _ = fmt.Fprintf(out, "%-20s %8s %12s\n", "ACTION", "TOTAL", "7 DERNIERS J")
	for _, c := range list {
		_, _ = fmt.Fprintf(out, "%-20s %8d %12d\n", c.label, c.total, c.week)
	}
	_, _ = fmt.Fprintf(out, "%-20s %8d\n", "TOTAL", len(entries))
	return nil
}
