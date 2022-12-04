package env

import (
	"errors"
	"os"
	"strings"
)

const (
	ENV_FILE_NAME = "config.env"
)

// Load config.env
func Load() error {
	raw_data, err := os.ReadFile(ENV_FILE_NAME)
	if err != nil {
		return err
	}

	data := string(raw_data)

	if data == "" {
		return nil
	}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			return errors.New("cannot extract information - check your syntax")
		}

		err := os.Setenv(pair[0], pair[1])
		if err != nil {
			return err
		}
	}

	return nil
}
