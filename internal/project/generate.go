package project

import (
	"embed"
	"fmt"
	"log"

	"github.com/ProstoyVadila/goproj/cmd/input"
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
func Generate(dataFromCli ...*models.Setup) {
	fmt.Println("Let's start!")

	var setup *models.Setup
	var err error

	if len(dataFromCli) == 1 {
		setup = dataFromCli[0]
	} else {
		setup, err = input.GetSetup()
		if err != nil {
			log.Fatal()
		}
	}

	projectInfo := models.NewProjectInfo(setup)
	projectInfo.EmbedFiles = EmbedFiles

	if err = files.Generate(projectInfo); err != nil {
		log.Fatal(err)
	}

	if err = folders.Create(projectInfo.Folders); err != nil {
		log.Fatal(err)
	}

	if projectInfo.InitGit {
		if err = git.InitGitRepo(projectInfo); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("successfully generated!")

	if projectInfo.InitVSCode {
		if err = vscode.InitVSCode(); err != nil {
			log.Fatal(err)
		}
	}
}
