package cli

import (
	"math/rand/v2"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

const stimLabel = "STIM"

// stimSuggestions are concrete, short sensory/regulation actions — unlike
// plop these are meant to actually be followed, not just read.
var stimSuggestions = []string{
	"🧶 Stim break. Serre un objet 10 secondes, relâche, recommence trois fois.",
	"🎧 Stim break. Casque sur les oreilles, bruit blanc deux minutes.",
	"🧊 Stim break. Glaçon dans la main jusqu'à ce qu'il fonde un peu.",
	"🌬️ Stim break. Respire : 4 secondes inspire, 4 secondes retiens, 6 secondes expire.",
	"🪢 Stim break. Triture un élastique ou une pâte à modeler pendant une minute.",
	"🚶 Stim break. Lève-toi, marche jusqu'au bout du couloir et reviens.",
	"🖐️ Stim break. Presse tes paumes l'une contre l'autre fort pendant 10 secondes.",
	"🎵 Stim break. Mets un son répétitif que tu aimes, ferme les yeux 30 secondes.",
	"🧣 Stim break. Enroule-toi dans une couverture ou un vêtement serré une minute.",
	"💡 Stim break. Baisse la lumière ou ferme les yeux, réduis le bruit ambiant.",
}

func newStimCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stim",
		Short: "Suggérer une pause sensorielle/de régulation, sans message à donner",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			suggestion := stimSuggestions[rand.IntN(len(stimSuggestions))]
			if _, err := store.Append(stimLabel, suggestion); err != nil {
				return err
			}
			cmd.Println(suggestion)
			return nil
		},
	}
}
