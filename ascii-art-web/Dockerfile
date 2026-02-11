FROM golang:1.24-alpine AS builder

# METADATA - Builder stage
LABEL maintainer="jvictori, phaturo, makerubo"
LABEL version="1.0.0"
LABEL description="ASCII Art Web Application - Builder Stage"
LABEL project="ascii-art-web"

WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o ascii-art-web .

# Production stage
FROM alpine:latest

# METADATA - Production image (PROJECT REQUIREMENT)
LABEL maintainer="jvictori, phaturo, makerubo"
LABEL version="1.0.0"
LABEL description="Production ASCII Art Web Server "
LABEL project="ascii-art-web"
LABEL dockerfile.version="1.0"

# Add non-root user for security
RUN addgroup -g 1000 -S appuser && \
    adduser -u 1000 -S appuser -G appuser

# Set working directory (STANDARD PRACTICE)
WORKDIR /app

# Copy binary and assets with correct permissions
COPY --from=builder --chown=appuser:appuser /app/ascii-art-web .
COPY --chown=appuser:appuser ./templates ./templates
COPY --chown=appuser:appuser ./banners ./banners

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check (Docker best practice)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./ascii-art-web"]