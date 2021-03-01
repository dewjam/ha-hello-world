build: test
	go build -o bin/ha-hello-world main.go
	
run: build
	go run main.go

test:
	go fmt
	go vet
	go test main.go