# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev git

# Copy only go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies with retries
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -o smsc-gateway ./cmd/smsc/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy only the binary from builder
COPY --from=builder /app/smsc-gateway .

EXPOSE 8080 2775 2776

CMD ["./smsc-gateway"] 