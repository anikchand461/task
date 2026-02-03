package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Done(idStr string) error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	// Parse comma-separated IDs
	idParts := strings.Split(idStr, ",")
	idsToMark := make(map[int]bool)

	for _, part := range idParts {
		trimmed := strings.TrimSpace(part)
		id, err := strconv.Atoi(trimmed)
		if err != nil {
			return fmt.Errorf("invalid task id: %s", trimmed)
		}
		idsToMark[id] = true
	}

	// Read all tasks
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var tasks []string
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if idsToMark[i] {
			// Mark this task as done
			parts := strings.Split(line, "|")
			if len(parts) >= 2 {
				parts[1] = "done"
				line = strings.Join(parts, "|")
			}
		}
		tasks = append(tasks, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back all tasks
	f, err = os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, task := range tasks {
		_, err = f.WriteString(task + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func DoneAll() error {
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

	// Mark all tasks as done
	for i, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) >= 2 {
			parts[1] = "done"
			lines[i] = strings.Join(parts, "|")
		}
	}

	// Write back to file
	return os.WriteFile(
		path,
		[]byte(strings.Join(lines, "\n")+"\n"),
		0644,
	)
}


