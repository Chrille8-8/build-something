# decrypt/Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o decrypt_service ./decrypt

CMD ["./decrypt_service"]

