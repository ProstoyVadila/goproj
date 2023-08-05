# Go Project Template Generator

An utility to initialize a Go project with default folders and files.

## Overview

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
- [Quick Start](#quick-start)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Command-line interface](#command-line-interface)
- [Project Structure](#project-structure)
    - [Full List of Files](#a-full-list-of-generated-files)
    - [Full List of Folders](#a-full-list-of-generated-folders)
- [Plans](#plans)

## Quick Start

### Installation
Default installation to `GOPATH/bin` folder
```bash
go install github.com/ProstoyVadila/goproj@latest
```
**!!! Make sure that `GOPATH/bin` is in your `PATH`.**

### Usage
Just type `goproj` in your terminal in <b>your project folder</b>
```bash
goproj
```
and answer a few questions:
```
Let's start!
Please, enter your name: Bob
Please, enter your new project (package) name: github.com/Bobert/new_app
Please, add a description to your project: my new project 
```

### Command-line interface
Another option is to use CLI arguments to make it more in a Go way:
```bash
goproj init <your_new_package>
```
\
You can specify some parameters with optional flags. For example:
```bash
goproj init github.com/Bobert/new_app -a Bob -d="My new project" -s="Dockerfile,.dokerignore"
```
\
There is a description of all flags and options:
```
Usage:
   init [flags]

Flags:
  -a, --author string         an optional flag to set your name
  -d, --description strings   an optional flag to set a description of your project
  -s, --skip strings          an optional flag to skip exact files from the generation
  -h, --help                  help for init
  ```
You can find more information with `-h` or `--help` flags.

## Project Structure
### A Full List of Generated Files
1. **go.mod** – generates with the entered package name and your version of Go.

2. **main.go** – an empty file with `package main` and `func main()`.

3. **.env** – an empty env file.

4. **LICENSE** – [The MIT license](https://opensource.org/license/mit/) with your entered name and the current year.

5. **Dockerfile** – multi-stage build dockerfile setup with your version of Go.
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

6. **Makefile** – a simple Makefile with `run`/`build`/`tests` commands.
```Makefile
run:
	@go run .

build:
	@go build -o bin/main

tests:
	@go tests .


.PHONY: run build tests
```

7. **.gitignore** – a default [gitignore template](https://github.com/github/gitignore/blob/main/Go.gitignore) for projects in Go.
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

8. **.dockerignore** – a default [dockerignore template](https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/.dockerignore) for projects in Go. 
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

### A Full List of Generated Folders
There are some standart folders for any project in Go:
1. **cmd**
2. **internal**
3. **pkg**
4. **tests**

That's it! \
Please, enjoy! :)

## Plans
- udpate cli ui/ux to make it prettier and more fun
- add tests
- add option to skip exact files
- add option to create exact folders
- update readme [usage section](#usage) with cli args examples
- add an option to choose a [license](https://choosealicense.com/)
- add initial setup configuration (to do setup only once and forever)
