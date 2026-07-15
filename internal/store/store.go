// Package store persists mental-cli entries as newline-delimited JSON
// under the XDG data directory.
package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Entry is a single logged thought.
type Entry struct {
	Time    time.Time `json:"time"`
	Action  string    `json:"action"`
	Message string    `json:"message"`
}

// Path returns the JSONL log file path, honoring XDG_DATA_HOME.
func Path() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home directory: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(dataHome, "mental", "dump.log"), nil
}

// Append writes a new entry to the log file, creating it (and its parent
// directory) if needed.
func Append(action, message string) (Entry, error) {
	entry := Entry{Time: time.Now(), Action: action, Message: message}

	path, err := Path()
	if err != nil {
		return entry, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return entry, fmt.Errorf("create log directory: %w", err)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return entry, fmt.Errorf("open log file: %w", err)
	}
	defer func() { _ = f.Close() }()

	line, err := json.Marshal(entry)
	if err != nil {
		return entry, fmt.Errorf("encode entry: %w", err)
	}
	if _, err := f.Write(append(line, '\n')); err != nil {
		return entry, fmt.Errorf("write entry: %w", err)
	}
	return entry, nil
}

// All reads every entry from the log file, oldest first. A missing log
// file yields an empty slice, not an error.
func All() ([]Entry, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("open log file: %w", err)
	}
	defer func() { _ = f.Close() }()

	var entries []Entry
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var e Entry
		if err := json.Unmarshal(line, &e); err != nil {
			return nil, fmt.Errorf("decode entry: %w", err)
		}
		entries = append(entries, e)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read log file: %w", err)
	}
	return entries, nil
}
