package project

import (
	"embed"
	"fmt"
	"log"

	"github.com/ProstoyVadila/goproj/cmd/input"
	"github.com/ProstoyVadila/goproj/internal/config"
	"github.com/ProstoyVadila/goproj/internal/git"
	"github.com/ProstoyVadila/goproj/internal/models"
	"github.com/ProstoyVadila/goproj/internal/vscode"
	"github.com/ProstoyVadila/goproj/pkg/files"
	"github.com/ProstoyVadila/goproj/pkg/folders"
	"github.com/ProstoyVadila/goproj/pkg/reader"
)

// some wierd Go embed magic here
//
//go:embed templates/* templates/files/*
var EmbedFiles embed.FS

// showSetup writes setup info from CLI to standart output.
func showSetup(setup *models.Setup) {
	fmt.Printf("\nProject (package) name: %s\n", setup.PackageName)
	fmt.Printf("Author: %s\n", setup.Author)
	fmt.Printf("Description: %s\n", setup.Description)
	fmt.Printf("Files to skip: %v\n", setup.FilesToSkip)
	fmt.Printf("Folders to skip: %v\n\n", setup.FoldersToSkip)
	fmt.Printf("Init Git Repo: %v\n", setup.InitGit)
	fmt.Printf("Open in VS Code: %v\n", setup.InitVSCode)
}

// Generate creates files and initialize git repo with data from CLI or input.
func Generate(dataFromCli ...*models.Setup) {
	fmt.Println("Let's start!")

	var setup *models.Setup
	var err error

	// trying to get setup from the configuration file
	conf, err := config.Get()
	if err != nil {
		log.Print(err)
	}

	log.Println(conf)
	if err == nil {
		filesToSkip := reader.GetFilesToSkip(conf.Skip)
		foldersToSkip := reader.GetFoldersToSkip(conf.Skip)
		setup = models.NewSetupFromConfig(conf, filesToSkip, foldersToSkip)
	}

	// getting package name and other info from CLI or input
	if len(dataFromCli) == 1 {
		models.UpdateSetup(setup, dataFromCli[0])
	} else {
		inputSetup, err := input.GetSetup()
		if err != nil {
			log.Fatal()
		}
		models.UpdateSetup(setup, inputSetup)
	}

	// Show final Setup
	showSetup(setup)

	// aggregating all info about the projct to generate
	projectInfo := models.NewProjectInfo(setup)
	projectInfo.EmbedFiles = EmbedFiles

	// generating files
	if err = files.Generate(projectInfo); err != nil {
		log.Fatal(err)
	}

	// generating folders
	if err = folders.Create(projectInfo.Folders); err != nil {
		log.Fatal(err)
	}

	// Git init
	if projectInfo.InitGit {
		if err = git.InitGitRepo(projectInfo); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("successfully generated!")

	// open VS Code
	if projectInfo.InitVSCode {
		if _ = vscode.InitVSCode(); err != nil {
			log.Println(err)
		}
	}
}
