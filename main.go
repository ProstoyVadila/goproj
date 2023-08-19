package main

import (
	"log"

	"github.com/ProstoyVadila/goproj/cmd/cli"
	"github.com/ProstoyVadila/goproj/internal/project"
	"github.com/ProstoyVadila/goproj/pkg/output"
)

const logo = `
_________              ________                ________ 
__  ____/______        ___  __ \______________ ______(_)
_  / __  _  __ \       __  /_/ /__  ___/_  __ \_____  / 
/ /_/ /  / /_/ /       _  ____/ _  /    / /_/ /____  /  
\____/   \____/        /_/      /_/     \____/ ___  /   
                                               /___/  
______________________________________________________ 
 
_______ A CLI tool to initialize a Go project ________
____ with customizable default folders and files _____
______________________________________________________

`

func main() {
	output.Info(logo)

	if cli.ArgsInCLI() {
		if err := cli.Execute(); err != nil {
			log.Fatal(err)
		}
	} else {
		project.Generate()
	}
}
