package main

import (
	"log"

	"github.com/ProstoyVadila/goproj/cmd/cli"
	"github.com/ProstoyVadila/goproj/internal/project"
)

func main() {

	if cli.ArgsInCLI() {
		if err := cli.Execute(); err != nil {
			log.Fatal(err)
		}
	} else {
		project.Generate()
	}
}
