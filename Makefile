
build:
	@echo "Building binary"
	@go build -o bin/goproj main.go

tests:
	@go test -v -cover ./...

i:
	@echo "Installing in ~/go/bin"
	@go install

tidy:
	@echo "Running go mod tidy"
	@go mod tidy

.PHONY: build tests i tidy
