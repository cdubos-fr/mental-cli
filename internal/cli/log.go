package cli

import (
	"fmt"
	"strings"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

func newLogCmd() *cobra.Command {
	var (
		actionFilter string
		limit        int
	)

	cmd := &cobra.Command{
		Use:     "log",
		Aliases: []string{"history"},
		Short:   "Lister les dernières pensées archivées",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entries, err := store.All()
			if err != nil {
				return err
			}

			filter := strings.ToUpper(actionFilter)
			var filtered []store.Entry
			for _, e := range entries {
				if filter != "" && e.Action != filter {
					continue
				}
				filtered = append(filtered, e)
			}

			if len(filtered) == 0 {
				cmd.Println("Aucune pensée archivée pour le moment.")
				return nil
			}

			start := 0
			if limit > 0 && len(filtered) > limit {
				start = len(filtered) - limit
			}
			for i := len(filtered) - 1; i >= start; i-- {
				e := filtered[i]
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s [%s] %s\n",
					e.Time.Local().Format("2006-01-02 15:04"), e.Action, e.Message)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&actionFilter, "action", "", "filtrer par type d'action (ex: DUMP, LOOP_DETECTED, PLOP)")
	cmd.Flags().IntVar(&limit, "limit", 20, "nombre maximum d'entrées à afficher (0 = tout)")

	_ = cmd.RegisterFlagCompletionFunc("action", func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		return actionLabels(), cobra.ShellCompDirectiveNoFileComp
	})

	return cmd
}

// actionLabels lists every action a thought can be logged under, action
// commands first — used for shell completion of --action.
func actionLabels() []string {
	labels := make([]string, 0, len(actions)+4)
	for _, a := range actions {
		labels = append(labels, a.label)
	}
	return append(labels, plopLabel, stimLabel, focusLabel, checkinLabel)
}
