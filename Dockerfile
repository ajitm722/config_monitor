# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the binary
RUN go build -o config-watcher .

# Stage 2: Minimal runtime image
FROM alpine:latest

# Set the working directory in the minimal image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/config-watcher .

# Set the entry point
ENTRYPOINT ["./config-watcher"]

