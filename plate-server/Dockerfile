FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go version
RUN go env
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /server /server

EXPOSE 5050

CMD ["/server"]