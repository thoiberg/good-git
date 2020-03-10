.PHONY: build format
# .DEFAULT build

build:
	go build -o gg

format:
	gofmt -w ./