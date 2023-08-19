package project

import (
	"embed"
	"fmt"
	"log"

	"github.com/ProstoyVadila/goproj/internal/git"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/vscode"
	"github.com/ProstoyVadila/goproj/pkg/files"
	"github.com/ProstoyVadila/goproj/pkg/folders"
)

// some wierd Go embed magic here
//
//go:embed templates/* templates/files/*
var EmbedFiles embed.FS

// Generate creates files and initialize git repo with data from CLI or input.
func Generate(ArgsSetup ...*models.Setup) {
	fmt.Println("Let's start!")

	// trying to get setup from the configuration file or CLI args, or Input
	setup := enrichSetup(ArgsSetup...)

	// Show final Setup
	setup.Show()

	// aggregating all info about the projct to generate
	projectInfo := models.NewProjectInfo(setup)
	projectInfo.EmbedFiles = EmbedFiles

	// generating files
	fmt.Println("Generating files...")
	if err := files.Generate(projectInfo); err != nil {
		log.Fatal(err)
	}

	// generating folders
	fmt.Println("Generating folders...")
	if err := folders.Create(projectInfo.Folders); err != nil {
		log.Fatal(err)
	}

	// Git init
	if projectInfo.InitGit {
		if err := git.InitGitRepo(projectInfo); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("\nSuccessfully generated!")

	// open VS Code
	if projectInfo.InitVSCode {
		if err := vscode.InitVSCode(); err != nil {
			log.Println(err)
		}
	}
}
