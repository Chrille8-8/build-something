FROM golang:latest
WORKDIR /server
COPY . .
RUN go mod download
RUN go build -o server
CMD ["./server"]
    