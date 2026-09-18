FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /server .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /server /server

EXPOSE 5050

CMD ["/server"]