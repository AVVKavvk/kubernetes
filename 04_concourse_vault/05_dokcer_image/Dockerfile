# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files first (better layer caching)
COPY . . 
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage - tiny final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]