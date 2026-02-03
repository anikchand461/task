package task

import (
		"strings"
	"fmt"
	"os"
)

func Add(task string) error {
		// Validate input: reject empty or whitespace-only strings
	if strings.TrimSpace(task) == "" {
		return fmt.Errorf("error: task description cannot be empty")
	}
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(task + "|pending\n")
	return err
}

