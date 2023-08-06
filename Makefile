
run:
	@go run .

build:
	@go build -o bin/goproj cmd/app/main.go

tests:
	@go test -v -cover ./...

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
