package watcher

import (
	"errors"
)

// StartPolling initializes the polling watcher
func StartPolling(filepath string) (<-chan string, <-chan error, error) {
	if filepath == "" {
		return nil, nil, errors.New("filepath cannot be empty")
	}
	return Poll(filepath)
}

// StartNotify initializes the OS notification watcher
func StartNotify(filepath string) (<-chan string, <-chan error, error) {
	if filepath == "" {
		return nil, nil, errors.New("filepath cannot be empty")
	}
	return Notify(filepath)
}

