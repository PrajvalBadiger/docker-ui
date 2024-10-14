build:
	go mod tidy
	go build -o bin/docker-ui main.go


run: build
	./bin/docker-ui

