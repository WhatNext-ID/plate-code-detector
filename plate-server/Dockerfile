# =========================
# Build stage
# =========================
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Dependencies first for Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Source code
COPY . .

# Build static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /server .


# =========================
# Runtime stage
# =========================
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /server /server

EXPOSE 5050

CMD ["/server"]