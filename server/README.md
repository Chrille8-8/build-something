# Start server

### Build this microservice
```bash
docker build -t start_server .
```


### Run this microservice
```bash
docker run -p 8080:5000 start_server
```

Accesss the server at:
```bash
http://172.17.0.2:8080/input
```


## Ideas for the databas
The database could be used like this:
When a user want a text encrypted, the user should enter his name as well so this could be stored in the database together with the key and text 



