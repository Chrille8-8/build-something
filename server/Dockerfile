FROM golang:latest

# Set the working directory inside the container to /server
WORKDIR /server

# Copy the stuff of the current dir into the conterners /server dir
COPY . .

# Download dependencies
RUN go mod download

RUN go build -o server

CMD ["./server", "-port", "5000"]
#CMD ["./server"]
    