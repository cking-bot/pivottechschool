default: test

test:
	cd calculator
	go test -v ./...

build:
	cd cmd/calculator && go test -v ./...
	go build -o calculator && go build -o calculator