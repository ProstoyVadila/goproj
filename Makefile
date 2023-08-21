
build:
	@echo "Building binary"
	@go build -o bin/goproj main.go

tests:
	@go test -v -cover ./...

i:
	@echo "Installing in ~/go/bin"
	@go install


.PHONY: build tests i
