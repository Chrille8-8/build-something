# decrypt/Dockerfile
FROM golang:latest
WORKDIR /encryption
COPY . .
RUN go mod download
RUN go build -o encrypt_service
CMD ["./encrypt_service"]

