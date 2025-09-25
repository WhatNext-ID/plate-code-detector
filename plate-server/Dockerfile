# Build stage
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Final stage
FROM alpine:latest
WORKDIR /
COPY --from=builder /server /server

# Minimal runtime
CMD ["/server"]
