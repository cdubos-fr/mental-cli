package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cdubos-fr/mental-cli/internal/cli"
)

// run executes the root command against whatever XDG_DATA_HOME is already
// set for the calling test, so repeated calls within a test share history.
func run(t *testing.T, args ...string) string {
	t.Helper()
	root := cli.NewRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs(args)

	if err := root.Execute(); err != nil {
		t.Fatalf("Execute(%v) error = %v", args, err)
	}
	return out.String()
}

func TestActionCommandsRequireAMessage(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())

	for _, use := range []string{"dump", "break", "loop", "ping", "refactor", "commit", "chop", "debug", "win"} {
		root := cli.NewRootCmd()
		var out bytes.Buffer
		root.SetOut(&out)
		root.SetErr(&out)
		root.SetArgs([]string{use})

		if err := root.Execute(); err == nil {
			t.Errorf("%s with no message: expected an error, got none", use)
		}
	}
}

func TestDumpLogsAndPrintsConfirmation(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	out := run(t, "dump", "a", "racing", "thought")

	if !strings.Contains(out, "RAM libérée") {
		t.Errorf("dump output = %q, want confirmation message", out)
	}

	history := run(t, "log")
	if !strings.Contains(history, "[DUMP] a racing thought") {
		t.Errorf("log output = %q, want it to contain the dumped message", history)
	}
}

func TestPlopNeedsNoMessageAndIsLogged(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	out := run(t, "plop")
	if !strings.Contains(out, "Plop") {
		t.Errorf("plop output = %q, want it to contain a plop phrase", out)
	}

	history := run(t, "log")
	if !strings.Contains(history, "[PLOP]") {
		t.Errorf("log output = %q, want a PLOP entry", history)
	}
}

func TestWinLogsAndPrintsConfirmation(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	out := run(t, "win", "shipped", "the", "thing")

	if !strings.Contains(out, "changelog des victoires") {
		t.Errorf("win output = %q, want confirmation message", out)
	}

	history := run(t, "log")
	if !strings.Contains(history, "[WIN] shipped the thing") {
		t.Errorf("log output = %q, want it to contain the win message", history)
	}
}

func TestStimNeedsNoMessageAndIsLogged(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	out := run(t, "stim")
	if !strings.Contains(out, "Stim break") {
		t.Errorf("stim output = %q, want it to contain a stim suggestion", out)
	}

	history := run(t, "log")
	if !strings.Contains(history, "[STIM]") {
		t.Errorf("log output = %q, want a STIM entry", history)
	}
}

func TestFocusRunsForGivenDurationAndLogsStartAndEnd(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	out := run(t, "focus", "5ms")

	if !strings.Contains(out, "commencée") || !strings.Contains(out, "Terminé") {
		t.Errorf("focus output = %q, want start and completion messages", out)
	}

	history := run(t, "log", "--action", "focus")
	if !strings.Contains(history, "started 5ms") || !strings.Contains(history, "completed 5ms") {
		t.Errorf("log output = %q, want FOCUS start and completion entries", history)
	}
}

func TestFocusRejectsInvalidDuration(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	root := cli.NewRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"focus", "not-a-duration"})

	if err := root.Execute(); err == nil {
		t.Error("focus with an invalid duration: expected an error, got none")
	}
}

func TestCheckinLogsAndPrintsStats(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	run(t, "dump", "earlier thought")

	out := run(t, "checkin", "tired but ok")

	if !strings.Contains(out, "Point fait") {
		t.Errorf("checkin output = %q, want confirmation message", out)
	}
	if !strings.Contains(out, "ACTION") || !strings.Contains(out, "DUMP") {
		t.Errorf("checkin output = %q, want it to include the stats table", out)
	}

	history := run(t, "log", "--action", "checkin")
	if !strings.Contains(history, "[CHECKIN] tired but ok") {
		t.Errorf("log output = %q, want it to contain the checkin message", history)
	}
}

func TestLogFilterByAction(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	run(t, "dump", "keep this out")
	run(t, "loop", "keep this in")

	out := run(t, "log", "--action", "loop_detected")
	if strings.Contains(out, "keep this out") {
		t.Errorf("log --action filter leaked an unrelated entry: %q", out)
	}
	if !strings.Contains(out, "keep this in") {
		t.Errorf("log --action filter dropped the matching entry: %q", out)
	}
}

func TestStatsCountsByAction(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	run(t, "dump", "one")
	run(t, "dump", "two")
	run(t, "loop", "three")

	out := run(t, "stats")
	if !strings.Contains(out, "DUMP") || !strings.Contains(out, "2") {
		t.Errorf("stats output = %q, want DUMP count of 2", out)
	}
}
