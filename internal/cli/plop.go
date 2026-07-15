package cli

import (
	"math/rand/v2"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

const plopLabel = "PLOP"

// plopPhrases are random, message-free tension releases: no thought to
// name, just noise leaving the system.
var plopPhrases = []string{
	"🫧 Plop. Une pensée aléatoire vient de s'auto-supprimer.",
	"🐸 Plop ! Un crapaud a traversé ta RAM et n'a rien laissé derrière lui.",
	"🎈 Plop. Ballon de stress dégonflé sans bruit.",
	"🧦 Plop. Une chaussette dépareillée est apparue, personne ne sait pourquoi.",
	"🛸 Plop. Un vaisseau alien a aspiré ta pensée avant qu'elle ne devienne un problème.",
	"🍄 Plop. Un champignon a poussé dans ton cerveau pendant la nuit. Il va bien.",
	"🪄 Plop. Abracadabra, le sujet a disparu.",
	"🎪 Plop. Le cirque mental ferme ses portes pour aujourd'hui.",
	"🧃 Plop. Une brique de jus s'est ouverte quelque part dans l'univers. Pas ton problème.",
	"🐙 Plop. Une pieuvre a piqué ta pensée et s'enfuit en jetant de l'encre.",
	"🚀 Plop. Décollage réussi vers nulle part en particulier.",
	"🎲 Plop. Lancer de dé random : résultat inutile, comme prévu.",
	"🦆 Plop. Un canard passe. Rien à signaler.",
	"🧊 Plop. Glaçon jeté dans le vide mental. Aucun bruit à l'impact.",
	"🪩 Plop. Boule disco activée. Ambiance non justifiée mais bienvenue.",
}

func newPlopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "plop",
		Short: "Lâcher une pensée aléatoire, sans message à donner",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			phrase := plopPhrases[rand.IntN(len(plopPhrases))]
			if _, err := store.Append(plopLabel, phrase); err != nil {
				return err
			}
			cmd.Println(phrase)
			return nil
		},
	}
}
