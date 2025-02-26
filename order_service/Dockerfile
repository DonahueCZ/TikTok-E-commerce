FROM golang:1.23-alpine AS builder

WORKDIR /usr/src/app

COPY . .

RUN go build -o /tmp/order ./rpc/order/.

FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /tmp/order /usr/src/app/

CMD ["./order"]