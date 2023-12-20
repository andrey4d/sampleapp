### This is a simple web application.

The application contents two servers:
1. API server
2. WEB Server

Application don't use any dependency packages.

### Build container from scratch
1. Build from github.com 
 ```shell
 podman build --tag=sampleapp https://github.com/andrey4d/sampleapp.git
 ```
2. Building with Makefile
```shell
make container
```
3. You can redefine build tool and application image name
```shell
make -e BUILDER=buildah -e IMAGE_NAME=go-demo-app container
```
### Running containers with podman
```shell
make run
```



## Local Build
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