# Start with a base image containing Go
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod ./
COPY main.go ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o payment-service .

# Run the application
CMD ["./payment-service"]
