
build:
	@echo "Building binary"
	@go build -o bin/goproj main.go

tests:
	@echo "Running tests"
	@go test .

i:
	@echo "Installing in ~/go/bin"
	@go install


.PHONY: build tests i
