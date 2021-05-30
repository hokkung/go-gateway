all: test build 

generate:
	go generate ./...

test: generate
	go test -tag=test

build: generate
	go build $(BUILDFLAGS)
