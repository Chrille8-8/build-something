# decrypt/Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download


RUN go build -o encrypt_service


CMD ["./encrypt_service", "-port", "5001"]

