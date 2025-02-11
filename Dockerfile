# Step 1: Build the Go binary inside a Go container
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy  # Ensures unused dependencies are cleaned up

# Copy the entire application code
COPY . .

# Build the Go binary
RUN go build -o server ./cmd/main.go

# Step 2: Use a lightweight image to run the application
FROM debian:bookworm-slim

# Install necessary libraries (e.g., CA certificates for TLS)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*  # Clean up apt cache to reduce image size

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server /server

# Ensure CA certificates are updated
RUN update-ca-certificates

# Expose the application port
EXPOSE 1000

# Run the compiled Go application
CMD ["/server"]
