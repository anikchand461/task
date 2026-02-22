package task

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)


func AcquireLock() (func(), error) {
	path, err := dataFilePath()
	if err != nil {
		return nil, err
	}
	
	// Lock file acts as a flag: if it exists, the DB is busy.
	lockPath := filepath.Join(filepath.Dir(path), "tasks.lock")

	for i := 0; i < 10; i++ { // Try 10 times (1 second total)
		
		f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL, 0644)
		if err == nil {
			
			f.Close()
			
			return func() {
				os.Remove(lockPath)
			}, nil
		}

		
		if !os.IsExist(err) {
			return nil, err
		}

		
		time.Sleep(100 * time.Millisecond)
	}

	return nil, fmt.Errorf("could not acquire lock: another task process is running")
}