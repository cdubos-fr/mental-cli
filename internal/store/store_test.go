package store_test

import (
	"path/filepath"
	"testing"

	"github.com/cdubos-fr/mental-cli/internal/store"
)

func TestPathHonorsXDGDataHome(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/tmp/xdg-example")

	got, err := store.Path()
	if err != nil {
		t.Fatalf("Path() error = %v", err)
	}
	want := filepath.Join("/tmp/xdg-example", "mental", "dump.log")
	if got != want {
		t.Errorf("Path() = %q, want %q", got, want)
	}
}

func TestAppendAndAll(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())

	if _, err := store.Append("DUMP", "first thought"); err != nil {
		t.Fatalf("Append() error = %v", err)
	}
	if _, err := store.Append("LOOP_DETECTED", "second thought"); err != nil {
		t.Fatalf("Append() error = %v", err)
	}

	entries, err := store.All()
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("All() returned %d entries, want 2", len(entries))
	}
	if entries[0].Action != "DUMP" || entries[0].Message != "first thought" {
		t.Errorf("entries[0] = %+v, want action DUMP message %q", entries[0], "first thought")
	}
	if entries[1].Action != "LOOP_DETECTED" || entries[1].Message != "second thought" {
		t.Errorf("entries[1] = %+v, want action LOOP_DETECTED message %q", entries[1], "second thought")
	}
}

func TestAllWithoutExistingLogFile(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())

	entries, err := store.All()
	if err != nil {
		t.Fatalf("All() error = %v", err)
	}
	if entries != nil {
		t.Errorf("All() = %v, want nil for missing log file", entries)
	}
}
