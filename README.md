# Go Project Template Generator

This small utility allows you to start your new project in Go with an already initialized standard `folders` and default support files such as `Makefile`, `Dockerfile`, `README.md`, `LICINSE` etc. It creates a git repo as well.

Init project structre:
```bash
.
├── cmd/
├── internal/
├── pkg/
├── tests/
├── main.go
├── Makefile
├── Dockerfile
├── LICENSE
├── README.md
├── go.mod
├── .dockerignore
├── .gitignore
├── .env
└── .git
```

### List of Generated Files
1. <b>go.mod</b> – generates with the entered package name and your version of Go.

2. <b>main.go</b> – an empty file with `package main` and `func main()`.

3. <b>.env</b> – an empty env file.

4. <b>LICENSE</b> – [The MIT license](https://opensource.org/license/mit/) with your entered name and the current year.

5. <b>Dockerfile</b> – two-stage build dockerfile setup with your version of Go.
```Dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN env GOOS=linux CGO_ENABLED=0 go build -ldflags "-w" -o main

FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

CMD ["/app/main"]
```

6. <b>Makefile</b> – a simple Makefile with `run`/`build`/`tests` commands.
```Makefile
run:
	@go run .

build:
	@go build -o bin/main

tests:
	@go tests .


.PHONY: run build tests
```

7. <b>.gitignore</b> – a default [gitignore template](https://github.com/github/gitignore/blob/main/Go.gitignore) for projects in Go.
```gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
```

8. <b>.dockerignore</b> – a default [dockerignore template](https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/.dockerignore) for projects in Go. 
```dockerignore
# The .dockerignore file excludes files from the container build process.
#
# https://docs.docker.com/engine/reference/builder/#dockerignore-file

# Exclude locally vendored dependencies.
vendor/

# Exclude "build-time" ignore files.
.dockerignore
.gcloudignore

# Exclude git history and configuration.
.gitignore
```

### List of Generated Folders
1. <b>cmd</b>
2. <b>internal</b>
3. <b>pkg</b>
4. <b>tests</b>

## Installation (TODO rewrite)
Default installation to Go root folder (as utility)
```bash
# go install github.com/ProstoyVadila/goprojtepmplate
```

## Usage
...
```bash
```

## Todo
- pretty cli ui
- add an option to choose a [license](https://choosealicense.com/)
- add cli args to skip a starter quiz
- add initial setup configuration (???)
- rewrite readme lol
