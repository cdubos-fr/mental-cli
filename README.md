# mental

A tiny CLI for dropping anxious or looping thoughts into a blackhole so you
can close the loop and get back to work.

Each command logs a timestamped entry and prints a short confirmation —
the point isn't to analyze the thought, it's to acknowledge it left your
head and move on.

## Install

```sh
go install github.com/cdubos-fr/mental-cli/cmd/mental@latest
```

Or grab a prebuilt binary from the [releases page](https://github.com/cdubos-fr/mental-cli/releases).

## Usage

```sh
mental dump "message"       # évacuer une pensée sans l'analyser
mental break "message"      # stopper une boucle infinie de rumination
mental loop "message"       # signaler une analyse sans nouvelles données
mental ping "message"       # signaler une impulsion (vérifier/écrire)
mental refactor "message"   # réécrire une pensée irrationnelle
mental commit "message"     # verrouiller une décision, sujet clos
mental chop "message"       # casser le rythme d'un échange qui dérape
mental debug "message"      # signaler un bug physique/sensoriel

mental plop                 # lâcher une pensée aléatoire, sans message à donner

mental log                  # lister les dernières pensées archivées
mental log --action dump    # filtrer par type d'action
mental log --limit 0        # tout afficher

mental stats                # compter les pensées archivées par type
```

Entries are appended as JSON Lines to `$XDG_DATA_HOME/mental/dump.log`
(defaults to `~/.local/share/mental/dump.log`).

## Development

Tool versions (Go, just, golangci-lint, goreleaser, pre-commit, zizmor) are
pinned in `mise.toml`. `.envrc` (direnv) checks they're present and
installs the git hooks automatically when you `cd` into the project —
run `mise install` once if a tool is missing.

```sh
just              # list all recipes
just build        # build ./bin/mental
just run dump "message"
just test
just check        # fmt-check + vet + lint + test + security — what CI runs
just release-snapshot  # local GoReleaser dry run, no publish
```

CI runs `just check` on every push/PR. Pushing a `vX.Y.Z` tag triggers
GoReleaser to cross-compile and publish a GitHub release.
