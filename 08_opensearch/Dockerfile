# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Stage 2: Final Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/server .

# Expose port and run
EXPOSE 8080
CMD ["./server"]