test:
	go test -v ./...

build:
	go build -o ./tmp/main cmd/web/main.go

run:
	go run cmd/web/main.go

