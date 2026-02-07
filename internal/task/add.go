package task

import "os"

func Add(tasks ...string) error {
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

	for _, task := range tasks {
		if task == "" {
			continue
		}
		_, err := f.WriteString(task + "|pending\n")
		if err != nil {
			return err
		}
	}

	return nil
}
