format:
	go fmt ./...

build:
	go build cmd/webserver/main.go

run: build
	./main