# Start from the official Golang image for building
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build tools for cgo
RUN apk add --no-cache build-base

# Copy go.mod and go.sum then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=1 go build -o cba cmd/cba/main.go

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary and config file
COPY --from=builder /app/cba .
COPY cmd/cba/cba.yaml .

# Run the application
CMD ["./cba"]