
run:
	@go run .

build:
	@go build -o bin/goproj

tests:
	@go tests .

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
