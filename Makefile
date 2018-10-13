all: build

build:
	go build -o bin/consumer cmd/consumer/*.go
	go build -o bin/producer cmd/producer/*.go

tools:
	go get -u github.com/golang/dep/cmd/dep

devtools:
	go get github.com/vektra/mockery/.../

deps:
	dep ensure --vendor-only

clean:
	rm -rf bin