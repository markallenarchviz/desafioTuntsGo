# Start from a minimal GoLang image
FROM golang:1.22-alpine3.18

# Set the Current Working Directory inside the container
WORKDIR /

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o build cmd/main.go


# Command to run the executable
CMD ["build/main"]
