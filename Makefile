
run:
	@go run .

build:
	@go build -o bin/goproj cmd/app/main.go

tests:
	@go test .

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
