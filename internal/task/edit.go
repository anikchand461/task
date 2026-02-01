package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EditTask updates the description of an existing task
func EditTask(taskNumber int, newDescription string) error {
	// Get absolute path to ~/.task/tasks.txt
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	// Open file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Validate task number
	if taskNumber < 1 || taskNumber > len(lines) {
		return fmt.Errorf("invalid task number: %d (valid range: 1-%d)", taskNumber, len(lines))
	}

	// Update task description (preserve status)
	parts := strings.Split(lines[taskNumber-1], "|")
	if len(parts) >= 2 {
		// Keep the status (pending/done), update description
		parts[0] = newDescription
		lines[taskNumber-1] = strings.Join(parts, "|")
	}

	// Write back to file
	return os.WriteFile(
		path,
		[]byte(strings.Join(lines, "\n")+"\n"),
		0644,
	)
}