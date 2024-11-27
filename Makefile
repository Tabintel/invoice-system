.PHONY: build test run clean

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

run:
	go run cmd/api/main.go

clean:
	rm -rf bin/
