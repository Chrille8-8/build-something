# Encryption server


### Build the encryption server
```bash
docker build -t encrypt_server .
```

### Run the encryption server
```bash
docker run -p 8080:8080 encrypt_server
```

### Run the encryption server

```bash
http://172.17.0.2:8080/encrypt
```
The service could be accesses through insomnia or curl, the service only takes a POST requests.

#### Body for the POST request
```bash
{"key": 3, "text": "abcz"}
```

### Debug the server
#### Check the ip-address to connect to the service
Get the name or id for the service to debug
```bash
docker ps
```
then inspect the service
```bash
docker inspect <image id | image name>
```