package main

import (
	"fmt"
	"log"

	"github.com/ProstoyVadila/goproj/cmd/cli"
	"github.com/ProstoyVadila/goproj/internal/project"
)

func main() {
	fmt.Println("Let's start!")

	if cli.ArgsInCLI() {
		if err := cli.Execute(); err != nil {
			log.Fatal(err)
		}
	} else {
		project.Generate()
	}

	fmt.Println("Successfully generated!")
}
