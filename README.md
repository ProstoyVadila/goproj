# Go Project Generator

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
![GitHub Release](https://badgen.net/github/release/ProstoyVadila/goproj)

A CLI tool to initialize a Go project with customizable default folders and files.

## Overview

This small utility allows you to start your new project in Go with an already initialized standard [folders](#list-of-generated-folders) and default support files such as [Makefile](#list-of-generated-files), [Dockerfile](#list-of-generated-files), [README.md](#list-of-generated-files), [LICINSE](#list-of-generated-files) etc. It **creates a git repo** as well and tries to **open the new project in VS Code** by default. (You can change this behaviour by flags in [CLI](#command-line-interface))

Init project structure:

```bash
.
├── cmd/
├── internal/
├── pkg/
├── main.go
├── Makefile
├── Dockerfile
├── LICENSE
├── README.md
├── go.mod
├── .dockerignore
└── .gitignore
```

## Content

- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Command-line interface](#command-line-interface)
- [Project Structure](#project-structure)
  - [Full List of Files](#a-full-list-of-generated-files)
  - [Full List of Folders](#a-full-list-of-generated-folders)
- [Configuration](#configuration)
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
![](/examples/usage/example_full_input.GIF)

### Command-line interface

Another option is to use CLI arguments to make it more in a Go way:

```bash
goproj init <your_new_package>
```

![](/examples/usage/example_init.GIF)

\
You can specify some parameters with optional flags. For example:

```bash
goproj init github.com/Bobert/new_app --author Bob -d="My new project" --skip="Dockerfile,.dokerignore,internal/,pkg/"
```

\
There is a description of all flags and options:

```
Flags:
  -a, --author string         an optional flag to set your name
  -d, --description strings   an optional flag to set a description of your project
  -g, --git                   an optional flag to define start git initialization or not (default true)
  -s, --skip                  an optional flag to skip exact files and/or folders (add /) from the generation.
  -c, --vscode                an optional flag to open the new project in VS Code (default true
  -h, --help                  help for init
```

You can find more information with `goproj init --help` command.

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
	@go test .


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

## Configuration

You can set a global configuration for your new projects by command `gorpoj config` with args in CLI. It will create `~/.goproj.config.toml` config file in your user folder. And generator will read it every time when you start a new project (additinal args will override config setup for that project only)

You can set a global config by providing a file. Goproj supports `json`, `yaml` and `toml` file extensions.
For example:

```bash
goproj config -f ~Documets/my_config.toml
```

You can find examples of files [here](examples)

Or you can set it with additional flags line in the gorpoj init mode.
For example:

```bash
goproj config -a "Bobert Doe" --skip="Dockerfile,.dockerignore,internal/,pkg/" --git=false --vscode=false
```

Or you can set/change a global config manually by creating/changing `~/.goproj.config.toml` file in your user folder (It doesn't exist by default).

There is a list of all flags:

```
Set up global configuration for all new generated projects to not do it every time

Usage:
   config [flags]

Examples:
goproj config -a "Bobert Doe" -s="Dockerfile,.dockerignore,internal/,pkg/" --git=false --vscode=false

Flags:
  -a, --author string   an optional flag to set author name
  -f, --file json       an optional flag to set information from yaml file (supprots `json`, `yaml`, `toml`)
  -g, --git             an optional flag to define start git initialization or not (default true)
  -h, --help            help for config
  -s, --skip /          an optional flag to skip exact files and/or folders (add /) from the generation
  -c, --vscode          an optional flag to open the new project in VS Code (default true)
```

You can find them with `gorpoj config --help` command.

That's it! \
Please, enjoy! :)

## Plans

- add CI tests
- add an option to choose a [license](https://choosealicense.com/)
