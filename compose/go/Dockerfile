FROM golang:1.24.7-bookworm AS builder

WORKDIR /app

COPY  go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main.o

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main.o .

CMD ["./main.o"]