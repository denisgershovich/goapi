package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Load reads environment variables from a .env file in the project root.
// Lines beginning with '#' are comments. Blank lines are ignored.
// Each non-empty line should be in the form KEY=VALUE. Whitespace around keys
// and values is trimmed. Double or single quotes around values are removed.
func Load() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get current file path for config")
	}
	// internal/config -> project root
	rootDir := filepath.Dir(filepath.Dir(filename))
	envPath := filepath.Join(rootDir, ".env")

	file, err := os.Open(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // .env is optional
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// support inline comments after value with a space then #
		if hash := strings.Index(line, " #"); hash != -1 {
			line = strings.TrimSpace(line[:hash])
		}
		eq := strings.Index(line, "=")
		if eq <= 0 {
			continue
		}
		key := strings.TrimSpace(line[:eq])
		val := strings.TrimSpace(line[eq+1:])
		val = strings.TrimSuffix(strings.TrimPrefix(val, "\""), "\"")
		val = strings.TrimSuffix(strings.TrimPrefix(val, "'"), "'")
		_ = os.Setenv(key, val)
	}
	return scanner.Err()
}
