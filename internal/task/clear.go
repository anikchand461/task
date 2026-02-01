package task

import (
	"os"
)

func Clear() error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}
	
	// Delete the file
	return os.Remove(path)
}