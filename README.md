
# File Watcher with Feature Flags

## Overview

This project is a file-watching application that monitors changes in a specified file (`chaotic_file`) using either the `polling` method or the `osnotify` method, depending on the configuration. The system detects file modifications and can respond to changes in real time.

The configuration and feature flags are managed via a `config.yaml` file, and the application can be run using `Makefile` commands. It provides the ability to simulate file changes and ensures the code reacts correctly without manual intervention.

## Prerequisites

Before running the application, ensure the following tools are installed on your system:

- **Docker**: For containerized environments.
- **Docker Compose**: For managing multi-container applications.
- **Make**: For automating the build process.

## Configuration (`config.yaml`)

Before running the application, edit the `config.yaml` file to specify the file to watch (`chaotic_file`) and the watch method (`polling` or `osnotify`).

Example configuration:

```yaml
file_path: ./chaotic_file  # Path to the file being monitored
watch_method: polling      # Method used for file watching ('polling' or 'osnotify')
```

### Watch Methods

- **Polling**: The file is checked for changes at regular intervals (e.g., every 2 seconds).
- **osnotify**: Uses OS-level notifications to detect file changes immediately.

## Workflow

### Step 1: Set up the configuration

Before running any `make` commands, ensure the `config.yaml` file is configured correctly. The file path should point to `chaotic_file`, and the watch method should be set to either `polling` or `osnotify`.

### Step 2: Rebuild the application

Use the `rebuild` command to clean up old containers, build the application, and start the necessary services.

```bash
make rebuild
```

### Step 3: Simulate changes to `chaotic_file`

Open a **second terminal** and use the `echo` command to simulate changes to the `chaotic_file`.

```bash
echo "some changes" >> chaotic_file
```

After running the above command, the application will detect the changes, and the file watcher will react based on the configured watch method.

### Step 4: Observe the behavior

Monitor the output of the file-watching application to see how it reacts when the file is changed. If you're using the `polling` method, the system will detect changes at regular intervals. If using `osnotify`, it will detect changes immediately.

### Step 5: Stop the application

After testing, you can stop the application and clean up resources by running:

```bash
make stop
```

### Step 6: Clean up

To remove all containers and volumes, run:

```bash
make clean
```

To Display all available `make` commands and their descriptions:

```bash
make help
```

## Troubleshooting

- **Watcher not detecting changes**:

  - Ensure that the correct watch method (`polling` or `osnotify`) is selected in the `config.yaml`.
  - If using the `polling` method, ensure the polling interval is appropriate for your use case.
  - Using editors like vim or gedit internally rename and replace the file, thus may not be appropriate for this task.
  
- **Docker issues**:
  - If you encounter issues with Docker containers, try running `make clean` to reset the environment.
  