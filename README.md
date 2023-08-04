# Go Project Template Generator

This small utility allows you to start your new project in Go with an already initialized standard [folders](#list-of-generated-folders) and default support files such as [Makefile](#list-of-generated-files), [Dockerfile](#list-of-generated-files), [README.md](#list-of-generated-files), [LICINSE](#list-of-generated-files) etc. It creates a git repo as well.

Init project structure:
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

## Content
- [File Description](#file-description)
- [Installation](#installation)
- [Usage](#usage)
- [Plans](#plans)

## File Description
### List of Generated Files
1. <b>go.mod</b> – generates with the entered package name and your version of Go.

2. <b>main.go</b> – an empty file with `package main` and `func main()`.

3. <b>.env</b> – an empty env file.

4. <b>LICENSE</b> – [The MIT license](https://opensource.org/license/mit/) with your entered name and the current year.

5. <b>Dockerfile</b> – multi-stage build dockerfile setup with your version of Go.
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

## Installation
Default installation to `GOPATH/bin` folder
```bash
go install github.com/ProstoyVadila/goproj@latest
```
!!! Make sure that `GOPATH/bin` is in your `PATH`.

## Usage
Just type `goproj` in your terminal in <b>your project folder</b>
```bash
goproj
```
and answer a few questions:
```
Let's start!
Please, enter your name: Bob
Please, enter your new project (package) name: github.com/Bobert/new_project
Please, add a description to your project: my new project 
```
That's it! \
Please, Enjoy! :)

## Plans
- udpate cli ui/ux to make it prettier and more fun
- add an option to choose a [license](https://choosealicense.com/)
- add cli args to skip a starter quiz
- add initial setup configuration (to do setup only once and forever)
