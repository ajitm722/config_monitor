package watcher

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"time"
)

// Poll uses hashing to detect changes in the file
func Poll(filepath string) (<-chan string, <-chan error, error) {
	changes := make(chan string)
	errs := make(chan error)
	hash := ""

	go func() {
		ticker := time.NewTicker(2 * time.Second) // Poll every 2 seconds
		defer ticker.Stop()

		fmt.Printf("Polling started: Checking the file every 2 seconds...\n\n")

		for range ticker.C {
			fmt.Println("Polling: Checking the file for changes...")
			newHash, err := calculateFileHash(filepath)
			if err != nil {
				fmt.Printf("Polling Error: Failed to calculate file hash: %v\n\n", err)
				errs <- err
				continue
			}

			if newHash != hash {
				fmt.Printf("File Change Detected: Previous Hash: %s, New Hash: %s\n\n", hash, newHash)
				hash = newHash
				changes <- filepath
			} else {
				fmt.Printf("No Change Detected: File content remains the same.\n\n")
			}
		}
	}()

	return changes, errs, nil
}

// calculateFileHash calculates the SHA-256 hash of the file
func calculateFileHash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
