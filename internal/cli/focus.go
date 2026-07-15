package cli

import (
	"fmt"
	"time"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

const focusLabel = "FOCUS"

const defaultFocusDuration = 25 * time.Minute

func newFocusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "focus [duration]",
		Short: "Démarrer une session de focus sur une seule tâche",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			duration := defaultFocusDuration
			if len(args) == 1 {
				d, err := time.ParseDuration(args[0])
				if err != nil {
					return fmt.Errorf("durée invalide %q (ex: 25m, 90s) : %w", args[0], err)
				}
				if d <= 0 {
					return fmt.Errorf("la durée doit être positive, reçu %q", args[0])
				}
				duration = d
			}

			if _, err := store.Append(focusLabel, fmt.Sprintf("started %s", duration)); err != nil {
				return err
			}
			cmd.Printf("🎯 [FOCUS] Session de %s commencée. Une seule tâche, pas de zapping.\n", duration)

			time.Sleep(duration)

			if _, err := store.Append(focusLabel, fmt.Sprintf("completed %s", duration)); err != nil {
				return err
			}
			cmd.Println("✅ [FOCUS] Terminé. Pause méritée.")
			return nil
		},
	}
}
