FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o chibi-maruko main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates curl

WORKDIR /app
COPY --from=builder /app/chibi-maruko .
COPY static ./static
EXPOSE 5544

CMD ["./chibi-maruko"]
