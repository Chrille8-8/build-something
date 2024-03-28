# Start service

## Build & start using minikube & kubernetes


### Steps

<ol>
    <li> go to project dir: <code>cd project</code>
    <li> run: <code>make build</code> then <code>make start_service</code>
    <li> Wait a minute or two so that the containers can be created.
    <li> If minikube wont open the site in a new window, use ctrl and left-click on the URL that is shown in the terminal, example belov:
</ol>

```console
minikube service server-service
|-----------|----------------|-------------|---------------------------|
| NAMESPACE |      NAME      | TARGET PORT |            URL            |
|-----------|----------------|-------------|---------------------------|
| default   | server-service |        5000 |  http://192.168.4.2:32596 |
|-----------|----------------|-------------|---------------------------|
```


#  Endpoints
The endpoints you can access is
Available Endpoints:

- `/`: Main endpoint, this is where you can encrypt your text
- `/getall`: Retrieve all data from database
