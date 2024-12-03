package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func Notify(filepath string) (<-chan string, <-chan error, error) {
	changes := make(chan string)
	errs := make(chan error)

	// Create a new fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("Notify Error: Failed to create a file watcher: %v\n", err)
		return nil, nil, err
	}

	fmt.Printf("Notify: Watching file '%s' for changes...\n", filepath)

	go func() {
		defer func() {
			log.Println("Notify: Closing the watcher...")
			close(errs)
			watcher.Close()
		}()

		// Add the file to the watcher
		err := watcher.Add(filepath)
		if err != nil {
			log.Printf("Notify Error: Failed to add file '%s': %v\n", filepath, err)
			errs <- err
			return
		}

		// Log all received events
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("Notify: Event channel closed.")
					return
				}
				log.Printf("Notify: Event received: %v\n", event)

				// Handle write events
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("Notify: File '%s' has been modified.\n", event.Name)
					changes <- event.Name
				} else {
					// Log and exit on other events
					fmt.Println("DEBUG INFO: Editors like 'vim', 'gedit' often save by renaming the file first, causing rename events.")
					fmt.Printf("Notify: Unsupported Event '%v' for file '%s'. Exiting...\n", event.Op, event.Name)
					return
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					log.Println("Notify: Error channel closed.")
					return
				}
				log.Printf("Notify Error: %v\n", err)
				errs <- err
			}
		}
	}()

	return changes, errs, nil
}
