FROM golang:1.20-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go/go.mod .
COPY go/go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY go .

# Build the Go application
RUN go build -o main .

# Command to run the app
CMD ["/app/main"]
