
run:
	@go run .

build:
	@go build -o bin/main

tests:
	@go tests .

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
