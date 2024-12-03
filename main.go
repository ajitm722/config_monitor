package main

import (
	"fmt"
	"log"
	"os"

	"config-watcher/watcher"

	"github.com/spf13/viper" // For configuration management
)

func main() {
	// Load configuration
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Get file path and watch method from configuration
	filepath := viper.GetString("file_path")
	watchMethod := viper.GetString("watch_method")

	// Dynamically decide the watching mechanism
	var changes <-chan string
	var errs <-chan error

	switch watchMethod {
	case "polling":
		fmt.Println("Using Polling to watch the configuration file...")
		changes, errs, err = watcher.StartPolling(filepath)
	case "osnotify":
		fmt.Println("Using OS Notifications to watch the configuration file...")
		changes, errs, err = watcher.StartNotify(filepath)
	default:
		log.Fatalf("Invalid watch_method: %v. Use 'polling' or 'osnotify'.", watchMethod)
	}

	if err != nil {
		log.Fatalf("Error initializing watcher: %v", err)
	}

	// Simulate the server running
	fmt.Println("Server is running. Press Ctrl+C to exit.")
	// Listen for changes or errors

	for {
		select {
		case change := <-changes:
			fmt.Printf("Configuration file changed: %v\n\n", change)
		case err := <-errs:
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

}
