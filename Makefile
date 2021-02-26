build: test
	go build -o bin/ha-hello-world main.go
	
run: build
	go run main.go

docker: test
	GOOS=linux GOARCH=amd64 go build -o bin/ha-hello-world main.go
	docker build -t ha-hello-world .
	docker run -it --rm --name ha-hello-world -p 8080:8080 ha-hello-world

test:
	go fmt
	go vet
	go test main.go