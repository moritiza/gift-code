# Simple Gift Code Management
## About
Very simple gift code management system just for presenting


## How can I run it?
### #1 Make sure the docker is installed
```bash
docker --version
```
##### If the docker is not installed, you can use this link
<a href="https://docs.docker.com/engine/install/">Install Docker</a>

### #2 Make sure the docker service is running
```bash
sudo systemctl start docker.service
```

### #3 Execute docker-compose build command
```bash
sudo docker-compose build --no-cache
```

### #4 Execute docker-compose up command
```bash
sudo docker-compose up --force-recreate
```

### #5 Make sure everything is OK! :wink:
##### Send request to:
###### #1 Discount Service 'http://localhost:7000/api/ping':
```bash
curl -X GET "http://localhost:7000/api/ping" -H "Accept: application/json"
```
##### Result sample:
```json
{
  "status": true,
  "message": "ok",
  "errors": null,
  "data": "pong"
}
```

###### #2 Report Service 'http://localhost:8000/api/ping':
```bash
curl -X GET "http://localhost:8000/api/ping" -H "Accept: application/json"
```
##### Result sample:
```json
{
  "status": true,
  "message": "ok",
  "errors": null,
  "data": "pong"
}
```

###### #3 Wallet Service 'http://localhost:9000/api/ping':
```bash
curl -X GET "http://localhost:9000/api/ping" -H "Accept: application/json"
```
##### Result sample:
```json
{
  "status": true,
  "message": "ok",
  "errors": null,
  "data": "pong"
}
```

<hr>

## How can I dive into?
### Just import `thunder-collection_Gift_Code_Management.json` to VSCode `Thunder Client` extension and enjoy it! :smile:

<hr>

## How can I test it?
### #1 Get into Docker container shell
```bash
sudo docker exec -it discount_service_app bash
```

### #2 Execute `go test` command
```bash
go clean -testcache
go test ./...
```