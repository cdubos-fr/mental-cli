package cli

import (
	"strings"

	"github.com/cdubos-fr/mental-cli/internal/store"
	"github.com/spf13/cobra"
)

// action describes one of the fixed mental-cli commands: it logs a message
// under actionLabel and prints output to close the loop.
type action struct {
	use    string
	short  string
	label  string
	output string
}

// actions is the exact port of the 8 commands from the original zsh
// `mental()` function, wording included.
var actions = []action{
	{
		use:    "dump",
		short:  "Évacuer une pensée sans l'analyser",
		label:  "DUMP",
		output: "🧠 [DUMP] Pensée évacuée et archivée. RAM libérée.",
	},
	{
		use:    "break",
		short:  "Stopper une boucle infinie de rumination",
		label:  "BREAK",
		output: "🛑 [BREAK] Boucle infinie stoppée.\nMessage enregistré. Interdiction d'analyser.",
	},
	{
		use:    "loop",
		short:  "Signaler une analyse sans nouvelles données",
		label:  "LOOP_DETECTED",
		output: "🔄 [LOOP] Alerte : Tu tentes de résoudre une équation sans nouvelles données.\nAction : Arrêt immédiat de l'analyse. Retourne coder.",
	},
	{
		use:    "ping",
		short:  "Signaler une impulsion (besoin de vérifier/écrire)",
		label:  "PING_IMPULSE",
		output: "⚡ [PING] Impulsion détectée (besoin de vérifier/écrire).\nAction requise : Pose ton téléphone. Bois un grand verre d'eau ou fais 10 pompes maintenant.",
	},
	{
		use:    "refactor",
		short:  "Réécrire une pensée irrationnelle",
		label:  "REFACTOR",
		output: "🔧 [REFACTOR] Pensée irrationnelle réécrite et corrigée en base de données.",
	},
	{
		use:    "commit",
		short:  "Verrouiller une décision, sujet clos",
		label:  "COMMIT",
		output: "💾 [COMMIT] Décision verrouillée en prod. Pas de rollback possible.\nAction : Sujet clos jusqu'à la date définie.",
	},
	{
		use:    "chop",
		short:  "Casser le rythme d'un échange mental qui dérape",
		label:  "CHOP_BLOCK",
		output: "🏓 [CHOP] Chop block ! Tu casses le rythme de l'échange mental.\nAction : Respire à fond, change d'activité immédiatement.",
	},
	{
		use:    "debug",
		short:  "Signaler un bug physique/sensoriel",
		label:  "ENVIRONMENT_DEBUG",
		output: "⚙️ [DEBUG] Bug physique/sensoriel identifié.\nAction : Isole-toi du bruit, aère la pièce ou change d'environnement immédiatement.",
	},
}

func newActionCmd(a action) *cobra.Command {
	return &cobra.Command{
		Use:   a.use + " <message>",
		Short: a.short,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			message := strings.Join(args, " ")
			if _, err := store.Append(a.label, message); err != nil {
				return err
			}
			cmd.Println(a.output)
			return nil
		},
	}
}
