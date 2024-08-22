# Use the official Golang image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go/go.mod ./
COPY go/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the Go files to the working directory
COPY go/ ./

# Build the Go application
RUN go build -o main .

# Specify the command to run the application
CMD ["./main"]
