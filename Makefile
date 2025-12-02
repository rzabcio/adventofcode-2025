.PHONY: clean test all

all: clean test build

clean:
	@rm -f aoc2024

test:
	@go test

build:
	@go build -o aoc2024
