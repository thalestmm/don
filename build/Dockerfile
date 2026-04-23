FROM golang:1.26-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 2000

CMD ["./main"]
