FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]