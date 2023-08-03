
run:
	go run .

build:
	go build .

tests:
	go tests .

cleanup:
	rm tests/files/*

.PHONY: run build tests cleanup
