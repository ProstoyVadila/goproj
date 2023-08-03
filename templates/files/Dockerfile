FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main

FROM alpine:3.17 
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

CMD ["/app/main"]