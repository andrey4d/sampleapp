### This is a simple web application.

The application contents two servers:
1. API server
2. WEB Server

### Build
```shell
make build
```
### Run API server
```shell
bin/bakend
```
### Run WEB server
```shell
BACKEND_URL="http://localhost:8070" bin/webserver
```

### Get about
```
curl localhost:8080/about
```