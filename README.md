# mock-oauth-svc

Simple oAuth mock server for microservices integration test.

It supports Basic Auth for clients - you can set it by environment variables AUTH_CLIENT and AUTH_SECRET. 
If not, it will use default client:secret. 

Main propose for this project is to emulate Spring Boot oAuth Server, so it should return erectly the same response 
as Spring.

For any user when password is equale to login name, this server will return new oAuth token

```
GET http://localhost:9005/oauth/token?grant_type=password&username=admin&password=admin
Accept: */*
Authorization: Basic Y2xpZW50OnNlY3JldA==  
```

Check token

```
GET http://localhost:9001/oauth/check_token?token=2e4228e4-ffd7-41ff-abf5-aa5d105abd79
Accept: application/json
Cache-Control: no-cache
```

This project use Go-kit and Allegro BigCache.

### TODO 

- docker container
- verify password == username
- refresh token suport
- gRPC endpoints
