# Go Project Generator

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
![GitHub Release](https://badgen.net/github/release/ProstoyVadila/goproj?cachebust=2)
[![Go Reference](https://pkg.go.dev/badge/github.com/ProstoyVadila/goproj.svg)](https://pkg.go.dev/github.com/ProstoyVadila/goproj)

A CLI tool to initialize a Go project with customizable default folders and files.

## Overview

This small utility allows you to start your new project in Go with an already initialized standard [folders](#list-of-generated-folders) and default support files such as [Makefile](#list-of-generated-files), [Dockerfile](#list-of-generated-files), [README.md](#list-of-generated-files), [LICINSE](#list-of-generated-files) etc. It **creates a git repo** as well and tries to **open the new project in VS Code** by default. (You can change this behaviour by flags in [CLI](#command-line-interface))

Init project structure:

```bash
.
└── your_project
    ├── cmd/
    ├── internal/
    ├── pkg/
    ├── main.go
    ├── go.mod
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── .env
    ├── .dockerignore
    └── .gitignore
```

### Available Commands:

- **new**                Generates a new Go porject with default files and folders in the new folder
- **init**                  Generates a new Go porject with default files and folders in the path
- **config**            Sets up global configuration for all new generated projects
- **help**                Help about any command
- **completion**    Generates the autocompletion script for the specified shell
- **version**           Gets version info

## Content

- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Command-line interface](#command-line-interface)
- [Project Structure](#project-structure)
  - [Full List of Files](#a-full-list-of-generated-files)
  - [Full List of Folders](#a-full-list-of-generated-folders)
- [Configuration](#configuration)
  - [Usage](#usage-1)
  - [Show](#show)
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

#### Init Command

```bash
goproj init <your_new_package>
```

![](/examples/usage/example_init.GIF)

or with

#### New Command

```bash
goproj new <your_new_package>
```

This command generates a new project in the new folder. It can resolve folder's name by the entered project's name. \
For example, `goproj new github.com/bob/mylib` generates everything in `mylib` folder.

#### Arguments

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
  -g, --git                   an optional flag to define start git initialization or not (default false)
  -c, --vscode                an optional flag to open the new project in VS Code (default false)
  -s, --skip                  an optional flag to skip exact files and/or folders (add / after folder's name) from the generation.
  -h, --help                  help for init
```

You can find more information with `goproj init --help` or `goproj new --help` commands.

## Project Structure

### A Full List of Generated Files

1. **go.mod** – generates with the entered package name and your version of Go.

2. **main.go** – an empty file with `package main` and `func main()`.

3. **.env** – an empty env file.

4. **LICENSE** – [The MIT license](https://opensource.org/license/mit/) with your entered name and the current year.

5. **Dockerfile** – multi-stage build dockerfile setup with your version of Go.

```Dockerfile
FROM golang:1.22-alpine AS builder
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

You can set a global configuration for your new projects by command `gorpoj config` with args in CLI. It will create `goproj.config.toml` file in `~/.config/goproj` folder. And generator will read it every time when you start a new project (additinal args will override config setup for that project)

### Usage

Just type

```bash
goproj config
```

and answer a few questions like in [`goproj init`](#quick-start) command.

You can set a global config by providing a file as well. Goproj supports `json`, `yaml` and `toml` file extensions.
For example:

```bash
goproj config -f my_config.toml
```

You can find examples of config files [here](examples)

Or you can set it with additional flags like in the [gorpoj init](#command-line-interface) mode.
For example:

```bash
goproj config -a "Bobert Doe" -p "github.com/bobert_doe" --skip="Dockerfile,.dockerignore,internal/,pkg/" --git=false --vscode=false
```

Or you can set/change a global config manually by creating/changing `~/.config/goproj/goproj.config.toml` file (It doesn't exist by default).

This is a list of all flags:

```
Set up global configuration for all new generated projects to not do it every time

Usage:
   config [flags]

Examples:
goproj config -a "Bobert Doe" -s="Dockerfile,.dockerignore,internal/,pkg/" --git=false --vscode=false

Flags:
  -a, --author string   an optional flag to set author name
  -f, --file json       an optional flag to set information from yaml file (supprots `json`, `yaml`, `toml`)
  -p, --prefix string   an optional flag to set repo prefix for your new projects (example: "github.com/<smth>")
  -g, --git             an optional flag to define start git initialization or not (default false)
  -c, --vscode          an optional flag to open the new project in VS Code (default false)
  -s, --skip            an optional flag to skip exact files and/or folders (add / after folder's name) from the generation
  -h, --help            help for config
```

You can find them with `gorpoj config --help` command.

### Show

You can check your global config with:

```bash
goproj config show
```

That's it! \
Please, enjoy! :)

## Plans

- an option to set another folders or empty files to generate in the global config.
- an option to choose a [license](https://choosealicense.com/).
