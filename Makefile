GONAME=$(shell basename "$(PWD)")
GOFILES=$(wildcard *.go)

all: clean test build

build:
	@echo "Building files"
	go build -o bin/$(GONAME) $(GOFILES) 

test:
	@echo "Running tests"
	go test $(GOFILES)

clean:
	@echo "Removing existing bins"
	go clean
	rm -rf bin/

run:
	@echo "Running the program"
	./bin/$(GONAME)

.PHONY: all build test clean setup run