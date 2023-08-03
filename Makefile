
run:
	go run .

build:
	go build

tests:
	go tests .

cleanup:
	rm tests/tempFiles/*

.PHONY: run build tests cleanup
