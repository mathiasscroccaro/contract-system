format:
	go fmt internal/...

build:
	go build cmd/webserver/main.go

run: build
	./main