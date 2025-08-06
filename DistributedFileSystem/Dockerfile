# Development Dockerfile with hot reload
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for go mod download)
RUN apk add --no-cache git ca-certificates

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

# Install Swagger for API documentation (optional)
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod files first for better caching
COPY go.mod* go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8084

# Command to run with hot reload
CMD ["air"]