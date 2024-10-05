build:
	go build -o bin/docker-ui main.go


run: build
	./bin/docker-ui

