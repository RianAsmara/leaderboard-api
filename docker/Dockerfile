# Use Golang base image
FROM golang:1.20 as builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder
COPY --from=builder /app/main /main

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
