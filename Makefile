commit := $(shell git rev-parse --short HEAD)

all: fmt build run

fmt:
	go fmt main.go

build:
	echo $(commit)
	go build 

run:
	./webapp

clear:
	rm ./webapp
