package project

import (
	"embed"
	"log"

	"github.com/ProstoyVadila/goproj/cmd/input"
	"github.com/ProstoyVadila/goproj/internal/git"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/pkg/files"
	"github.com/ProstoyVadila/goproj/pkg/folders"
)

//go:embed templates/* templates/files/*
var EmbedFiles embed.FS

// Generate creates files and initialize git repo with data from CLI or input.
func Generate(dataFromCli ...*models.Setup) {
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

	err = files.Generate(projectInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = folders.Create(projectInfo.Folders)
	if err != nil {
		log.Fatal(err)
	}

	if projectInfo.InitGit {
		err = git.InitGitRepo(projectInfo)
		if err != nil {
			log.Fatal(err)
		}
	}
}
