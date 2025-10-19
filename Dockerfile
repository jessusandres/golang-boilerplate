# Build stage
FROM golang:1.25.1-alpine AS builder

# Create a working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies (will be cached if go.mod/go.sum don't change)
RUN go mod download

# Copy the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o app ./cmd/server

# Final stage: Use a minimal alpine image
FROM alpine:3.18

# FROM gcr.io/distroless/static-debian11

# WORKDIR /app

# COPY --from=builder /app/app .

# EXPOSE 8080

# # The distroless image doesn't have a shell, so use the binary directly
# CMD ["/app/app"]


# Add CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create a non-root user to run the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy only the binary from the build stage
COPY --from=builder /app/app .

# Use the non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
