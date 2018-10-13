all: build

build:
	go build -o bin/go-simple-rabbitmq-server-client cmd/*.go

tools:
	go get -u github.com/golang/dep/cmd/dep

devtools:
	go get github.com/vektra/mockery/.../

deps:
	dep ensure --vendor-only

clean:
	rm -rf bin