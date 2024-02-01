# Use the official Golang image as the base image
FROM golang:1.21.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Install gcc and musl-dev for cgo
RUN apk add --no-cache gcc musl-dev

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . ./

# Remove test related files
RUN rm -r ./test
RUN find . -name "*_test.go" -exec rm {} \;

# Build the Go application
RUN CGO_ENABLED=1 go build -o ./main ./cmd/app

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]
