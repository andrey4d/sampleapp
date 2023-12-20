.PHONY: help clean build

all: help

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

include .env
PROJECTNAME=$(shell basename "$(PWD)")

# Container builder
CONTAINER=podman
BUILDER=podman
IMAGE_NAME=sampleapp

BACKEND_URL="http://$(word 1, $(shell hostname -I)):8070"

# Go related variables.
BASE=$(shell pwd)
BIN=$(BASE)/bin
FILES=$(wildcard **/*.go)


build: go-build

clean: go-clean


## go-build: build binary `go build` internally.
go-build: 
	@echo "  >  Building binary..."
	go build -o $(BIN)/backend $(BASE)/cmd/backend/main.go
	go build -o $(BIN)/webserver $(BASE)/cmd/webserver/main.go

## go-clean: Clean build files. Runs `go clean` internally.
go-clean:
	@echo "  >  Cleaning build cache"
	rm -rf ./bin 
	go clean

## container: Build container from scratch
container:
	$(BUILDER) build -t $(IMAGE_NAME) -f Dockerfile .

## run: Start containers 
run:
	$(CONTAINER) run --rm -d -p 8070:8070 --name sampleapp-api localhost/sampleapp
	$(CONTAINER) run --rm -d -p 8080:8080 -e BACKEND_URL=$(BACKEND_URL) --name sampleapp-web localhost/sampleapp /opt/webserver

## stop: Stop containers 
stop:
	$(CONTAINER) stop sampleapp-api sampleapp-web
