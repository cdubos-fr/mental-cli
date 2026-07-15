package cli

import (
	"strings"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

const checkinLabel = "CHECKIN"

func newCheckinCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "checkin [message]",
		Short: "Faire un point rapide, avec un résumé des stats juste après",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			message := strings.Join(args, " ")
			if _, err := store.Append(checkinLabel, message); err != nil {
				return err
			}
			cmd.Println("📍 [CHECKIN] Point fait. Voici où tu en es :")
			cmd.Println()
			return printStats(cmd.OutOrStdout())
		},
	}
}
