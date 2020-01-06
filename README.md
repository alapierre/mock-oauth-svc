# mock-oauth-svc

Simple oAuth mock server for microservices integration test.

It supports Basic Auth for clients - you can set it by environment variables AUTH_CLIENT and AUTH_SECRET. 
If not, it will use default credentials client:secret. 

Main propose for this project is to emulate Spring Boot oAuth Server, so it should return exactly the same response 
as Spring.

## Run with Docker

docker-compose.yml

```
version: '3'
services:
  auth:
    image: lapierre/mock-oauth:0.0.1
    ports:
     - "9005:9005"
``` 

For any user with password equals to login name, this server will return new oAuth token. Grant type is ignored.

```
GET http://localhost:9005/oauth/token?grant_type=password&username=admin&password=admin
Accept: */*
Authorization: Basic Y2xpZW50OnNlY3JldA==  
```

Check token

```
GET http://localhost:9005/oauth/check_token?token=2e4228e4-ffd7-41ff-abf5-aa5d105abd79
Accept: application/json
Cache-Control: no-cache
```

This project use Go-kit and Allegro BigCache.

### TODO 
 
- users and client details from json file (almost ready)
- run docker container on non root user
- refresh token suport
- gRPC endpoints (preparing on branch)
