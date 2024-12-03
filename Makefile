# Define Docker Compose options
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yml

# Define image and container names
IMAGE_NAME = config-watcher
CONTAINER_NAME = config-watcher-container

# Default target: Build and start the services
.PHONY: up
up: build
	@$(DOCKER_COMPOSE) up

# Build the Docker images (without cache to ensure fresh builds)
.PHONY: build
build:
	@$(DOCKER_COMPOSE) build --no-cache

# Clean up stopped containers, networks, and volumes
.PHONY: clean
clean:
	@$(DOCKER_COMPOSE) down -v

# Stop the containers without removing volumes
.PHONY: stop
stop:
	@$(DOCKER_COMPOSE) down

# Start the containers without rebuilding
.PHONY: start
start:
	@$(DOCKER_COMPOSE) up

# Rebuild images and restart the containers
.PHONY: rebuild
rebuild: clean build up

# Help section: Show available commands and their explanations
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  up       - Build and start the Docker containers."
	@echo "  build    - Build the Docker images without cache."
	@echo "  clean    - Stop the containers, networks, and volumes."
	@echo "  stop     - Stop the containers without removing volumes."
	@echo "  start    - Start the containers without rebuilding."
	@echo "  rebuild  - Clean, rebuild the images, and start the containers."
	@echo "  help     - Show this help message."


