
run:
	@go run .

build:
	@go build -o bin/goproj main.go

tests:
	@go test .

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
